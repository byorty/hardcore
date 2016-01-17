package plugin

import (
	"github.com/byorty/hardcore/meta/types"
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
    {{.Name}}{{end}}
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

import ({{range .AutoImports}}
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

func ({{$shortName}} {{$name}}{{.GetName}}) Call(ctrl interface{}, scope types.RequestScope) { {{range .GetParams}}
	var {{.GetName}} {{.GetDefineVarKind}}{{end}}{{range .GetParams}}
	{{.GetName}}Prim := prim.{{.GetPrimitive}}("{{.GetName}}"){{if .IsRequired}}
	{{.GetName}}Prim.Required(){{end}}
	{{.GetName}}Prim.Export(&{{.GetName}})
	{{end}}
	form := form.New(){{range .GetParams}}
	form.Add({{.GetName}}Prim){{end}}

//		Check(scope)

	{{$shortName}}(ctrl.(*{{$name}}), {{.GetDefineVars}})
//	scope.GetWriter().Write([]byte(result))
}
{{end}}{{end}}
`
)

type Controller struct {}

func (c *Controller) Do(env types.Environment) {
//	logger := env.GetLogger()
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
				autoImports := imports
				if hasForm {
					autoImports = append(autoImports, "github.com/byorty/hardcore/form")
					autoImports = append(autoImports, "github.com/byorty/hardcore/form/prim")
				}
				tplParams := map[string]interface{}{
					"ShortName": strings.ToLower(controllerEntity.GetName()[0:1]),
					"Name": controllerEntity.GetName(),
					"Package": container.GetShortPackage(),
					"Imports": imports,
					"AutoImports": autoImports,
					"Actions": controllerEntity.GetActions(),
				}
//				logger.Info(controllerEntity.GetName(), append([]string{
//					types.DefaultImport,
//				}, controllerEntity.GetImports()...))

				env.GetConfiguration().AddAutoFile(controllerEntity.GetAutoFilename(), autoControllerTpl, tplParams)
				env.GetConfiguration().AddFile(controllerEntity.GetFilename(), controllerTpl, tplParams)
			}
		}
	}
}