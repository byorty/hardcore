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
				mux.{{.GetRouteMethod}}("{{.GetRoute}}", {{$package}}.{{$ctrlName}}{{.GetName}}Action),{{else}}
				mux.{{.GetRouteMethod}}("{{.GetRoute}}", (*{{$package}}.{{$ctrlName}}).{{.GetName}}),
			{{end}}{{end}}
			){{range .Befores}}.
			Before({{.GetName}}){{end}}{{range .Afters}}.
			After({{.GetName}}){{end}},{{end}}
		){{range .Befores}}.
		Before({{.GetName}}){{end}}{{range .Afters}}.
		After({{.GetName}}){{end}},
		{{end}}
	}{{end}}
)
`
)

type Router struct {
	PackageImporter
}

func (r *Router) Do(env types.Environment) {
	containers := make([]*controller.Container, 0)
	for _, container := range env.GetConfiguration().GetContainers() {
		if container.GetContainerKind() == types.ControllerContainerKind {
			ctrlContainer := container.(*controller.Container)
			containers = append(containers, ctrlContainer)
			r.addImport(ctrlContainer.GetImport())
			r.addMiddlewaresImports(ctrlContainer.Befores)
			r.addMiddlewaresImports(ctrlContainer.Afters)
			for _, entity := range container.GetEntities() {
				if entity.GetEntityKind() == types.ControllerEntityKind {
					ctrl := entity.(*controller.Controller)
					r.addMiddlewaresImports(ctrl.Befores)
					r.addMiddlewaresImports(ctrl.Afters)
				}
			}

		}
	}

	if len(containers) > 0 {
		tplParams := map[string]interface{}{
			"AutoImports": append([]string{
				types.DefaultImport,
				types.MuxImport,
			}, r.imports...),
			"Containers": containers,
		}

		filename := filepath.Join(env.GetAbsPath(), "configs", "router")
		env.GetConfiguration().AddAutoFile(filename+"_auto", autoRouterTpl, tplParams)
	}
}

func (r *Router) addMiddlewaresImports(slice types.EntitySlice) {
	for i := 0; i < slice.Len(); i++ {
		r.addImport(slice.Get(i).(types.Middleware).GetEntity().GetContainer().GetImport())
	}
}
