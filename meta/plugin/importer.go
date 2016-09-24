package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"strings"
)

var (
	importerTpl = `{{$name := .Name}}` +
		`{{$sourceName := .SourceName}}` +
		`{{$sourceVarName := .SourceVarName}}` +
		`package {{.Package}}

import ({{range .Imports}}
	"{{.}}"{{end}}
)

type {{$name}}PropertyImpl struct {
	importer.PropertyImpl
	closure func({{$sourceName}}, interface{})
}

func new{{$name}}Property(kind types.ProtoKind, closure func({{$sourceVarName}} {{$sourceName}}, value interface{})) types.ImportableProperty {
	return &{{$name}}PropertyImpl{
		importer.NewProperty(types.ScalarImportablePropertyKind, kind),
		closure,
	}
}

func ({{.ShortName}} {{$name}}PropertyImpl) SetValue(model interface{}, value interface{}) {
	{{.ShortName}}.closure(model.({{$sourceName}}), value)
}

{{if .IsMutiple}}func New{{$name}}({{.ImportableVarName}} {{.ImportableName}}) types.Importer {
	imp := new(importer.BaseImpl)
	imp.SetProperties({{.VarName}}Properties)
	imp.SetImportable({{.ImportableVarName}})
	return imp
}
{{else}}
func New{{$name}}({{.ImportableVarName}} {{.ImportableName}}) types.Importer {
	imp := new(importer.BaseImpl)
	imp.SetProperties({{.VarName}}Properties)
	imp.SetImportable({{.ImportableVarName}})
	return imp
}

func New{{.MultipleName}}({{.ImportablesVarName}} {{.ImportablesName}}) types.Importer {
	imp := new(importer.BaseImpl)
	imp.SetProperties({{.VarName}}Properties)
	imp.SetImportable({{.ImportablesVarName}})
	return imp
}
{{end}}
var (
	{{.VarName}}Properties = types.ImportableProperties{ {{range .Properties}}{{if .HasModelProperty}}
		"{{.GetName}}": new{{$name}}Property(types.{{.GetModelProperty.GetProtoKind}}, func({{$sourceVarName}} {{$sourceName}}, value interface{}) {
			{{$sourceVarName}}.{{.GetSetterName}}(value.({{.GetModelProperty.GetKind}}))
		}),{{end}}{{end}}
	}
)
`
)

type Importer struct{}

func (i *Importer) Do(env types.Environment) {
	logger := env.GetLogger()
	config := env.GetConfiguration()
	for _, container := range config.GetContainers() {
		if container.GetContainerKind() == types.ImporterContainerKind {
			for _, entity := range container.GetEntities() {
				impEntity := entity.(types.ImporterEntity)

				srcEntity := config.GetEntity(impEntity.GetSource())
				if srcEntity == nil {
					logger.Error("source %s for %s not found", impEntity.GetSource(), impEntity.GetName())
				} else {
					impEntity.AddImport(srcEntity.GetContainer().GetImport())
					if srcEntity.GetEntityKind() == types.ModelEntityKind {
						modelEntity := srcEntity.(types.ModelEntity)
						for _, prop := range impEntity.GetProperties() {
							for _, modelProp := range modelEntity.GetProperties() {
								if prop.GetName() == modelProp.GetName() {
									prop.SetModelProperty(modelProp)
									break
								}
							}
						}
					}
				}

				var importableName, importableVarName string
				isMutiple := impEntity.GetSource() == srcEntity.GetFullMultipleName()
				if isMutiple {
					importableName = srcEntity.GetFullMultipleName()
					importableVarName = utils.LowerFirst(srcEntity.GetMultipleName())
				} else {
					importableName = srcEntity.GetPointerFullName()
					importableVarName = utils.LowerFirst(srcEntity.GetName())
				}
				tplParams := map[string]interface{}{
					"Name":         impEntity.GetName(),
					"MultipleName": impEntity.GetMultipleName(),
					"ShortName":    strings.ToLower(impEntity.GetName()[0:1]),
					"Package":      container.GetShortPackage(),
					"Imports": append([]string{
						types.DefaultImport,
						types.ImporterImport,
					}, impEntity.GetImports()...),
					"Properties":         impEntity.GetProperties(),
					"VarName":            utils.LowerFirst(impEntity.GetName()),
					"ImportableName":     importableName,
					"ImportableVarName":  importableVarName,
					"ImportablesName":    srcEntity.GetFullMultipleName(),
					"ImportablesVarName": utils.LowerFirst(srcEntity.GetMultipleName()),
					"SourceName":         srcEntity.GetPointerFullName(),
					"SourceVarName":      utils.LowerFirst(srcEntity.GetName()),
					"IsMutiple":          isMutiple,
				}

				env.GetConfiguration().AddAutoFile(impEntity.GetFilename(), importerTpl, tplParams)
			}
		}
	}
}
