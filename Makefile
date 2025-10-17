all: generate-openapi-yaml openapicodegen generate-handlers cross-compile 

generate-handlers:
	go run cmd/handlers_generator/main.go

generate-openapi-yaml:
	python3 tasks_to_openapi.py

openapicodegen:
	mkdir -p pkg/openapi/v1 | exit;
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.16.3 \
	--package=openapiv1 --generate=types,models,echo-server,strict-server,spec \
	api/openapi/openapi.yaml > pkg/openapi/v1/openapi.gen.go

cross-compile:
	env GOOS=linux GOARCH=arm go build -o server cmd/main/main.go

# .PHONY: generate-openapi-yaml generate-openapi cross-compile