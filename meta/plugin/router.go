package plugin

import (
	"github.com/byorty/hardcore/meta/controller"
	"github.com/byorty/hardcore/meta/types"
	"path/filepath"
)

var (
	autoRouterTpl = `package configs

import ({{range .AutoImports}}
	"{{.}}"{{end}}
)

var (
	Routers = []types.Route{
		{{if .Containers}}{{range .Containers}}{{$package := .GetShortPackage}}mux.Path("{{.GetRoute}}", {{range .Controllers}}{{$ctrlName := .GetName}}
			mux.Controller("{{.GetRoute}}", {{$package}}.New{{.GetName}}).Batch({{range .GetActions}}{{if .HasForm}}
				mux.{{.GetMethod}}("{{.GetRoute}}", {{$package}}.{{$ctrlName}}{{.GetName}}Action),{{else}}
				mux.{{.GetMethod}}("{{.GetRoute}}", (*{{$package}}.{{$ctrlName}}).{{.GetName}}),
			{{end}}{{end}}
			),{{end}}
		),{{end}}
	}{{end}}
)
`
)

type Router struct{}

func (r *Router) Do(env types.Environment) {
	autoImports := make([]string, 0)
	containers := make([]*controller.Container, 0)
	for _, container := range env.GetConfiguration().GetContainers() {
		if container.GetContainerKind() == types.ControllerContainerKind {
			ctrlContainer := container.(*controller.Container)
			ctrlContainer.Controllers = make([]*controller.Controller, 0)
			for _, entity := range container.GetEntities() {
				ctrlContainer.Controllers = append(ctrlContainer.Controllers, entity.(*controller.Controller))
			}
			containers = append(containers, ctrlContainer)

			hasImport := false
			for _, existsImport := range autoImports {
				if existsImport == container.GetImport() {
					hasImport = true
					break
				}
			}
			if !hasImport {
				autoImports = append(autoImports, container.GetImport())
			}
		}
	}
	tplParams := map[string]interface{}{
		"AutoImports": append([]string{
			types.DefaultImport,
			types.MuxImport,
		}, autoImports...),
		"Containers": containers,
	}

	filename := filepath.Join(env.GetAbsPath(), "configs", "router")
	env.GetConfiguration().AddAutoFile(filename+"_auto", autoRouterTpl, tplParams)
}
