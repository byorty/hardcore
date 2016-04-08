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

func new{{$name}}Property(name string, closure func({{$sourceVarName}} {{$sourceName}}) interface{}) types.ExportableProperty {
	return &{{$name}}PropertyImpl{
		exporter.NewProperty(name),
		closure,
	}
}

func ({{.ShortName}} {{$name}}PropertyImpl) GetValue(model interface{}) interface{} {
	return {{.ShortName}}.closure(model.({{$sourceName}}))
}
{{if .IsMutiple}}
func New{{$name}}({{.ExportableVarName}} {{.ExportableName}}) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties({{.VarName}}Properties)
	exp.SetExportable({{.ExportableVarName}})
	return exp
}
{{else}}
func New{{$name}}({{.ExportableVarName}} {{.ExportableName}}) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties({{.VarName}}Properties)
	exp.SetExportable({{.ExportableVarName}})
	return exp
}

func New{{.MultipleName}}({{.ExportablesVarName}} {{.ExportablesName}}) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties({{.VarName}}Properties)
	exp.SetExportable({{.ExportablesVarName}})
	return exp
}
{{end}}
var (
	{{.VarName}}Properties = []types.ExportableProperty{ {{range .Properties}}
		new{{$name}}Property("{{.GetName}}", func({{$sourceVarName}} {{$sourceName}}) interface{} {
			return {{if .HasGetter}}{{$sourceVarName}}.{{.GetGetterName}}(){{else}}nil{{end}}
		}),{{end}}
	}
)
`
)

type Exporter struct{}

func (e *Exporter) Do(env types.Environment) {
	logger := env.GetLogger()
	config := env.GetConfiguration()
	for _, container := range config.GetContainers() {
		if container.GetContainerKind() == types.ExporterContainerKind {
			for _, entity := range container.GetEntities() {
				expEntity := entity.(types.ExporterEntity)

				srcEntity := config.GetEntity(expEntity.GetSource())
				if srcEntity == nil {
					logger.Error("source %s for %s not found", expEntity.GetSource(), expEntity.GetName())
				} else {
					expEntity.AddImport(srcEntity.GetContainer().GetImport())
					if srcEntity.GetEntityKind() == types.ModelEntityKind {
						modelEntity := srcEntity.(types.ModelEntity)
						for _, prop := range expEntity.GetProperties() {
							for _, modelProp := range modelEntity.GetProperties() {
								if prop.GetName() == modelProp.GetName() && modelProp.GetRelation().IsNone() {
									prop.SetHasGetter(true)
									break
								}
							}
						}
					} else if srcEntity.GetEntityKind() == types.EnumEntityKind {
						for _, prop := range expEntity.GetProperties() {
							if prop.GetName() == "id" || prop.GetName() == "name" {
								prop.SetHasGetter(true)
								break
							}
						}
					}
				}

				var exportableName, exportableVarName string
				isMutiple := expEntity.GetSource() == srcEntity.GetFullMultipleName()
				if isMutiple {
					exportableName = srcEntity.GetFullMultipleName()
					exportableVarName = utils.LowerFirst(srcEntity.GetMultipleName())
				} else {
					exportableName = srcEntity.GetPointerFullName()
					exportableVarName = utils.LowerFirst(srcEntity.GetName())
				}
				tplParams := map[string]interface{}{
					"Name":         expEntity.GetName(),
					"MultipleName": expEntity.GetMultipleName(),
					"ShortName":    strings.ToLower(expEntity.GetName()[0:1]),
					"Package":      container.GetShortPackage(),
					"Imports": append([]string{
						types.DefaultImport,
						types.ExporterImport,
					}, expEntity.GetImports()...),
					"Properties":         expEntity.GetProperties(),
					"VarName":            utils.LowerFirst(expEntity.GetName()),
					"ExportableName":     exportableName,
					"ExportableVarName":  exportableVarName,
					"ExportablesName":    srcEntity.GetFullMultipleName(),
					"ExportablesVarName": utils.LowerFirst(srcEntity.GetMultipleName()),
					"SourceName":         srcEntity.GetPointerFullName(),
					"SourceVarName":      utils.LowerFirst(srcEntity.GetName()),
					"IsMutiple":          isMutiple,
				}

				env.GetConfiguration().AddFile(expEntity.GetFilename(), tpl, tplParams)
			}
		}
	}
}
