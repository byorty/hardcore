package meta

import (
	"reflect"
	"fmt"
	"strings"
)

type Struct struct {
	Name       string
	PluralName string
	Pkg        string
	PkgPath    string
	fullname   string
//	Value     reflect.Value
//	Type      reflect.Type
	FieldMap   FieldMap
	Fields     Fields
	IdField    *Field
}

type StructMap map[string]*Struct

type FieldSign int

const (
	NoneSign      FieldSign = iota
	PtrStructSign
)

type Field struct {
	Name     string
	Kind     reflect.Kind
	TypeName string
	Tag      reflect.StructTag
	Pkg      string
	PkgPath  string
	Sign     FieldSign
}

func (f *Field) IsId() bool {
	return f.Name == "Id"
}

func (f *Field) IsName() bool {
	return f.Name == "Name"
}

type FieldMap map[string]*Field
type Fields   []*Field

var (
	structMap = make(StructMap)
)

func GetStructFieldMap(data interface{}) FieldMap {
	metaStruct := AnalyzeStruct(data)
	if metaStruct == nil {
		return nil
	} else {
		return metaStruct.FieldMap
	}
}

func getValue(data interface{}) reflect.Value {
	value := reflect.ValueOf(data)
	if isPtr(data) {
		value = reflect.Indirect(value)
	}
	return value
}

func isPtr(data interface{}) bool {
	return reflect.ValueOf(data).Kind() == reflect.Ptr
}

func AnalyzeStruct(data interface{}) *Struct {
	val := getValue(data)
	if val.Kind() == reflect.Struct {
		valType := val.Type()
		name := valType.Name()
		fullname := fmt.Sprintf("%s/%s", valType.PkgPath(), name)
		if metaStruct, ok := structMap[fullname]; ok {
			return metaStruct
		} else {
			parts := strings.Split(valType.PkgPath(), "/")
			pluralName := name
			if endSRegex.MatchString(pluralName) {
				pluralName = pluralName + "es"
			} else if endYRegex.MatchString(pluralName) {
				pluralName = endYRegex.ReplaceAllString(name, "ie") + "s"
			} else {
				pluralName = pluralName + "s"
			}
			metaStruct := &Struct{
				Name: name,
				PluralName: pluralName,
				Pkg: parts[len(parts)-1],
				PkgPath: valType.PkgPath(),
				fullname: fullname,
//				Value: val,
//				Type: valType,
			}
			metaStruct.FieldMap = make(FieldMap)
			metaStruct.Fields = make(Fields, 0)
			for i := 0;i < valType.NumField();i++ {
				structField := valType.Field(i)
				field := &Field{
					Name: structField.Name,
					Tag: structField.Tag,
				}

				fieldType := structField.Type
				if fieldType.Kind() == reflect.Ptr {
					if structField.Type.Elem().Kind() == reflect.Struct {
						fieldType = structField.Type.Elem()
						field.Sign = PtrStructSign
					}
				}
				field.Kind = fieldType.Kind()

				if field.Kind == reflect.Struct {
					parts := strings.Split(fieldType.PkgPath(), "/")
					field.Pkg = parts[len(parts)-1]
					field.PkgPath = fieldType.PkgPath()
					field.TypeName = fieldType.Name()
				}
				if field.IsId() {
					metaStruct.IdField = field
				}
				metaStruct.FieldMap[structField.Name] = field
				metaStruct.Fields = append(metaStruct.Fields, field)
			}
			structMap[fullname] = metaStruct
			return metaStruct
		}
	} else {
		return nil
	}
}
