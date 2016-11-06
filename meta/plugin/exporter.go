package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"strings"
)

var (
	exporterAutoTpl = `{{$name := .Name}}` +
		`{{$sourceName := .SourceName}}` +
		`{{$sourceVarName := .SourceVarName}}` +
`package {{.Package}}

import ({{range .AutoImports}}
	"{{.}}"{{end}}
)

type _{{$name}}Impl struct {
	model {{$sourceName}}
	props []_{{$name}}PropertyImpl
	kind types.ProtoKind
}

func New{{$name}}(model {{$sourceName}}) types.Exporter {
	return new{{$name}}(model, {{.VarName}}Properties)
}

func new{{$name}}(model {{$sourceName}}, props []_{{$name}}PropertyImpl) types.Exporter {
	exp := new(_{{$name}}Impl)
	exp.props = props
	exp.model = model
	exp.kind = types.ProtoModelKind
	return exp
}

func ({{.ShortName}} _{{$name}}Impl) Len() int {
	return len({{.ShortName}}.props)
}

func ({{.ShortName}} _{{$name}}Impl) Get(i int) types.ExportableProperty {
	return {{.ShortName}}.props[i]
}

func ({{.ShortName}} _{{$name}}Impl) GetProtoKind() types.ProtoKind {
	return {{.ShortName}}.kind
}

func ({{.ShortName}} _{{$name}}Impl) Export(i int, encoder types.Encoder) {
	{{.ShortName}}.props[i].closure({{.ShortName}}.model, encoder)
}

type _{{.MultipleName}}Impl struct {
	{{.ExportablesVarName}} {{.ExportablesName}}
}

func New{{.MultipleName}}({{.ExportablesVarName}} {{.ExportablesName}}) types.SliceExporter {
	exp := new(_{{.MultipleName}}Impl)
	exp.{{.ExportablesVarName}} = {{.ExportablesVarName}}
	return exp
}

func ({{.ShortName}} _{{.MultipleName}}Impl) Len() int {
	return {{.ShortName}}.{{.ExportablesVarName}}.Len()
}

func ({{.ShortName}} _{{.MultipleName}}Impl) ExportItem(i int, encoder types.Encoder) {
	encoder.One(new{{$name}}({{.ShortName}}.{{.ExportablesVarName}}[i], {{.VarName}}Properties))
}

type _{{$name}}PropertyImpl struct {
	exporter.PropertyImpl
	closure func({{$sourceName}}, types.Encoder)
}

func new{{$name}}Property(name string, kind types.ProtoKind, closure func({{$sourceName}}, types.Encoder)) _{{$name}}PropertyImpl {
	return _{{$name}}PropertyImpl{
		exporter.NewPropertyWithKind(name, kind),
		closure,
	}
}`
exporterTpl = `{{$name := .Name}}` +
`{{$sourceName := .SourceName}}` +
`{{$sourceVarName := .SourceVarName}}` +
`package {{.Package}}

import ({{range .Imports}}
	"{{.}}"{{end}}
)

var (
    {{.VarName}}Properties = []_{{$name}}PropertyImpl{ {{range .Properties}}
		new{{$name}}Property("{{.GetAliasName}}", {{if .HasGetter}}types.{{.GetProtoKind}}{{else}}types.ProtoUnkownKind{{end}}, func({{$sourceVarName}} {{$sourceName}}, encoder types.Encoder) {
			{{if .HasGetter}}encoder.{{.GetMethod}}({{$sourceVarName}}.{{.GetGetterName}}()){{end}}
		}),{{end}}
	}
)`
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
									prop.SetProperty(modelProp)
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
					"AutoImports": append([]string{
						types.DefaultImport,
						types.ExporterImport,
					}, expEntity.GetImports()...),
					"Imports": append([]string{
						types.DefaultImport,
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

				env.GetConfiguration().AddAutoFile(expEntity.GetAutoFilename(), exporterAutoTpl, tplParams)
				env.GetConfiguration().AddFile(expEntity.GetFilename(), exporterTpl, tplParams)
			}
		}
	}
}
