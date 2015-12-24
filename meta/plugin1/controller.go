package plugin

import (
	"github.com/byorty/hardcore/meta"
	"strings"
)

var (
	controllerTpl = `package {{.Package}}{{if .HasImports}}

import ({{range .Imports}}
    "{{.}}"{{end}}
)
{{else}}
{{end}}
type {{.Name}} struct { {{range .Extends}}
    {{.Name}}{{end}}
}

func New{{.Name}}() types.ActionController {
    return nil
}
`
	autoControllerTpl = `package {{.Package}}

import ({{range .AutoImports}}
    "{{.}}"{{end}}
)

func ({{.ShortName}} *{{.Name}}) CallAction(action interface{}, scope types.RequestScope) {
    action.(func(*{{.Name}}, types.RequestScope))(t, scope)
}
`
)

type Controller struct {}

func (c *Controller) Do(env *meta.Environment) {
    conf := env.Configuration
	for _, container := range conf.ControllerContainers {
		for _, controller := range container.Controllers {
            imports := make([]string, 0)
            imports = append(imports, "github.com/byorty/hardcore/types")
            for _, extend := range controller.Extends {
                imports = append(imports, extend.Import)
            }

            autoImports := make([]string, 0)
            autoImports = append(autoImports, "github.com/byorty/hardcore/types")

			tmplParams := map[string]interface{}{
				"ShortName": strings.ToLower(controller.Name[0:1]),
				"Name": controller.Name,
				"Package": container.ShortPackage,
				"HasImports": len(imports) > 0,
				"Imports": imports,
                "Extends": controller.Extends,
				"AutoImports": autoImports,
			}

            env.Configuration.AddAutoFile(controller.AutoFilename, autoControllerTpl, tmplParams)
            env.Configuration.AddFile(controller.Filename, controllerTpl, tmplParams)
		}
	}
}
