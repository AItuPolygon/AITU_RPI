package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"
)

const handlerTemplate = `package app

import (
    "context"
    . "rpi-test/pkg/openapi/v1"
)

type Server struct {
    app *App
}

var _ StrictServerInterface = (*Server)(nil)

func newServer(app *App) *Server {
    return &Server{
        app: app,
    }
}

{{range .PostEndpoints}}
// {{.Comment}}
// (POST {{.Path}})
func (s *Server) {{.Name}}(ctx context.Context, request {{.RequestType}}) ({{.ResponseTypeInterface}}, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return {{.ResponseType}}{}, nil
}
{{end}}

{{range .GetEndpoints}}
// {{.Comment}}
// (GET {{.Path}})
func (s *Server) {{.Name}}(ctx context.Context, request {{.RequestType}}) ({{.ResponseTypeInterface}}, error) {
    result := &{{.ResponseType}}{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}
{{end}}
`

type Endpoint struct {
	Comment               string
	Path                  string
	Name                  string
	RequestType           string
	ResponseType          string
	ResponseTypeInterface string
}

type Data struct {
	PostEndpoints map[string]Endpoint
	GetEndpoints  map[string]Endpoint
}

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "pkg/openapi/v1/openapi.gen.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := Data{
		PostEndpoints: make(map[string]Endpoint),
		GetEndpoints:  make(map[string]Endpoint),
	}

	for _, f := range node.Decls {
		if fn, isFn := f.(*ast.FuncDecl); isFn {
			if fn.Recv != nil && len(fn.Recv.List) > 0 {
				if fn.Recv.List[0].Type != nil {
					methodName := fn.Name.Name
					comment := fn.Doc.Text()
					path := extractPath(comment)
					requestType := fmt.Sprintf("%sRequestObject", methodName)
					responseTypeInterface := fmt.Sprintf("%sResponseObject", methodName)
					responseType := fmt.Sprintf("%s200JSONResponse", methodName)
					if strings.HasPrefix(methodName, "Post") {
						data.PostEndpoints[methodName] = Endpoint{
							Comment:               comment,
							Path:                  path,
							Name:                  methodName,
							RequestType:           requestType,
							ResponseTypeInterface: responseTypeInterface,
							ResponseType:          responseType,
						}
					} else if strings.HasPrefix(methodName, "Get") {
						data.GetEndpoints[methodName] = Endpoint{
							Comment:               comment,
							Path:                  path,
							Name:                  methodName,
							RequestType:           requestType,
							ResponseTypeInterface: responseTypeInterface,
							ResponseType:          responseType,
						}
					}
				}
			}
		}
	}

	tmpl, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("internal/app/handlers.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Handlers generated successfully!")
}

func extractPath(comment string) string {
	lines := strings.Split(comment, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "(POST") || strings.HasPrefix(line, "(GET") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				return parts[1]
			}
		}
	}
	return ""
}
