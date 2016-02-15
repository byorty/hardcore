package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"path/filepath"
	"github.com/byorty/hardcore/meta/controller"
)

var (
	routerTpl = `package configs

import ({{range .Imports}}
	"{{.}}"{{end}}
)

func makeRouter() *mux.Route {
	return nil
}
`
	autoRouterTpl = `package configs

import ({{range .AutoImports}}
	"{{.}}"{{end}}
)

var (
	Router = mux.NewRouter(){{if .Containers}}.Add({{range .Containers}}{{$package := .GetShortPackage}}
		mux.Path("{{.GetRoute}}", {{range .Controllers}}{{$ctrlName := .GetName}}
			mux.Controller("{{.GetRoute}}", {{$package}}.New{{.GetName}}).Batch({{range .GetActions}}{{if .HasForm}}
				mux.{{.GetMethod}}("{{.GetRoute}}", {{$package}}.{{$ctrlName}}{{.GetName}}Action),{{else}}
				mux.{{.GetMethod}}("{{.GetRoute}}", (*{{$package}}.{{$ctrlName}}).{{.GetName}}),
			{{end}}{{end}}
			),{{end}}
		),
	{{end}}){{end}}.Add(makeRouter())
)
`
)

type Router struct {}

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
		"Imports": []string{
			types.MuxImport,
		},
		"AutoImports": append([]string{
			types.MuxImport,
		}, autoImports...),
		"Containers": containers,
	}

	filename := filepath.Join(env.GetAbsPath(), "configs", "router")
	env.GetConfiguration().AddAutoFile(filename + "_auto", autoRouterTpl, tplParams)
	env.GetConfiguration().AddFile(filename, routerTpl, tplParams)
}
