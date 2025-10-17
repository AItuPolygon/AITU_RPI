import yaml
import re

# Custom representer to ensure strings are treated literally
def str_presenter(dumper, data):
    return dumper.represent_scalar('tag:yaml.org,2002:str', data)

yaml.add_representer(str, str_presenter)

def parse_custom_yaml(file_path):
    """Reads and parses the custom YAML."""
    with open(file_path, 'r', encoding='utf-8') as file:
        return yaml.safe_load(file)

def handle_special_values(values):
    """Replace True/False with on/off where applicable."""
    new_keys = {}
    for key, val in values.items():
        if isinstance(key, bool):
            new_key = "on" if key else "off"
        else:
            new_key = key
        new_keys[new_key] = val
    return new_keys

def generate_openapi_spec(tasks, classes):
    """Generates the OpenAPI spec based on the tasks and class definitions."""
    openapi_spec = {
        "openapi": "3.0.0",
        "info": {
            "title": "Generated API",
            "version": "1.0.0"
        },
        "paths": {},
        "components": {
            "schemas": {}
        }
    }
    
    # Create reusable schema definitions
    for class_name, class_fields in classes.items():
        properties = {}
        required_fields = []

        for field_name, field in class_fields.items():
            field_schema = {"type": field["type"]}
            extra_tags = {}

            if field["type"] == "enum" and "values" in field:
                enum_values = handle_special_values(field["values"])
                field_schema["type"] = "string"
                field_schema["enum"] = list(enum_values.keys())
                extra_tags = {
                    "binding": f"required,oneof={','.join(field_schema['enum'])}"
                }
            elif field["type"] == "int" and "min" in field and "max" in field:
                field_schema["type"] = "integer"
                field_schema["minimum"] = field["min"]
                field_schema["maximum"] = field["max"]
                extra_tags = {
                    "binding": f"required,gte={field['min']},lte={field['max']}"
                }
            else:
                print(f"Warning: Unhandled or missing type fields for field '{field_name}' in class '{class_name}'")

            if extra_tags:
                field_schema["x-oapi-codegen-extra-tags"] = extra_tags

            properties[field_name] = field_schema

            if field.get("required", True):
                required_fields.append(field_name)

        openapi_spec["components"]["schemas"][class_name] = {
            "type": "object",
            "properties": properties,
            "required": required_fields
        }
    
    # Create paths with references to reusable components
    for task_name, task_details in tasks.items():
        for element_name, element_details in task_details["elements"].items():
            class_name = element_details["class"]
            if class_name not in classes:
                print(f"Warning: No details found for class '{class_name}'")
                continue
            description = element_details["description"]

            endpoint = f"/api/{task_name}/{element_name}"
            ref_schema = {
                "$ref": f"#/components/schemas/{class_name}"
            }

            openapi_spec["paths"][endpoint] = {
                "get": {
                    "summary": f"Get information from {endpoint}",
                    "description": f'{description}',
                    "responses": {
                        "200": {
                            "description": "Successful response",
                            "content": {
                                "application/json": {
                                    "schema": ref_schema
                                }
                            }
                        }
                    }
                },
                "post": {
                    "summary": f"Post data to {endpoint}",
                    "description": f"{description}",
                    "requestBody": {
                        "required": True,
                        "content": {
                            "application/json": {
                                "schema": ref_schema
                            }
                        }
                    },
                    "responses": {
                        "200": {
                            "description": "Successful response",
                            "content": {
                                "application/json": {
                                    "schema": ref_schema
                                }
                            }
                        }
                    }
                }
            }
    
    return openapi_spec

def write_openapi_file(openapi_spec, output_path):
    """Writes the OpenAPI spec to a file."""
    yaml_str = yaml.dump(openapi_spec, sort_keys=False, default_flow_style=False, allow_unicode=True)

    with open(output_path, 'w', encoding='utf-8') as file:
        file.write(yaml_str)

if __name__ == "__main__":
    input_path = "tasks.yaml"
    output_path = "./api/openapi/openapi.yaml"
    
    custom_data = parse_custom_yaml(input_path)
    openapi_spec = generate_openapi_spec(custom_data['tasks'], custom_data['classes'])
    write_openapi_file(openapi_spec, output_path)
