package meta

import (
	"fmt"
	"strings"
	"reflect"
	"bytes"
	"github.com/byorty/hardcore"
)

var enumTypes = map[string]string{
		reflect.Int.String()       : "hardcore.Enum",
		reflect.Int8.String()      : "hardcore.Enum8",
		reflect.Int16.String()     : "hardcore.Enum16",
		reflect.Int32.String()     : "hardcore.Enum32",
		reflect.Int64.String()     : "hardcore.Enum64",
		reflect.Uint.String()      : "hardcore.UEnum",
		reflect.Uint8.String()     : "hardcore.UEnum8",
		reflect.Uint16.String()    : "hardcore.UEnum16",
		reflect.Uint32.String()    : "hardcore.UEnum32",
		reflect.Uint64.String()    : "hardcore.UEnum64",
		reflect.Float32.String()   : "hardcore.Float32Enum",
		reflect.Float64.String()   : "hardcore.Float64Enum",
		reflect.Complex64.String() : "hardcore.Complex64Enum",
		reflect.Complex128.String(): "hardcore.Complex128Enum",
		reflect.String.String()    : "hardcore.StringEnum",
	}

func AddEnumType(enumType, interfaceName string) {
	enumTypes[enumType] = interfaceName
}

type enumInst struct {
	Id   string
	Var  string
}

type enumBuilder struct {}

func (e *enumBuilder) Build(enums ...interface{}) []*BuildResult {
	results := make([]*BuildResult, 0)
	for _, enum := range enums {
		info := AnalyzeStruct(enum)
		if info == nil {
			panic(fmt.Sprintf("enum should be a struct, '%v' given", enum))
		} else {
			buf := new(bytes.Buffer)
			var valuesCount int
			initParams := make(map[int]map[int]string)
			for i, field := range info.Fields {
				if len(field.Tag.Get("enum")) > 0 {
					values := strings.Split(field.Tag.Get("enum"), ",")
					if field.IsId() {
						valuesCount = len(values)
					}
					for j, value := range values {
						if _, ok := initParams[j]; !ok {
							initParams[j] = make(map[int]string)
						}
						value = strings.TrimSpace(value)
						if field.Kind == reflect.String {
							value = fmt.Sprintf("\"%s\"", value)
						}
						initParams[j][i] = value
					}
				}
			}
			insts := make([]*enumInst, 0)
			if valuesCount > 0 {
				buf.WriteString("var (")
				buf.WriteRune(hardcore.EOL)
				for i := 0;i < valuesCount;i++ {
					instName := fmt.Sprintf("__autoEnum%sInst%d", info.Name, i)
					params := make([]string, 0)
					for j, field := range info.Fields {
						if initParam, ok := initParams[i][j]; ok {
							params = append(params, fmt.Sprintf("%s: %s", field.Name, initParam))
							if field.IsId() {
								insts = append(insts, &enumInst{Id: initParam, Var: instName})
							}
						}
					}
					buf.WriteString(fmt.Sprintf("    %s = &%s{%s}", instName, info.Name, strings.Join(params, ", ")))
					buf.WriteRune(hardcore.EOL)
				}

				mapName := fmt.Sprintf("__auto%sMap", info.Name)
				mapNameAsEnum := fmt.Sprintf("__auto%sAsEnumMap", info.Name)
				sliceName := fmt.Sprintf("__auto%sSlice", info.Name)
				sliceNameAsEnum := fmt.Sprintf("__auto%sAsEnumSlice", info.Name)
				enumType, ok := enumTypes[info.IdField.Kind.String()]
				if !ok {
					panic(fmt.Sprintf("unknow enum type \"%s\"", info.IdField.Kind.String()))
				}

				e.writeMap(buf, mapName, info.IdField.Kind.String(), info.Name, true, insts)
				e.writeMap(buf, mapNameAsEnum, info.IdField.Kind.String(), enumType, false, insts)
				e.writeSlice(buf, sliceName, info.Name, true, insts)
				e.writeSlice(buf, sliceNameAsEnum, enumType, false, insts)

				buf.WriteString(")")
				buf.WriteRune(hardcore.EOL)
				buf.WriteRune(hardcore.EOL)

				for _, field := range info.Fields {
					if len(field.Tag.Get("enum")) > 0 {
						e.writeGetter(buf, info, field)
					}
				}

				e.writeMapGetter(buf, info.PluralName + "AsMap", info.IdField.Kind.String(), info.Name, mapName, true)
				e.writeMapGetter(buf, info.PluralName + "AsEnumMap", info.IdField.Kind.String(), enumType, mapNameAsEnum, false)
				e.writeSliceGetter(buf, info.PluralName + "AsSlice", info.Name, sliceName, true)
				e.writeSliceGetter(buf, info.PluralName + "AsEnumSlice", enumType, sliceNameAsEnum, false)

				e.writeByIdGetter(buf, info.Name, info.IdField.Kind.String(), info.Name, insts, true)
				e.writeByIdGetter(buf, info.Name + "AsEnum", info.IdField.Kind.String(), enumType, insts, false)

				results = append(results, &BuildResult{Pkg: info.Pkg, PkgPath: info.PkgPath, Bytes: buf.Bytes()})
			}
		}
	}
	return results
}

func (e *enumBuilder) writeMap(buf *bytes.Buffer, name, keyType, mapType string, isPointer bool, insts []*enumInst) {
	buf.WriteString(fmt.Sprintf("    %s = map[%s]", name, keyType))
	if isPointer {
		buf.WriteString("*")
	}
	buf.WriteString(mapType)
	buf.WriteString("{")
	buf.WriteRune(hardcore.EOL)
	for _, inst := range insts {
		buf.WriteString("        ")
		buf.WriteString(inst.Id)
		buf.WriteString(": ")
		buf.WriteString(inst.Var)
		buf.WriteString(",")
		buf.WriteRune(hardcore.EOL)
	}
	buf.WriteString("    }")
	buf.WriteRune(hardcore.EOL)
}

func (e *enumBuilder) writeSlice(buf *bytes.Buffer, name, mapType string, isPointer bool, insts []*enumInst) {
	buf.WriteString(fmt.Sprintf("    %s = []", name))
	if isPointer {
		buf.WriteString("*")
	}
	buf.WriteString(mapType)
	buf.WriteString("{")
	buf.WriteRune(hardcore.EOL)
	for _, inst := range insts {
		buf.WriteString("        ")
		buf.WriteString(inst.Var)
		buf.WriteString(",")
		buf.WriteRune(hardcore.EOL)
	}
	buf.WriteString("    }")
	buf.WriteRune(hardcore.EOL)
}

func (e *enumBuilder) writeMapGetter(buf *bytes.Buffer, name, keyType, returnType, varName string, isPointer bool) {
	buf.WriteString(fmt.Sprintf("func Get%s() map[%s]", name, keyType))
	if isPointer {
		buf.WriteString("*")
	}
	buf.WriteString(fmt.Sprintf("%s {", returnType))
	buf.WriteRune(hardcore.EOL)
	buf.WriteString("    return ")
	buf.WriteString(varName)
	buf.WriteRune(hardcore.EOL)
	buf.WriteString("}")
	buf.WriteRune(hardcore.EOL)
	buf.WriteRune(hardcore.EOL)
}

func (e *enumBuilder) writeSliceGetter(buf *bytes.Buffer, name, returnType, varName string, isPointer bool) {
	buf.WriteString(fmt.Sprintf("func Get%s() []", name))
	if isPointer {
		buf.WriteString("*")
	}
	buf.WriteString(fmt.Sprintf("%s {", returnType))
	buf.WriteRune(hardcore.EOL)
	buf.WriteString("    return ")
	buf.WriteString(varName)
	buf.WriteRune(hardcore.EOL)
	buf.WriteString("}")
	buf.WriteRune(hardcore.EOL)
	buf.WriteRune(hardcore.EOL)
}

func (e *enumBuilder) writeByIdGetter(buf *bytes.Buffer, name, keyType, returnType string, insts []*enumInst, isPointer bool) {
	buf.WriteString(fmt.Sprintf("func Get%sById(id %s) ", name, keyType))
	if isPointer {
		buf.WriteString("*")
	}
	buf.WriteString(fmt.Sprintf("%s {", returnType))
	buf.WriteRune(hardcore.EOL)
	buf.WriteString("    switch id {")
	for _, inst := range insts {
		buf.WriteRune(hardcore.EOL)
		buf.WriteString(fmt.Sprintf("    case %s: return %s", inst.Id, inst.Var))
	}
	buf.WriteRune(hardcore.EOL)
	buf.WriteString("    default: return nil")
	buf.WriteRune(hardcore.EOL)
	buf.WriteString("    }")
	buf.WriteRune(hardcore.EOL)
	buf.WriteString("}")
	buf.WriteRune(hardcore.EOL)
	buf.WriteRune(hardcore.EOL)
}

func (e *enumBuilder) writeGetter(buf *bytes.Buffer, info *Struct, field *Field) {
	structShortName := strings.ToLower(string([]rune(info.Name)[0]))
	buf.WriteString(
		fmt.Sprintf(
			"func (%s *%s) Get%s() %s {",
			structShortName,
			info.Name,
			field.Name,
			field.Kind.String(),
		),
	)
	buf.WriteRune(hardcore.EOL)
	buf.WriteString(fmt.Sprintf("   return %s.%s", structShortName, field.Name))
	buf.WriteRune(hardcore.EOL)
	buf.WriteString("}")
	buf.WriteRune(hardcore.EOL)
	buf.WriteRune(hardcore.EOL)
}
