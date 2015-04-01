package meta

import (
	"fmt"
	"strings"
	"reflect"
	"bytes"
	"github.com/byorty/hardcore"
)

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
				e.writeMap(buf, mapName, info.IdField.Kind.String(), info.Name, true, insts)
				e.writeMap(buf, mapNameAsEnum, info.IdField.Kind.String(), "hardcore.Enum", false, insts)
				e.writeSlice(buf, sliceName, info.Name, true, insts)
				e.writeSlice(buf, sliceNameAsEnum, "hardcore.Enum", false, insts)

				buf.WriteString(")")
				buf.WriteRune(hardcore.EOL)
				buf.WriteRune(hardcore.EOL)

				e.writeMapGetter(buf, info.Name + "Map", info.IdField.Kind.String(), info.Name, mapName, true)
				e.writeMapGetter(buf, info.Name + "AsEnumMap", info.IdField.Kind.String(), "hardcore.Enum", mapNameAsEnum, false)
				e.writeSliceGetter(buf, info.Name + "Slice", info.Name, sliceName, true)
				e.writeSliceGetter(buf, info.Name + "AsEnumSlice", "hardcore.Enum", sliceNameAsEnum, false)

				e.writeByIdGetter(buf, info.Name, info.IdField.Kind.String(), info.Name, insts, true)
				e.writeByIdGetter(buf, info.Name + "AsEnum", info.IdField.Kind.String(), "hardcore.Enum", insts, false)

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
