package plugin

import (
	"github.com/byorty/hardcore/meta"
	"strings"
	"github.com/byorty/hardcore/utils"
	"os"
	"regexp"
	"bytes"
	"text/template"
    "fmt"
    "path/filepath"
)

var (
	controllerRenameRegex = regexp.MustCompile(`([A-Z])`)
	controllerTpl = `package {{.Package}}

{{if .HasImports}}import ({{range .Imports}}
    "{{.}}"{{end}}
){{end}}

type {{.Name}} struct {
    {{range .Extends}}{{.Name}}
{{end}}}
`
	autoControllerTpl = `package {{.Package}}

import ({{range .AutoImports}}
    "{{.}}"{{end}}
)

func New{{.Name}}() types.ActionController {
    return new({{.Name}})
}

func ({{.ShortName}} *{{.Name}}) CallAction(action interface{}, scope types.RequestScope) {
    action.(func(*{{.Name}}, types.RequestScope))(t, scope)
}
`
)

type Controller struct {}

func (c *Controller) Do(env *meta.Environment) {
    conf := env.Configuration
	for _, container := range conf.ControllerContainers {
		if !utils.FileExists(container.Path) {
			err := os.MkdirAll(container.Path, os.ModePerm)
			if err != nil {
				env.Logger.Error(err)
				os.Exit(1)
			}
		}
		for _, controller := range container.Controllers {
			filename := strings.ToLower(controllerRenameRegex.ReplaceAllString(controller.Name, "_$1"))
			if filename[0] == '_' {
				filename = filename[1:]
			}
			filename = filepath.Join(container.Path, filename)
            autoFilename := fmt.Sprintf("%s_auto", filename)

            imports := make([]string, 0)
            for _, extend := range controller.Extends {
                imports = append(imports, extend.Import)
            }

            autoImports := make([]string, 0)
            autoImports = append(autoImports, "github.com/byorty/hardcore/types")

			tmplParams := map[string]interface{}{
				"ShortName": strings.ToLower(controller.Name[0:1]),
				"Name": controller.Name,
				"Package": container.Package,
				"HasImports": len(imports) > 0,
				"Imports": imports,
                "Extends": controller.Extends,
				"AutoImports": autoImports,
			}
			buf := new(bytes.Buffer)
			tmpl := template.New(filename  + "_tpl")
			tmpl.Parse(controllerTpl)
			tmpl.Execute(buf, tmplParams)
			env.Logger.Info(buf.String())


			buf = new(bytes.Buffer)
			tmpl = template.New(autoFilename  + "_tpl")
			tmpl.Parse(autoControllerTpl)
			tmpl.Execute(buf, tmplParams)
			env.Logger.Info(buf.String())


		}
	}
}
