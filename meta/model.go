package meta

import (
	"fmt"
	"text/template"
	"bytes"
	"strings"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/types"
	"reflect"
)

type ModelByDb struct {
	dbName string
	models []interface{}
}

var (

	ImportByDBKind = map[types.DBKind]string {
		types.SqlDB: "database/sql",
	}
	RawDataTypeByDBKind = map[types.DBKind]string {
		types.SqlDB: "*sql.Row",
	}
	ParentDAOByDBKind = map[types.DBKind]string {
		types.SqlDB: "dao.SqlDAO",
	}
	ModelTmpl = `{{$s := .ShortStructName}}` +
`{{$lowerType := .LowerType}}` +
`func New{{.Type}}() *{{.Type}} {
    return new({{.Type}})
}

func({{$s}} {{.Type}}) DAO() {{.DaoType}} {
    return {{$lowerType}}DAO
}

type {{.PluralType}} []*{{.Type}}

func({{$s}} {{.PluralType}}) DAO() {{.DaoType}} {
    return {{$lowerType}}DAO
}

var (
    {{$lowerType}}DAO {{.DaoType}}
)

type {{.DaoType}} struct {
    {{.ParentDAO}}
}

func ({{$s}} {{.DaoType}}) Scan(data interface{}, model interface{}) {
	rawData := data.({{.RawDataType}})
	{{$lowerType}} := model.(*{{.Type}})
	rawData.Scan()
}
`
)

type modelBuilder struct {
	enums []*Struct
	models []*Struct
}

func (m *modelBuilder) Build(modelsByDbs ...interface{}) []*BuildResult {
	if len(itemsByKind[enumBuilderKind]) > 0 {
		m.enums = make([]*Struct, 0)
		for _, enum := range itemsByKind[enumBuilderKind] {
			metaEnum := AnalyzeStruct(enum)
			if metaEnum != nil {
				m.enums = append(m.enums, metaEnum)
			}
		}
	}

	m.models = make([]*Struct, 0)
	for _, rawModelsByDb := range modelsByDbs {
		modelsByDb := rawModelsByDb.(ModelByDb)
		for _, model := range modelsByDb.models {
			metaModel := AnalyzeStruct(model)
			if metaModel != nil {
				m.models = append(m.models, metaModel)
			}
		}
	}

	results := make([]*BuildResult, 0)
	for _, rawModelsByDb := range modelsByDbs {
		modelsByDb := rawModelsByDb.(ModelByDb)
		currentDb := db.Pool().ByName(modelsByDb.dbName)
		if currentDb == nil {
			panic(fmt.Sprintf("can't receive db '%s'", modelsByDb.dbName))
		}
		for _, model := range modelsByDb.models {
			metaModel := AnalyzeStruct(model)
			if metaModel == nil {
				panic(fmt.Sprintf("model should be a struct, '%v' given", model))
			} else {
				for _, field := range metaModel.Fields {
					tag := field.Tag.Get("orm")
					if tag == "-" {
						continue
					}
					switch field.Kind {
					case reflect.Struct:

					case reflect.Slice:

					default:
						
					}
				}

				buf := new(bytes.Buffer)
				tmplParams := map[string]interface{}{
					"ShortStructName": strings.ToLower(metaModel.Name[0:1]),
					"Type": metaModel.Name,
					"DaoType": metaModel.Name + "DAO",
					"PluralType": metaModel.PluralName,
					"LowerType": strings.ToLower(metaModel.Name),
					"ParentDAO": ParentDAOByDBKind[currentDb.GetKind()],
					"RawDataType": RawDataTypeByDBKind[currentDb.GetKind()],
				}
				tmpl := template.New(metaModel.Name + "Template")
				tmpl.Parse(ControllerTmpl)
				tmpl.Execute(buf, tmplParams)
				results = append(results, &BuildResult{
					Pkg: metaModel.Pkg,
					PkgPath: metaModel.PkgPath,
					Bytes: buf.Bytes(),
					Imports: []string{
						ImportByDBKind[currentDb.GetKind()],
						"github.com/byorty/hardcore/orm/dao",
					},
				})
			}
		}
	}
	return results
}

func (m *modelBuilder) isEnum(pkgPath, structName string) bool {
	isExists := false
	for _, enum := range m.enums {

	}
	return isExists
}

func (m *modelBuilder) isExists(metaStructs []*Struct, pkgPath, structName string) bool {
	isExists := false
	for _, metaStruct := range metaStructs {

	}
	return isExists
}

func (m *modelBuilder) isModel(pkgPath, structName string) bool {
	isExists := false
	for _, enum := range m.models {

	}
	return isExists
}
