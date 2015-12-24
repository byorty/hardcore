package plugin

import (
    "github.com/byorty/hardcore/meta"
    "strings"
    "fmt"
    "github.com/byorty/hardcore/utils"
    "github.com/byorty/hardcore/meta/model"
)

var (
    modelTpl = `package {{.Package}}

type {{.Name}} struct {
	{{.AutoName}}
}

type {{.MultipleName}} []*{{.Name}}

type {{.DaoName}} struct {
	dao.Base
}
`
    autoModelTpl = `{{$autoName := .AutoName}}` +
`package {{.Package}}

import ({{range .AutoImports}}
    "{{.}}"{{end}}
)

type {{.AutoName}} struct { {{range .Properties}}
    {{.Name}} {{if and .HasRelation .Relation.NeedMany}}{{end}}{{.Kind}}{{if .HasRelation}}
    {{.Child.Name}} {{.Child.Kind}}{{end}}
{{end}}}

func(a *{{.AutoName}}) DAO() types.DAO {
	return {{.VarDaoName}}
}

func (a *{{.AutoName}}) Proto() types.Proto {
	return {{.VarProtoName}}
}

{{range .Properties}}
func (a {{$autoName}}) Set{{.UpperName}}({{.Name}} {{.Kind}}) {
    a.{{.Name}} = {{.Name}}
}{{if .HasRelation}}

func (a {{$autoName}}) Set{{.Child.UpperName}}({{.Child.Name}} {{.Child.Kind}}) {
    a.{{.Child.Name}} = {{.Child.Name}}
}
{{end}}

func (a {{$autoName}}) Get{{.UpperName}}() {{.Kind}} {
{{if .HasRelation}}
    {{if .Relation.IsOneToOne}}
    if a.{{.Name}} == nil {
        {{if .IsEnum}}
        var {{.Name}} {{.Kind}}
        a.{{.Name}} = &{{.Name}}
        {{else}}
        a.{{.Name}} = new({{.Kind}})
        {{end}}
        a.{{.Name}}.DAO().ById(a.{{.Child.Name}}).One(a.{{.Name}})
    }
    {{else if .Relation.IsOneToMany}}
        {{if .IsEnum}}
    if a.{{.Name}} == nil {
        var {{.Name}} {{.Kind}}
        {{.Name}}.DAO().ByIds(a.{{.Child.Name}}).All({{.Name}})
        a.{{.Name}} = &{{.Name}}
    }
        {{else}}
    if a.{{.Name}} == nil {

    }
        {{end}}
    {{else if .Relation.IsManyToMany}}
    {{end}}
{{end}}
    return a.{{.Name}}
}

{{end}}

func({{.ShortName}} *{{.MultipleName}}) DAO() types.DAO {
	return {{.VarDaoName}}
}

func ({{.ShortName}} *{{.MultipleName}}) Proto() types.Proto {
	return {{.VarProtoName}}
}

func ({{.ShortName}} {{.DaoName}}) Proto() types.Proto {
	return {{.VarProtoName}}
}

func ({{.ShortName}} {{.DaoName}}) GetDB() string {
	return "{{.DBName}}"
}

func ({{.ShortName}} {{.DaoName}}) GetTable() string {
	return "{{.TableName}}"
}

func ({{.ShortName}} {{.DaoName}}) ScanAll(rows interface{}, rawModels interface{}) {
	models := rawModels.(*{{.MultipleName}})
	model := new({{.Name}})
	{{.ShortName}}.Scan(rows, model)
	(*models) = append((*models), model)
}

func ({{.ShortName}} {{.DaoName}}) Scan(row interface{}, rawModel interface{}) {
	model := rawModel.(*{{.Name}})
	err := row.(types.SqlModelScanner).Scan({{range .Properties}}{{if .HasRelation}}
	    &model.{{.Child.Name}},
	{{else}}
	    &model.{{.Name}},
    {{end}}{{end}})
	if err != nil {
		fmt.Println(err)
	}
}
`
)

type Model struct {}

func (e *Model) Do(env *meta.Environment) {
    for _, container := range env.Configuration.ModelContainers {
        for _, m := range container.Models {

            for _, prop := range m.Properties {
                prop.Child = new(model.Property)
//                env.Logger.Info(prop)
            }

            varName := utils.LowerFirst(m.Name)
            tmplParams := map[string]interface{}{
                "ShortName": strings.ToLower(m.Name[0:1]),
                "Name": m.Name,
                "AutoName": fmt.Sprintf("Auto%s", m.Name),
                "DaoName": fmt.Sprintf("%sDao", m.Name),
                "Package": container.ShortPackage,
                "AutoImports": m.Imports,
                "VarDaoName": fmt.Sprintf("%sDao", varName),
                "VarProtoName": fmt.Sprintf("%sProto", varName),
                "Properties": m.Properties,
            }

            env.Configuration.AddAutoFile(m.AutoFilename, autoModelTpl, tmplParams)
            env.Configuration.AddFile(m.Filename, modelTpl, tmplParams)
        }
    }
}
