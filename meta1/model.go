package meta

import (
	"text/template"
	"bytes"
	"strings"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/types"
	"reflect"
	"regexp"
	"github.com/byorty/hardcore/logger"
	"fmt"
)

type ModelByDb struct {
	dbName string
	models []interface{}
}

type ormTag map[string]string

func (o ormTag) Required() bool {
	if o["required"] == "true" {
		return true
	} else {
		return false
	}
}

const (
	ProtoBasicKind = "types.ProtoBasicKind"
	ProtoEnumKind  = "types.ProtoEnumKind"
	ProtoModelKind = "types.ProtoModelKind"
)

const (
	ProtoNoneRelation       = "types.ProtoNoneRelation"
	ProtoOneToOneRelation   = "types.ProtoOneToOneRelation"
	ProtoOneToManyRelation  = "types.ProtoOneToManyRelation"
	ProtoManyToManyRelation = "types.ProtoManyToManyRelation"
)

type protoProperty struct {
	Name string
	FieldName string
	Kind string
	Relation string
	Required string
	Proto string
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
    {{$lowerType}}Proto = proto.New(){{range .ProtoProps}}.
        Set({{.}}){{end}}
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
    {{range .EnumFields}}{{.}}.ById({{.}}.Id){{end}}
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
				enumFields := make([]string, 0)
				scanFields := make([]string, 0)
				protoProps := make([]string, 0)
				pointerFields := make([]*Field, 0)
//				protoStruct := proto.New()
				for _, field := range metaModel.Fields {

					tag := field.Tag.Get("orm")
					if tag == "-" {
						continue
					}

//					modelTag := ormTag{
//						"default": "",
//						"required": "false",
//						"": "",
//						"": "",
//						"": "",
//						"": "",
//					}

					fieldName := m.toSqlName(field.Name)

					var kind, relation, required, relationProto string
					required = "true"
					propConstruct := "proto.NewProperty"

					switch field.Kind {
					case reflect.Struct:
						isEnum := m.isEnum(field.PkgPath, field.TypeName)
						if isEnum || m.isModel(field.PkgPath, field.TypeName) {
							fieldName =	fmt.Sprintf("%s_id", fieldName)
							relation = ProtoOneToOneRelation
							if isEnum {
								kind = ProtoEnumKind
								enumFields = append(enumFields, fmt.Sprintf("model.%s", field.Name))
							} else {
								propConstruct = "proto.NewRelationProperty"
								relationProto = fmt.Sprintf("%sProto", m.toFirstLower(field.TypeName))
								kind = ProtoModelKind
							}
							if field.Sign == PtrStructSign {
								pointerFields = append(pointerFields, field)
							}
							scanFields = append(scanFields, fmt.Sprintf("&(model.%s).Id", field.Name))
						} else {
							kind = ProtoBasicKind
							relation = ProtoNoneRelation
							scanFields = append(scanFields, fmt.Sprintf("&model.%s", field.Name))
						}
					case reflect.Slice:

					default:
						kind = ProtoBasicKind
						relation = ProtoNoneRelation
						scanFields = append(scanFields, fmt.Sprintf("&model.%s", field.Name))
					}

					params := []string{fmt.Sprintf(`"%s"`, fieldName), kind, relation, required}
					if len (relationProto) > 0 {
						params = append(params, relationProto)
					}
					protoProps = append(protoProps, fmt.Sprintf(
						`"%s", %s(%s)`,
						field.Name,
						propConstruct,
						strings.Join(params, ", "),
					))
				}

				buf := new(bytes.Buffer)
				tmplParams := map[string]interface{}{
					"ShortStructName": strings.ToLower(metaModel.Name[0:1]),
					"Type"           : metaModel.Name,
					"DaoType"        : metaModel.Name + "DAO",
					"PluralType"     : metaModel.PluralName,
					"LowerType"      : m.toFirstLower(metaModel.Name),
					"ParentDAO"      : ParentDAOByDBKind[currentDb.GetKind()],
					"DataType"       : RawDataTypeByDBKind[currentDb.GetKind()],
					"PointerFields"  : pointerFields,
					"ScanFields"     : strings.Join(scanFields, ", "),
					"EnumFields"     : enumFields,
					"DB"             : modelsByDb.dbName,
					"Table"          : m.toSqlName(metaModel.Name),
					"ProtoProps"     : protoProps,
				}
				tmpl := template.New(metaModel.Name + "Template")
				tmpl.Parse(ModelTmpl)
				tmpl.Execute(buf, tmplParams)
//				fmt.Println(buf.String())
				fmt.Println()
				results = append(results, &BuildResult{
					Pkg: metaModel.Pkg,
					PkgPath: metaModel.PkgPath,
					Bytes: buf.Bytes(),
					Imports: []string{
						ImportByDBKind[currentDb.GetKind()],
						"github.com/byorty/hardcore/proto",
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

func (m *modelBuilder) toFirstLower(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}
