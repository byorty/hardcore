package meta

import (
	"fmt"
	"text/template"
	"bytes"
)

var ControllerTmpl = `func New{{.Type}}() mux.ActionController {
    return new({{.Type}})
}

func (t *{{.Type}}) CallAction(action interface{}, scope mux.RequestScope) {
    action.(func(*{{.Type}}, mux.RequestScope))(t, scope)
}
`

type controllerBuilder struct {}

func (e *controllerBuilder) Build(controllers ...interface{}) []*BuildResult {
	results := make([]*BuildResult, 0)
	for _, controller := range controllers {
		info := AnalyzeStruct(controller)
		if info == nil {
			panic(fmt.Sprintf("controller should be a struct, '%v' given", controller))
		} else {
			buf := new(bytes.Buffer)
			tmplParams := map[string]interface{}{
				"Type": info.Name,
			}
			tmpl := template.New(info.Name + "Template")
			tmpl.Parse(ControllerTmpl)
			tmpl.Execute(buf, tmplParams)
			results = append(results, &BuildResult{
					Pkg: info.Pkg,
					PkgPath: info.PkgPath,
					Bytes: buf.Bytes(),
					Imports: []string{"github.com/byorty/hardcore/mux"},
				})
		}
	}
	return results
}
