package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"strings"
	"fmt"
)

var (
	modelTpl = `package {{.Package}}

type {{.Name}} struct {
	{{.AutoName}}
}

type {{.MultipleName}} []*{{.Name}}

type {{.DaoName}} struct {
	{{.AutoDaoName}}
}
`

	autoModelTpl = `{{$name := .Name}}` +
`{{$shortName := .ShortName}}` +
`{{$varName := .VarName}}` +
`{{$lastPropertyIndex := .LastPropertyIndex}}` +
`{{$upperIdentifierKind := .UpperIdentifierKind}}` +
`package {{.Package}}

import ({{range .AutoImports}}
	"{{.}}"{{end}}
)

type {{.AutoName}} struct {` +
`
{{range .Properties}}` +
`	{{.GetName}} {{.GetDefineKind}}
{{end}}}
{{range .Properties}}
func ({{$shortName}} {{$name}}) Get{{.GetUpperName}}() {{.GetDefineKind}} { {{if .GetRelation.IsOneToOne}}
	if {{$shortName}}.{{.GetName}} == nil {
		var {{.GetName}} {{.GetVariableKind}}
		{{.GetName}}.DAO().ById({{$shortName}}.{{.GetName}}Id).One(&{{.GetName}})
		{{$shortName}}.{{.GetName}} = &{{.GetName}}
	}
	return {{$shortName}}.{{.GetName}}
{{else if .GetRelation.IsOneToMany}}
	if {{$shortName}}.{{.GetName}} == nil {
		var {{.GetName}} {{.GetVariableKind}}{{if .GetEntity.GetEntityKind.IsModel}}
		dao.New{{$upperIdentifierKind}}OneToMany("{{.GetRelationProperty.GetName}}").All(&{{.GetName}})
		{{else}}
		{{.GetName}}.DAO().ByIds({{$shortName}}.{{.GetName}}Id).All(&{{.GetName}})
		{{end}}{{$shortName}}.{{.GetName}} = &{{.GetName}}
	}
	return {{$shortName}}.{{.GetName}}
{{else}}
	return {{$shortName}}.{{.GetName}}
{{end}}}

func ({{$shortName}} *{{$name}}) Set{{.GetUpperName}}({{.GetName}} {{.GetDefineKind}}) *{{$name}} {
	{{$shortName}}.{{.GetName}} = {{.GetName}}{{if .GetRelation.IsOneToOne}}
	{{$shortName}}.Set{{.GetUpperName}}Id({{.GetName}}.GetId()){{end}}
	return {{$shortName}}
}{{end}}

func({{.ShortName}} *{{.Name}}) CommonDAO() types.ModelDAO {
	return {{.VarDaoName}}
}

func({{.ShortName}} *{{.Name}}) DAO() {{.DaoName}} {
	return {{.VarDaoName}}
}

func ({{.ShortName}} *{{.Name}}) Proto() types.Proto {
	return {{.VarProtoName}}
}

func ({{.ShortName}} {{.MultipleName}}) Get(i int) *{{.Name}} {
	return {{.ShortName}}[i]
}

func ({{.ShortName}} {{.MultipleName}}) Len() int {
	return len({{.ShortName}})
}

func({{.ShortName}} *{{.MultipleName}}) CommonDAO() types.ModelDAO {
	return {{.VarDaoName}}
}

func({{.ShortName}} *{{.MultipleName}}) DAO() {{.DaoName}} {
	return {{.VarDaoName}}
}

func ({{.ShortName}} *{{.MultipleName}}) Proto() types.Proto {
	return {{.VarProtoName}}
}

type {{.AutoDaoName}} struct {
	dao.{{$upperIdentifierKind}}Impl
}

func ({{.ShortName}} *{{.DaoName}}) Proto() types.Proto {
	return {{.VarProtoName}}
}
{{range .Properties}}
type {{$name}}{{.GetUpperName}}Setter func(*{{$name}}, {{.GetDefineKind}})

func ({{$shortName}} {{$name}}{{.GetUpperName}}Setter) Call(model interface{}, {{.GetName}} interface{}) {
	{{$shortName}}(model.(*{{$name}}), {{.GetName}}.({{.GetDefineKind}}))
}

type {{$name}}{{.GetUpperName}}Getter func(*{{$name}}) {{.GetDefineKind}}

func ({{$shortName}} {{$name}}{{.GetUpperName}}Getter) Call(model interface{}) interface{} {
	return {{$shortName}}(model.(*{{$name}}))
}
{{end}}
var ({{range .Properties}}
	{{$varName}}{{.GetUpperName}}Setter {{$name}}{{.GetUpperName}}Setter = (*{{$name}}).Set{{.GetUpperName}}
	{{$varName}}{{.GetUpperName}}Getter {{$name}}{{.GetUpperName}}Getter = (*{{$name}}).Get{{.GetUpperName}}{{end}}
	{{.VarDaoName}} {{.DaoName}}
	{{.VarProtoName}} = proto.New().{{range $i, $property := .Properties}}
		Set("{{.GetName}}", proto.NewProperty("{{.GetField}}", types.{{.GetProtoKind}}, types.{{.GetRelation.AsProtoRelation}}, {{.IsRequired}}, {{$varName}}{{.GetUpperName}}Setter, {{$varName}}{{.GetUpperName}}Getter)){{if lt $i $lastPropertyIndex}}.{{end}}{{end}}
)
`

	valueTpl = `package {{.Package}}

type {{.Name}} struct {
	{{.AutoName}}
}

type {{.MultipleName}} []*{{.Name}}
`
	autoValueTpl = `{{$name := .Name}}` +
`{{$shortName := .ShortName}}` +
`package {{.Package}}

import ({{range .AutoImports}}
	"{{.}}"{{end}}
)

type {{.AutoName}} struct {` +
`
{{range .Properties}}` +
`	{{.GetName}} {{.GetDefineKind}}
{{end}}}
{{range .Properties}}
func ({{$shortName}} {{$name}}) Get{{.GetUpperName}}() {{.GetDefineKind}} {
	return {{$shortName}}.{{.GetName}}
}

func ({{$shortName}} *{{$name}}) Set{{.GetUpperName}}({{.GetName}} {{.GetDefineKind}}) *{{$name}} {
	{{$shortName}}.{{.GetName}} = {{.GetName}}
	return {{$shortName}}
}{{end}}

func ({{.ShortName}} {{.MultipleName}}) Get(i int) *{{.Name}} {
	return {{.ShortName}}[i]
}

func ({{.ShortName}} {{.MultipleName}}) Len() int {
	return len({{.ShortName}})
}
`
)

type Model struct {}

func (m *Model) Do(env types.Environment) {
	for _, container := range env.GetConfiguration().GetContainers() {
		if container.GetContainerKind() == types.EntityContainerKind {
			for _, entity := range container.GetEntities() {
				if entity.GetEntityKind() == types.ModelEntityKind {

					modelEntity := entity.(types.ModelEntity)
					varName := utils.LowerFirst(entity.GetName())
					tplParams := map[string]interface{}{
						"ShortName": strings.ToLower(entity.GetName()[0:1]),
						"Name": entity.GetName(),
						"MultipleName": modelEntity.GetMultipleName(),
						"AutoName": fmt.Sprintf("Auto%s", entity.GetName()),
						"Package": container.GetShortPackage(),
						"AutoImports": append([]string{
							types.DefaultImport,
							types.DaoImport,
							types.ProtoImport,
						}, entity.GetImports()...),
						"Properties": modelEntity.GetProperties(),
					}

					var tpl, autoTpl string
					if modelEntity.GetPattern() == types.StraightMappingPattern {
						tpl = modelTpl
						autoTpl = autoModelTpl

						tplParams["UpperIdentifierKind"] = utils.UpperFirst(modelEntity.GetProperties()[0].GetKind())
						tplParams["VarName"] = varName
						tplParams["VarDaoName"] = fmt.Sprintf("%sDao", varName)
						tplParams["VarProtoName"] = fmt.Sprintf("%sProto", varName)
						tplParams["LastPropertyIndex"] = len(modelEntity.GetProperties()) - 1
						tplParams["DaoName"] = fmt.Sprintf("%sDao", entity.GetName())
						tplParams["AutoDaoName"] = fmt.Sprintf("Auto%sDao", entity.GetName())
					} else {
						tpl = valueTpl
						autoTpl = autoValueTpl
					}

					env.GetConfiguration().AddAutoFile(modelEntity.GetAutoFilename(), autoTpl, tplParams)
					env.GetConfiguration().AddFile(modelEntity.GetFilename(), tpl, tplParams)
				}
			}
		}
	}
}