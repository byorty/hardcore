package meta

import (
	"text/template"
	"bytes"
	"strings"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/types"
	"reflect"
	"github.com/byorty/hardcore/proto"
	"regexp"
	"github.com/byorty/hardcore/logger"
	"fmt"
)

type ModelByDb struct {
	dbName string
	models []interface{}
}

var (
	upperRegex = regexp.MustCompile(`([A-Z])`)
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
`{{$proto := .Proto}}` +
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

func ({{$s}} {{.DaoType}}) GetDB() string {
    return "{{.DB}}"
}

func ({{$s}} {{.DaoType}}) GetTable() string {
    return "{{.Table}}"
}

func ({{$s}} {{.DaoType}}) Scan(rawData interface{}, rawModel interface{}) {
    data := rawData.({{.DataType}})
    model := rawModel.(*{{.Type}})
    {{range .PointerFields}}model.{{.Name}} = new({{.TypeName}})
    {{end}}data.Scan({{.ScanFields}})
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
			logger.Err("can't receive db '%s'", modelsByDb.dbName)
		}
		for _, model := range modelsByDb.models {
			metaModel := AnalyzeStruct(model)
			if metaModel == nil {
				logger.Err("model should be a struct, '%v' given", model)
			} else {
				scanFields := make([]string, 0)
				pointerFields := make([]*Field, 0)
				protoStruct := proto.New()
				for _, field := range metaModel.Fields {

					tag := field.Tag.Get("orm")
					if tag == "-" {
						continue
					}

					required := false
					fieldName := m.toSqlName(field.Name)

					var kind types.ProtoKind
					var relation types.ProtoRelation

					switch field.Kind {
					case reflect.Struct:
						isEnum := m.isEnum(field.PkgPath, field.TypeName)
						if isEnum || m.isModel(field.PkgPath, field.TypeName) {
							fieldName =	fmt.Sprintf("%s_id", fieldName)
							relation = types.ProtoOneToOneRelation
							if isEnum {
								kind = types.ProtoEnumKind
							} else {
								kind = types.ProtoModelKind
								if field.Sign == PtrStructSign {
									pointerFields = append(pointerFields, field)
								}
							}
							scanFields = append(scanFields, fmt.Sprintf("&(model.%s).Id", field.Name))
						} else {
							kind = types.ProtoBasicKind
							relation = types.ProtoNoneRelation
							scanFields = append(scanFields, fmt.Sprintf("&model.%s", field.Name))
						}
					case reflect.Slice:

					default:
						kind = types.ProtoBasicKind
						relation = types.ProtoNoneRelation
						scanFields = append(scanFields, fmt.Sprintf("&model.%s", field.Name))
					}

					protoStruct.Set(
						field.Name,
						proto.NewProperty(
							fieldName,
							kind,
							relation,
							required,
						),
					)
				}

				buf := new(bytes.Buffer)
				tmplParams := map[string]interface{}{
					"ShortStructName": strings.ToLower(metaModel.Name[0:1]),
					"Type"           : metaModel.Name,
					"DaoType"        : metaModel.Name + "DAO",
					"PluralType"     : metaModel.PluralName,
					"LowerType"      : strings.ToLower(metaModel.Name[0:1]) + metaModel.Name[1:],
					"ParentDAO"      : ParentDAOByDBKind[currentDb.GetKind()],
					"DataType"       : RawDataTypeByDBKind[currentDb.GetKind()],
					"PointerFields"  : pointerFields,
					"ScanFields"     : strings.Join(scanFields, ", "),
					"DB"             : modelsByDb.dbName,
					"Table"          : m.toSqlName(metaModel.Name),
				}
				tmpl := template.New(metaModel.Name + "Template")
				tmpl.Parse(ModelTmpl)
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
	return m.isExists(m.enums, pkgPath, structName)
}

func (m *modelBuilder) isExists(metaStructs []*Struct, pkgPath, structName string) bool {
	isExists := false
	for _, metaStruct := range metaStructs {
		if metaStruct.PkgPath == pkgPath && metaStruct.Name == structName {
			isExists = true
			break
		}
	}
	return isExists
}

func (m *modelBuilder) isModel(pkgPath, structName string) bool {
	return m.isExists(m.models, pkgPath, structName)
}

func (m *modelBuilder) toSqlName(name string) string {
	name = upperRegex.ReplaceAllString(name, "_$1")
	name = strings.TrimLeft(name, "_")
	name = strings.ToLower(name)
	return name
}
