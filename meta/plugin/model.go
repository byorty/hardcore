package plugin

import (
    "github.com/byorty/hardcore/meta"
)

var (
    modelTpl = `package {{.Package}}

import ({{range .Imports}}
    "{{.}}"{{end}}
)

type {{.Name}} struct {
	Auto{{.Name}}
}

type {{.MultipleName}} []*{{.Name}}

type {{.Name}}DAO struct {
	dao.Base
}
`
    autoModelTpl = `package {{.Package}}

import ({{range .AutoImports}}
    "{{.}}"{{end}}
)

type {{.AutoName}} struct { {{range .Properties}}
    {{.Name}} {{.Kind}}{{if .HasRelation}}
    {{.RelationName}} {{.RelationKind}}{{end}}
{{end}}}

func(a *{{.AutoName}}) DAO() types.DAO {
	return {{.VarDaoName}}
}

func (a *{{.AutoName}}) Proto() types.Proto {
	return {{.VarProtoName}}
}

{{range .Properties}}
func (a {{.AutoName}}) Set({{.Name}} {{.Kind}}) {
    a.{{.Name}} = {{.Name}}
}{{if .HasRelation}}

func (a {{.AutoName}}) Set({{.RelationName}} {{.RelationKind}}) {
    a.{{.RelationName}} = {{.RelationName}}
}
{{end}}

func (a {{.AutoName}}) Get() {{.Kind}} {
{{if .HasRelation}}
    {{if .Relation.IsOneToOne}}
    if a.{{.Name}} == nil {
        a.{{.Name}} = new({{.Kind}})
        a.{{.Name}}.DAO().ById(a.{{.RelationName}}).One(a.{{.Name}})
    }
    {{else if .Relation.IsOneToMany}}
        {{if .IsEnum}}
    if a.{{.Name}} == nil {
        var {{.Name}} {{.Kind}}
        {{.Name}}.DAO().ByIds(a.{{.RelationName}}).All({{.Name}})
        a.{{.Name}} = &{{.Name}}
    }
        {{else}}
    if a.{{.Name}} == nil {
        a.{{.Name}} = make({{.Kind}}, 0)
        a.{{.Name}}.O2MDAO().
    }
        {{end}}
    {{else if .Relation.ManyOneToMany}}
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

func ({{.ShortName}} {{.VarDaoName}}) GetDB() string {
	return "{{.DBName}}"
}

func ({{.ShortName}} {{.VarDaoName}}) GetTable() string {
	return "{{.TableName}}"
}

func ({{.ShortName}} {{.VarDaoName}}) ScanAll(rows interface{}, rawModels interface{}) {
	models := rawModels.(*{{.MultipleName}})
	model := new({{.Name}})
	{{.ShortName}}.Scan(rows, model)
	(*models) = append((*models), model)
}

func ({{.ShortName}} {{.VarDaoName}}) Scan(row interface{}, rawModel interface{}) {
	model := rawModel.(*{{.Name}})
	err := row.(types.SqlModelScanner).Scan({{range .Properties}}{{if .HasRelation}}
	    &model.{{.RelationName}},
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
        for _, model := range container.Models {

            for _, prop := range model.Properties {
                env.Logger.Info(prop)
            }
        }
    }
}
