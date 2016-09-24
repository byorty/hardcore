package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"sort"
	"strings"
)

var (
	controllerTpl = `{{$name := .Name}}` +
		`{{$shortName := .ShortName}}` +
		`package {{.Package}}

import ({{range .Imports}}
    "{{.}}"{{end}}
)

type {{.Name}} struct { {{range .Extends}}
    {{.GetFullName}}{{end}}
}

func New{{.Name}}() types.ActionController {
    return nil
}
{{range .Actions}}
func ({{$shortName}} *{{$name}}) {{.GetName}}({{.GetDefineParams}}){{if .HasReturn}} {{.GetReturn}}{{end}} {

	{{if .HasReturn}}return nil{{end}}
}
{{end}}
`
	autoControllerTpl = `{{$name := .Name}}` +
		`{{$shortName := .ShortName}}` +
		`package {{.Package}}

import ({{range .AutoImports}}
    "{{.}}"{{end}}
)

func ({{.ShortName}} *{{.Name}}) CallAction(action interface{}, rs types.RequestScope) { {{if .Kind.IsWebsocket}}
	action.(func(*{{.Name}}, types.WebsocketScope))({{.ShortName}}, rs.(types.WebsocketScope)){{else}}
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call({{.ShortName}}, rs)
	} else {
		v := action.(func(*{{.Name}}, types.RequestScope) types.View)({{.ShortName}}, rs)
		v.SetController({{.ShortName}})
		v.SetScope(rs)
		v.Render()
	}{{end}}
}
{{range .Actions}}{{if .HasForm}}
type {{$name}}{{.GetName}} func(*{{$name}}, {{.GetDefineKinds}}) {{.GetReturn}}

func ({{$shortName}} {{$name}}{{.GetName}}) Call(rawCtrl interface{}, rs types.RequestScope) {
	form := form.New(){{range .GetParams}}{{if .IsInjection}}{{if .GetInjection.IsMustWrite}}
	{{.GetInjection.GetBody}}{{end}}{{else}}
	var {{.GetName}} {{.GetDefineVarKind}}
	{{.GetName}}Prim := prim.{{.GetPrimitive}}("{{.GetName}}"){{if .IsRequired}}
	{{.GetName}}Prim.Required(){{end}}
	{{.GetName}}Prim.SetSource({{.GetSource}})
	{{.GetName}}Prim.Export(&{{.GetName}})
	form.Add({{.GetName}}Prim){{end}}{{end}}

	var v types.View
	if form.Check(rs) {
		ctrl := rawCtrl.(*{{$name}})
		v = {{$shortName}}(ctrl, {{.GetDefineVars}})
		v.SetController(ctrl)
	} else {
		handler, ok := rawCtrl.(types.FormErrorsHandler)
		if ok {
			v = handler.HandleFormErrors(form.GetErrors())
		} else {
			v = view.BadRequest()
		}
	}
	v.SetScope(rs)
	v.Render()
}
{{end}}{{end}}{{if .HasActionWithForm}}
var ({{range .Actions}}{{if .HasForm}}
	{{$name}}{{.GetName}}Action {{$name}}{{.GetName}} = (*{{$name}}).{{.GetName}}{{end}}{{end}}
)
{{end}}`
)

type Controller struct{}

func (c *Controller) Do(env types.Environment) {
	config := env.GetConfiguration()
	for _, container := range config.GetContainers() {
		if container.GetContainerKind() == types.ControllerContainerKind {
			for _, entity := range container.GetEntities() {
				if entity.GetEntityKind() == types.ControllerEntityKind {
					controllerEntity := entity.(types.ControllerEntity)
					for _, action := range controllerEntity.GetActions() {
						for _, param := range action.GetParams() {
							existsEntity := config.GetEntity(param.GetKind())
							if existsEntity != nil {
								param.SetEntity(existsEntity)
								controllerEntity.AddImport(existsEntity.GetContainer().GetImport())
							}
						}
					}

					var autoInjectionImports sort.StringSlice
					hasForm := false
					for _, action := range controllerEntity.GetActions() {
						if action.HasForm() {
							hasForm = true
							for _, param := range action.GetParams() {
								if param.IsInjection() {
									for _, newImport := range param.GetInjection().GetAutoImports() {
										autoInjectionImports.Sort()
										i := autoInjectionImports.Search(newImport)
										if i == autoInjectionImports.Len() {
											autoInjectionImports = append(autoInjectionImports, newImport)
										}
									}
								}
							}
							break
						}
					}

					imports := append(
						[]string{types.DefaultImport},
						controllerEntity.GetImports()...,
					)
					autoImports := []string{types.DefaultImport}
					if hasForm {
						autoImports = append(
							autoImports,
							"github.com/byorty/hardcore/form",
							"github.com/byorty/hardcore/form/prim",
							"github.com/byorty/hardcore/view",
						)
						autoImports = append(
							autoImports,
							controllerEntity.GetImports()...,
						)
						autoImports = append(
							autoImports,
							autoInjectionImports...,
						)
					}
					tplParams := map[string]interface{}{
						"Extends":           controllerEntity.GetExtends(),
						"ShortName":         strings.ToLower(controllerEntity.GetName()[0:1]),
						"Name":              controllerEntity.GetName(),
						"Package":           container.GetShortPackage(),
						"Imports":           imports,
						"AutoImports":       autoImports,
						"Actions":           controllerEntity.GetActions(),
						"HasActionWithForm": hasForm,
						"Kind":              controllerEntity.GetKind(),
					}

					env.GetConfiguration().AddAutoFile(controllerEntity.GetAutoFilename(), autoControllerTpl, tplParams)
					env.GetConfiguration().AddFile(controllerEntity.GetFilename(), controllerTpl, tplParams)
				}
			}
		}
	}
}
