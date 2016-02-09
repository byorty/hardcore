package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"strings"
)

var (
	tpl = `{{$name := .Name}}` +
`{{$sourceName := .SourceName}}` +
`{{$sourceVarName := .SourceVarName}}` +
`package {{.Package}}

import ({{range .Imports}}
	"{{.}}"{{end}}
)

type {{$name}}PropertyImpl struct {
	exporter.PropertyImpl
	closure func({{$sourceName}}) interface{}
}

func New{{$name}}Property(name string, closure func({{$sourceVarName}} {{$sourceName}}) interface{}) types.ExportableProperty {
	return &{{$name}}PropertyImpl{
		exporter.NewProperty(name),
		closure,
	}
}

func ({{.ShortName}} {{$name}}PropertyImpl) GetValue(model interface{}) interface{} {
	return {{.ShortName}}.closure(model.({{$sourceName}}))
}

func New{{$name}}({{.ExportableVarName}} {{.ExportableName}}) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties({{.VarName}}Properties)
	exp.SetExportable({{.ExportableVarName}})
	return exp
}

var (
	{{.VarName}}Properties = []types.ExportableProperty{ {{range .Properties}}
		New{{$name}}Property("{{.GetName}}", func({{$sourceVarName}} {{$sourceName}}) interface{} {
			return {{if .HasGetter}}{{$sourceVarName}}.{{.GetGetterName}}(){{else}}nil{{end}}
		}),{{end}}
	}
)
`
)

type Exporter struct {}

func (e *Exporter) Do(env types.Environment) {
	for _, container := range env.GetConfiguration().GetContainers() {
		if container.GetContainerKind() == types.ExporterContainerKind {
			for _, entity := range container.GetEntities() {

				expEntity := entity.(types.ExporterEntity)
				entity := expEntity.GetSourceEntity()
				var exportableName, exportableVarName string
				if expEntity.GetSource() == entity.GetFullMultipleName() {
					exportableName = entity.GetFullMultipleName()
					exportableVarName = utils.LowerFirst(entity.GetMultipleName())
				} else {
					exportableName = entity.GetPointerFullName()
					exportableVarName = utils.LowerFirst(entity.GetName())
				}
				tplParams := map[string]interface{}{
					"Name": expEntity.GetName(),
					"ShortName": strings.ToLower(expEntity.GetName()[0:1]),
					"Package": container.GetShortPackage(),
					"Imports": append([]string{
						types.DefaultImport,
						types.ExporterImport,
					}, expEntity.GetImports()...),
					"Properties": expEntity.GetProperties(),
					"VarName": utils.LowerFirst(expEntity.GetName()),
					"ExportableName": exportableName,
					"ExportableVarName": exportableVarName,
					"SourceName": entity.GetPointerFullName(),
					"SourceVarName": utils.LowerFirst(entity.GetName()),
				}

				env.GetConfiguration().AddFile(expEntity.GetFilename(), tpl, tplParams)
			}
		}
	}
}
