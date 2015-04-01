package meta

import (
	"reflect"
	"fmt"
	"strings"
)

type Struct struct {
	Name      string
	Pkg       string
	PkgPath   string
	fullname  string
//	Value     reflect.Value
//	Type      reflect.Type
	FieldMap  FieldMap
	Fields    Fields
	IdField   *Field
}

type StructMap map[string]*Struct

type Field struct {
	Name  string
	Kind  reflect.Kind
//	Value reflect.Value
//	Type  reflect.Type
	Tag   reflect.StructTag
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
			metaStruct := &Struct{
				Name: name,
				Pkg: parts[len(parts)-1],
				PkgPath: valType.PkgPath(),
				fullname: fullname,
//				Value: val,
//				Type: valType,
			}
			metaStruct.FieldMap = make(FieldMap)
			metaStruct.Fields = make(Fields, 0)
			for i := 0;i < valType.NumField();i++ {
				fieldType := valType.Field(i)
				field := &Field{
					Name: fieldType.Name,
					Kind: fieldType.Type.Kind(),
					Tag: fieldType.Tag,
				}
				if field.IsId() {
					metaStruct.IdField = field
				}
				metaStruct.FieldMap[fieldType.Name] = field
				metaStruct.Fields = append(metaStruct.Fields, field)
			}
			structMap[fullname] = metaStruct
			return metaStruct
		}
	} else {
		return nil
	}
}
