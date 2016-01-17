package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"strings"
	"github.com/byorty/hardcore/meta/controller"
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
	// implement me
    return nil
}

{{range .Actions}}
func ({{$shortName}} *{{$name}}) {{.GetName}}({{.GetDefineParams}}) {

}
{{end}}
`
	autoControllerTpl = `{{$name := .Name}}` +
`{{$shortName := .ShortName}}` +
`package {{.Package}}

import ({{range .FormImports}}
    "{{.}}"{{end}}
)

func ({{.ShortName}} *{{.Name}}) CallAction(action interface{}, scope types.RequestScope) {
	if callable, ok := action.(types.CallableAction); ok {
		callable.Call({{.ShortName}}, scope)
	} else {
		action.(func(*{{.Name}}, types.RequestScope))({{.ShortName}}, scope)
	}
}

{{range .Actions}}{{if .HasForm}}
type {{$name}}{{.GetName}} func(*{{$name}}, {{.GetDefineKinds}})

func ({{$shortName}} {{$name}}{{.GetName}}) Call(rawCtrl interface{}, scope types.RequestScope) {
	form := form.New()
	{{range .GetParams}}{{if .IsReserved}}
	{{else}}
	var {{.GetName}} {{.GetDefineVarKind}}
	{{.GetName}}Prim := prim.{{.GetPrimitive}}("{{.GetName}}")
	{{if .IsRequired}}{{.GetName}}Prim.Required(){{end}}
	{{.GetName}}Prim.SetSource({{.GetSource}})
	{{.GetName}}Prim.Export(&{{.GetName}})
	{{end}}{{end}}

	form.Check(scope)

	ctrl := rawCtrl.(*{{$name}})
	{{$shortName}}(ctrl, {{.GetDefineVars}})
//	{{$shortName}}(ctrl, New{{$name}}{{.GetName}}Form(ctrl.(*{{$name}})))
}
{{end}}{{end}}
`
//	formTpl = `{{$name := .Name}}` +
//`package {{.Package}}
//
//import ({{range .FormImports}}
//    "{{.}}"{{end}}
//)
//
//{{range .Actions}}{{if .HasForm}}
//func New{{$name}}{{.GetName}}Form(ctrl *{{$name}}) ({{.GetDefineKinds}}) { {{range .GetParams}}
//	var {{.GetName}} {{.GetDefineVarKind}}{{end}}{{range .GetParams}}
//	{{.GetName}}Prim := prim.{{.GetPrimitive}}("{{.GetName}}"){{if .IsRequired}}
//	{{.GetName}}Prim.Required(){{end}}
//	{{.GetName}}Prim.Export(&{{.GetName}})
//	{{end}}
//	form := form.New(){{range .GetParams}}
//	form.Add({{.GetName}}Prim){{end}}
//
//	return {{.GetDefineVars}}
//}
//{{end}}{{end}}
//`
)

type Controller struct {}

func (c *Controller) Do(env types.Environment) {
//	logger := env.GetLogger()

	container := new(controller.Container)
	container.SetShortPackage("mux")
	parent := new(controller.Controller)
	parent.Name = "ControllerImpl"
	parent.SetContainer(container)

	for _, container := range env.GetConfiguration().GetContainers() {
		if container.GetContainerKind() == types.ControllerContainerKind {
			for _, entity := range container.GetEntities() {

				controllerEntity := entity.(types.ControllerEntity)
				hasForm := false
				for _, action := range controllerEntity.GetActions() {
					if action.HasForm() {
						hasForm = true
						break
					}
				}

				imports := append(
					[]string{types.DefaultImport},
					controllerEntity.GetImports()...
				)
				var formImports []string
				if hasForm {
					formImports = append(
						[]string{
							types.DefaultImport,
							"github.com/byorty/hardcore/form",
							"github.com/byorty/hardcore/form/prim",
						},
						controllerEntity.GetImports()...
					)
				}
				tplParams := map[string]interface{}{
					"Extends": controllerEntity.GetExtends(),
					"ShortName": strings.ToLower(controllerEntity.GetName()[0:1]),
					"Name": controllerEntity.GetName(),
					"Package": container.GetShortPackage(),
					"Imports": imports,
					"FormImports": formImports,
					"Actions": controllerEntity.GetActions(),
				}

				env.GetConfiguration().AddAutoFile(controllerEntity.GetAutoFilename(), autoControllerTpl, tplParams)
				env.GetConfiguration().AddFile(controllerEntity.GetFilename(), controllerTpl, tplParams)
//				env.GetConfiguration().AddFile(controllerEntity.GetFilename() + "_form", formTpl, tplParams)
			}
		}
	}
}