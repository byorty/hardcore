package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"strings"
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
	{{.AutoDaoName}}
}

func ({{.ShortName}} *{{.DaoName}}) Init(db types.DB) {}
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
func ({{$shortName}} {{$name}}) Get{{.GetUpperName}}() {{.GetMethodDefineKind}} { {{if .GetRelation.IsOneToOne}}
	if {{$shortName}}.{{.GetName}} == nil {
		var {{.GetName}} {{.GetVariableKind}}
		{{.GetName}}.DAO().ById({{$shortName}}.{{.GetName}}Id).One(&{{.GetName}})
		{{$shortName}}.{{.GetName}} = &{{.GetName}}
	}
	return {{if .GetEntity.GetEntityKind.IsEnum}}*({{$shortName}}.{{.GetName}}){{else}}{{$shortName}}.{{.GetName}}{{end}}
{{else if .GetRelation.IsOneToMany}}
	if {{$shortName}}.{{.GetName}} == nil { {{if .GetEntity.GetEntityKind.IsModel}}
		dao.New{{$upperIdentifierKind}}OneToMany("{{.GetRelationProperty.GetName}}").ById({{$shortName}}.GetId()).All(&{{$shortName}}.{{.GetName}}){{else}}
		{{.GetName}}.DAO().ByIds({{$shortName}}.{{.GetName}}Id).All(&{{$shortName}}.{{.GetName}}){{end}}
	}
	return {{$shortName}}.{{.GetName}}
{{else}}
	return {{$shortName}}.{{.GetName}}
{{end}}}

func ({{$shortName}} *{{$name}}) Set{{.GetUpperName}}({{.GetName}} {{.GetMethodDefineKind}}) *{{$name}} {
	{{if .GetRelation.IsOneToOne}}{{$shortName}}.{{.GetName}} = {{if .GetEntity.GetEntityKind.IsEnum}}&{{end}}{{.GetName}}
	{{$shortName}}.Set{{.GetUpperName}}Id({{.GetName}}.GetId()){{else}}{{$shortName}}.{{.GetName}} = {{.GetName}}{{end}}
	return {{$shortName}}
}
{{end}}
func({{.ShortName}} *{{.Name}}) CommonDAO() types.ModelDAO {
	return {{.ShortName}}.DAO()
}

func({{.ShortName}} *{{.Name}}) KindDAO() types.{{$upperIdentifierKind}}ModelDAO {
	return {{.ShortName}}.DAO()
}

func({{.ShortName}} *{{.Name}}) DAO() *{{.DaoName}} {
	return {{.DaoName}}Inst()
}

func ({{.ShortName}} *{{.Name}}) Proto() types.Proto {
	return {{.VarProtoName}}
}

func ({{.ShortName}} {{.Name}}) IsScanned() bool {
	return {{.ShortName}}.GetId() != 0
}

func ({{.ShortName}} {{.MultipleName}}) Len() int {
	return len({{.ShortName}})
}

func ({{.ShortName}} {{.MultipleName}}) Less(x, y int) bool {
	return {{.ShortName}}[x].GetId() < {{.ShortName}}[y].GetId()
}

func ({{.ShortName}} {{.MultipleName}}) Swap(x, y int) {
	{{.ShortName}}[x], {{.ShortName}}[y] = {{.ShortName}}[y], {{.ShortName}}[x]
}

func ({{.ShortName}} {{.MultipleName}}) GetRaw(x int) interface{} {
	return {{.ShortName}}.Get(x)
}

func ({{.ShortName}} {{.MultipleName}}) Get(x int) *{{.Name}} {
	return {{.ShortName}}[x]
}

func({{.ShortName}} *{{.MultipleName}}) CommonDAO() types.ModelDAO {
	return {{.ShortName}}.DAO()
}

func({{.ShortName}} *{{.MultipleName}}) KindDAO() types.{{$upperIdentifierKind}}ModelDAO {
	return {{.ShortName}}.DAO()
}

func({{.ShortName}} *{{.MultipleName}}) DAO() *{{.DaoName}} {
	return {{.DaoName}}Inst()
}

func ({{.ShortName}} *{{.MultipleName}}) Proto() types.Proto {
	return {{.VarProtoName}}
}

func ({{.ShortName}} {{.MultipleName}}) IsScanned() bool {
	return {{.ShortName}}.Len() > 0 && {{.ShortName}}.Get(0).GetId() != 0
}

type {{.AutoDaoName}} struct {
	dao.{{$upperIdentifierKind}}Impl
}

func {{.DaoName}}Inst() *{{.DaoName}} {
	if {{.VarDaoName}} == nil {
		{{.VarDaoName}} = new({{.DaoName}})
	}
	return {{.VarDaoName}}
}

func ({{.ShortName}} {{.DaoName}}) GetDB() string {
	return "{{.Source}}"
}

func ({{.ShortName}} {{.DaoName}}) GetTable() string {
	return "{{.Table}}"
}

func ({{.ShortName}} {{.DaoName}}) Proto() types.Proto {
	return {{.VarProtoName}}
}

func ({{.ShortName}} {{.DaoName}}) ScanAll(rows interface{}, model interface{}) error {
	var err error
	items := model.(*{{.MultipleName}})
	item := new({{.Name}})
	err = {{.ShortName}}.Scan(rows, item)
	(*items) = append((*items), item)
	return err
}

func ({{.ShortName}} {{.DaoName}}) Scan(row interface{}, model interface{}) error {
	item := model.(*{{.Name}})
	return row.(types.DBScanner).Scan({{range .Properties}}{{if .GetRelation.IsNone}}
		&item.{{.GetName}},{{end}}{{end}}
	)
}

func ({{.ShortName}} *{{.DaoName}}) Add(model *{{.Name}}) {
	db := pool.DB().ByDAO({{.ShortName}})
	if db.SupportLastInsertId() {
		{{.ShortName}}.InsertStmt.Exec({{range .Properties}}{{if and .NotIdentifier .GetRelation.IsNone}}
			model.{{.GetName}},{{end}}{{end}}
		).One(model)
	} else if db.SupportReturningId() {
		{{.ShortName}}.InsertStmt.Custom({{range .Properties}}{{if and .NotIdentifier .GetRelation.IsNone}}
			model.{{.GetName}},{{end}}{{end}}
		).One(&model.id)
	}
}

func ({{.ShortName}} *{{.DaoName}}) Save(model *{{.Name}}) {
	{{.ShortName}}.UpdateStmt.Exec({{range .Properties}}{{if and .NotIdentifier .GetRelation.IsNone}}
		model.{{.GetName}},{{end}}{{end}}
		model.id,
	).One(model)
}

func ({{.ShortName}} *{{.DaoName}}) Take(model *{{.Name}}) {
	if model.IsScanned() {
		 {{.ShortName}}.Save(model)
	} else {
		 {{.ShortName}}.Add(model)
	}
}

func ({{.ShortName}} *{{.DaoName}}) AutoInit(db types.DB) {
	{{.ShortName}}.ByIdStmt = db.Prepare(criteria.SelectByDAO({{.ShortName}}).And(expr.Eq("id", nil)))
	//{{.ShortName}}.ByIdsStmt = db.Prepare(criteria.SelectByDAO({{.ShortName}}).And(expr.In("id", nil)))
	{{.ShortName}}.InsertStmt = db.Prepare(criteria.InsertByDao({{.ShortName}}))
	{{.ShortName}}.UpdateStmt = db.Prepare(criteria.UpdateByDAO({{.ShortName}}).And(expr.Eq("id", nil)))
}
{{range .Properties}}
func {{$varName}}{{.GetUpperName}}Setter(model interface{}, {{.GetName}} interface{}) {
	model.(*{{$name}}).Set{{.GetUpperName}}({{.GetName}}.({{.GetMethodDefineKind}}))
}

func {{$varName}}{{.GetUpperName}}Getter(model interface{}) interface{} {
	return model.(*{{$name}}).Get{{.GetUpperName}}()
}
{{end}}
var (
	{{.VarDaoName}} *{{.DaoName}}
	{{.VarProtoName}} = proto.New().{{range $i, $property := .Properties}}
		Set("{{.GetName}}", proto.NewProperty("{{.GetField}}", types.{{.GetProtoKind}}, types.{{.GetRelation.AsProtoRelation}}, {{.IsRequired}}, {{$varName}}{{.GetUpperName}}Setter, {{$varName}}{{.GetUpperName}}Getter)){{if lt $i $lastPropertyIndex}}.{{end}}{{end}}
)
`

	valueTpl = `package {{.Package}}

type {{.Name}} struct {
	{{.AutoName}}
}

type {{.MultipleName}} []*{{.Name}}

// implement me!
func ({{.ShortName}} {{.MultipleName}}) Less(x, y int) bool {
	return true
}
`
	autoValueTpl = `{{$name := .Name}}` +
		`{{$shortName := .ShortName}}` +
		`package {{.Package}}

type {{.AutoName}} struct {` +
		`
{{range .Properties}}` +
		`	{{.GetName}} {{.GetDefineKind}}
{{end}}}
{{range .Properties}}
func ({{$shortName}} {{$name}}) Get{{.GetUpperName}}() {{.GetMethodDefineKind}} { {{if .GetRelation.IsOneToOne}}
	return {{if .GetEntity.GetEntityKind.IsEnum}}*({{$shortName}}.{{.GetName}}){{else}}{{$shortName}}.{{.GetName}}{{end}}{{else}}return {{$shortName}}.{{.GetName}}{{end}}
}

func ({{$shortName}} *{{$name}}) Set{{.GetUpperName}}({{.GetName}} {{.GetMethodDefineKind}}) *{{$name}} {
	{{$shortName}}.{{.GetName}} = {{if .GetRelation.IsOneToOne}}{{if .GetEntity.GetEntityKind.IsEnum}}&({{$shortName}}.{{.GetName}}){{end}}{{else}}{{.GetName}}{{end}}
	return {{$shortName}}
}{{end}}

func ({{.ShortName}} {{.MultipleName}}) Len() int {
	return len({{.ShortName}})
}

func ({{.ShortName}} {{.MultipleName}}) Swap(x, y int) {
	{{.ShortName}}[x], {{.ShortName}}[y] = {{.ShortName}}[y], {{.ShortName}}[x]
}

func ({{.ShortName}} {{.MultipleName}}) GetRaw(x int) interface{} {
	return {{.ShortName}}.Get(x)
}

func ({{.ShortName}} {{.MultipleName}}) Get(x int) *{{.Name}} {
	return {{.ShortName}}[x]
}
`
)

type Model struct{}

func (m *Model) Do(env types.Environment) {
	config := env.GetConfiguration()
	for _, container := range config.GetContainers() {
		if container.GetContainerKind() == types.EntityContainerKind {
			for _, entity := range container.GetEntities() {
				if entity.GetEntityKind() == types.ModelEntityKind {

					modelEntity := entity.(types.ModelEntity)

					if modelEntity.GetPattern() == types.StraightMappingPattern {
						for _, property := range modelEntity.GetProperties() {
							relation := property.GetRelation()
							hasMany := relation.IsOneToMany() || relation.IsManyToMany()
							isModel := property.GetEntity() != nil && property.GetEntity().GetEntityKind() == types.ModelEntityKind
							if hasMany && isModel {
								relModel := property.GetEntity().(types.ModelEntity)
								for _, relProp := range relModel.GetProperties() {
									relPropRelation := relProp.GetRelation()
									hasRelMany := relPropRelation.IsOneToOne() || relPropRelation.IsManyToMany()
									equalKind := entity.GetFullName() == relProp.GetKind()
									if hasRelMany && equalKind {
										property.SetRelationProperty(relProp)
										break
									}
								}
							}
						}
					}

					varName := utils.LowerFirst(entity.GetName())
					tplParams := map[string]interface{}{
						"ShortName":    strings.ToLower(entity.GetName()[0:1]),
						"Name":         entity.GetName(),
						"MultipleName": modelEntity.GetMultipleName(),
						"AutoName":     fmt.Sprintf("Auto%s", entity.GetName()),
						"Package":      container.GetShortPackage(),
						"Properties": modelEntity.GetProperties(),
					}

					var tpl, autoTpl string
					if modelEntity.GetPattern() == types.StraightMappingPattern {
						tpl = modelTpl
						autoTpl = autoModelTpl
						tplParams["AutoImports"] = append([]string{
							types.DefaultImport,
							types.DaoImport,
							types.ProtoImport,
							types.CriteriaImport,
							types.ExprImport,
							types.PoolImport,
						}, modelEntity.GetImports()...)
						tplParams["UpperIdentifierKind"] = utils.UpperFirst(modelEntity.GetProperties()[0].GetKind())
						tplParams["VarName"] = varName
						tplParams["VarDaoName"] = fmt.Sprintf("%sDao", varName)
						tplParams["VarProtoName"] = fmt.Sprintf("%sProto", varName)
						tplParams["LastPropertyIndex"] = len(modelEntity.GetProperties()) - 1
						tplParams["DaoName"] = fmt.Sprintf("%sDao", entity.GetName())
						tplParams["AutoDaoName"] = fmt.Sprintf("Auto%sDao", entity.GetName())
						tplParams["Source"] = modelEntity.GetSource()
						tplParams["Table"] = modelEntity.GetTable()
						tplParams["Imports"] = []string{
							types.DefaultImport,
						}
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
