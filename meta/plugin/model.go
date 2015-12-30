package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"strings"
	"fmt"
)

var (
	modelTpl = `package {{.Package}}

import ({{range .Imports}}
	"{{.}}"{{end}}
)

type {{.Name}} struct {
	{{.AutoName}}
}

type {{.MultipleName}} []*{{.Name}}

type {{.DaoName}} struct {
	dao.Base
}
`

	autoModelTpl = `{{$name := .Name}}` +
`{{$shortName := .ShortName}}` +
`package {{.Package}}

type {{.AutoName}} struct {` +
`
{{range .Properties}}` +
`	{{.GetName}} {{.GetDefineKind}}
{{end}}}

{{range .Properties}}

func ({{$shortName}} {{$name}}) Get{{.GetUpperName}}() {{.GetDefineKind}} {
{{if .GetRelation.IsOneToOne}}
	if {{$shortName}}.{{.GetName}} == nil {
		var {{.GetName}} {{.GetVariableKind}}
		{{.GetName}}.DAO().ById({{$shortName}}.{{.GetName}}Id).One(&{{.GetName}})
		{{$shortName}}.{{.GetName}} = &{{.GetName}}
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

func ({{.ShortName}} {{.MultipleName}}) Get(i int) *{{.Name}} {
	return {{.ShortName}}[i]
}

func ({{.ShortName}} {{.MultipleName}}) Len() int {
	return len({{.ShortName}})
}
`
)

type Model struct {

}

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
						"DaoName": fmt.Sprintf("%sDao", entity.GetName()),
						"Package": container.GetShortPackage(),
//						"AutoImports": m.Imports,
						"Imports": []string{
							types.DaoImport,
						},
						"VarDaoName": fmt.Sprintf("%sDao", varName),
						"VarProtoName": fmt.Sprintf("%sProto", varName),
						"Properties": modelEntity.GetProperties(),
					}

					env.GetConfiguration().AddAutoFile(modelEntity.GetAutoFilename(), autoModelTpl, tplParams)
					env.GetConfiguration().AddFile(modelEntity.GetFilename(), modelTpl, tplParams)
				}
			}
		}
	}
}