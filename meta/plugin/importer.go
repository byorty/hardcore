package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"strings"
)

var (
	importerAutoTpl = `package {{.Package}}

import ({{range .Imports}}
	"{{.}}"{{end}}
)

type _{{.Name}}Impl struct {
	model {{.SourceName}}
	props map[string]_{{.Name}}PropertyImpl
}

func ({{.ShortName}} _{{.Name}}Impl) Get(key string) (types.ImportableProperty, bool) {
	prop, ok := {{.ShortName}}.props[key]
	return prop, ok
}

func ({{.ShortName}} _{{.Name}}Impl) Decode(key string, decoder types.Decoder, value []byte) {
	{{.ShortName}}.props[key].closure({{.ShortName}}.model, decoder, value)
}

type _{{.Name}}PropertyImpl struct {
	kind    types.ProtoKind
	closure func({{.SourceName}}, types.Decoder, []byte)
}

func ({{.ShortName}} _{{.Name}}PropertyImpl) GetProtoKind() types.ProtoKind {
	return {{.ShortName}}.kind
}

func new{{.Name}}Property(kind types.ProtoKind, closure func({{.SourceName}}, types.Decoder, []byte)) _{{.Name}}PropertyImpl {
	return _{{.Name}}PropertyImpl{
		kind,
		closure,
	}
}

func New{{.Name}}({{.SourceVarName}} {{.SourceName}}) types.Importer {
	imp := new(_{{.Name}}Impl)
	imp.model = {{.SourceVarName}}
	imp.props = _{{.VarName}}Properties
	return imp
}
`
	importerTpl = `{{$name := .Name}}` +
		`{{$sourceName := .SourceName}}` +
		`{{$sourceVarName := .SourceVarName}}` +
		`package {{.Package}}

import ({{range .Imports}}
	"{{.}}"{{end}}
)

var (
	_{{.VarName}}Properties = map[string]_{{$name}}PropertyImpl{ {{range .Properties}}
		"{{.GetAliasName}}": new{{$name}}Property(types.{{.GetModelProperty.GetProtoKind}}, func({{$sourceVarName}} {{$sourceName}}, decoder types.Decoder, value []byte) { {{if .HasModelProperty}}
			{{if .GetModelProperty.HasRelation}}{{if .GetModelProperty.GetEntity.GetEntityKind.IsEnum}}var {{.GetModelProperty.GetName}} {{.GetModelProperty.GetEntity.GetFullName}}
			{{.GetModelProperty.GetName}}.DAO().ById(decoder.{{.GetMethod}}(value)).One(&{{.GetModelProperty.GetName}})
			{{$sourceVarName}}.{{.GetSetterName}}({{.GetModelProperty.GetName}}){{end}}{{else}}{{$sourceVarName}}.{{.GetSetterName}}(decoder.{{.GetMethod}}(value)){{end}}
		{{end}}}),{{end}}
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

				env.GetConfiguration().AddFile(impEntity.GetAutoFilename(), importerAutoTpl, tplParams)
				env.GetConfiguration().AddFile(impEntity.GetFilename(), importerTpl, tplParams)
			}
		}
	}
}
