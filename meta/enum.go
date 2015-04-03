package meta

import (
	"fmt"
	"strings"
	"reflect"
	"bytes"
	"text/template"
)

var EnumTypes = map[string]string{
	reflect.Int.String()       : "hardcore.IntEnum",
	reflect.Int8.String()      : "hardcore.Int8Enum",
	reflect.Int16.String()     : "hardcore.Int16Enum",
	reflect.Int32.String()     : "hardcore.Int32Enum",
	reflect.Int64.String()     : "hardcore.Int64Enum",
	reflect.Uint.String()      : "hardcore.UintEnum",
	reflect.Uint8.String()     : "hardcore.Uint8Enum",
	reflect.Uint16.String()    : "hardcore.Uint16Enum",
	reflect.Uint32.String()    : "hardcore.Uint32Enum",
	reflect.Uint64.String()    : "hardcore.Uint63Enum",
	reflect.Float32.String()   : "hardcore.Float32Enum",
	reflect.Float64.String()   : "hardcore.Float64Enum",
	reflect.Complex64.String() : "hardcore.Complex64Enum",
	reflect.Complex128.String(): "hardcore.Complex128Enum",
	reflect.String.String()    : "hardcore.StringEnum",
}

var EnumTmpl = `{{$s := .ShortStructName}}` +
`{{$type := .StructType}}` +
`{{$types := .StructsType}}` +
`func New{{$type}}ById(id {{.IdType}}) *{{$type}} {
    {{$s}} := new({{$type}}).ById(id)
    if {{$s}} == nil {
        return nil
    } else {
        return {{$s}}.(*{{$type}})
    }
}

func ({{$s}} *{{$type}}) GetId() {{.IdType}} {
    return {{$s}}.Id
}

func ({{$s}} *{{$type}}) GetName() string {
    return {{$s}}.Name
}

func ({{$s}} *{{$type}}) ById(id {{.IdType}}) {{.Identifiable}} {
    switch id {
    {{range .InitParams}}case {{.IdValue}}:
        {{range $name, $value := .Params}}{{$s}}.{{$name}} = {{$value}}
        {{end}}break
    {{end}}default: return nil
    }
    return {{$s}}
}

type {{$types}} []*{{$type}}

func ({{$s}} *{{$types}}) All() {
{{range .InitParams}}    (*{{$s}}) = append((*{{$s}}), New{{$type}}ById({{.IdValue}}))
{{end}}}

func ({{$s}} *{{$types}}) ByIds(ids []{{.IdType}}) {
    enumMap := New{{$type}}Map()
    enumMap.All()
    for _, id := range ids {
        if enum, ok := enumMap[id]; ok {
            (*{{$s}}) = append((*{{$s}}), enum)
        }
    }
}

type {{$type}}Map map[{{.IdType}}]*{{$type}}

func New{{$type}}Map() {{$type}}Map {
    return make({{$type}}Map)
}

func ({{$s}} {{$type}}Map) All() {
{{range .InitParams}}    {{$s}}[{{.IdValue}}] = New{{$type}}ById({{.IdValue}})
{{end}}}

func ({{$s}} {{$type}}Map) ByIds(ids []{{.IdType}}) {
    enumMap := New{{$type}}Map()
    enumMap.All()
    for _, id := range ids {
        if enum, ok := enumMap[id]; ok {
            {{$s}}[id] = enum
        }
    }
}

`

type enumParams struct {
	IdValue string
	Params  map[string]string
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
			initParams := make([]*enumParams, 0)
			for _, field := range info.Fields {
				if len(field.Tag.Get("enum")) > 0 {
					values := strings.Split(field.Tag.Get("enum"), ",")
					for j, value := range values {
						if field.IsId() {
							initParam := &enumParams{
								IdValue: value,
								Params: make(map[string]string),
							}
							initParams = append(initParams, initParam)
						}
						if field.Kind == reflect.String {
							value = fmt.Sprintf("\"%s\"", value)
						}
						initParams[j].Params[field.Name] = value
					}
				}
			}

			if len(initParams) > 0 {
				tmplParams := map[string]interface{}{
					"ShortStructName": strings.ToLower(info.Name[0:1]),
					"StructType"     : info.Name,
					"StructsType"    : info.PluralName,
					"Identifiable"   : IdentifiableByType[info.IdField.Kind.String()],
					"IdType"         : info.IdField.Kind.String(),
					"InitParams"     : initParams,
				}
				tmpl := template.New(info.Name + "Template")
				tmpl.Parse(EnumTmpl)
				tmpl.Execute(buf, tmplParams)

				results = append(results, &BuildResult{Pkg: info.Pkg, PkgPath: info.PkgPath, Bytes: buf.Bytes()})
			}
		}
	}
	return results
}
