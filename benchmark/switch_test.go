package benchmark

import (
	"bytes"
	"fmt"
	"github.com/byorty/hardcore/types"
	"strconv"
	"testing"
)

type Property struct {
	value interface{}
	kind  types.ProtoKind
}

var (
	props = []Property{
		Property{123456789, types.ProtoIntKind},
		Property{"string", types.ProtoStringKind},
		Property{123.456, types.ProtoFloat32Kind},
		Property{true, types.ProtoBoolKind},
	}
	marshals = map[types.ProtoKind]func(*bytes.Buffer, interface{}){
		types.ProtoIntKind: func(buf *bytes.Buffer, value interface{}) {
			buf.WriteString(strconv.FormatInt(int64(value.(int)), 10))
		},
		types.ProtoStringKind: func(buf *bytes.Buffer, value interface{}) {
			buf.WriteString(value.(string))
		},
		types.ProtoFloat32Kind: func(buf *bytes.Buffer, value interface{}) {
			buf.WriteString(strconv.FormatFloat(value.(float64), 'E', -1, 32))
		},
		types.ProtoBoolKind: func(buf *bytes.Buffer, value interface{}) {
			buf.WriteString(strconv.FormatBool(value.(bool)))
		},
	}
)

func BenchmarkSwitchType(b *testing.B) {
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(props); j++ {
			prop := props[j]
			switch value := prop.value.(type) {
			case int:
				buf.WriteString(strconv.FormatInt(int64(value), 10))
			case string:
				buf.WriteString(value)
			case float64:
				buf.WriteString(strconv.FormatFloat(value, 'E', -1, 32))
			case bool:
				buf.WriteString(strconv.FormatBool(value))
			}
			buf.WriteString(",")
		}
	}
	//fmt.Println(buf.String())
}

func BenchmarkSwitchProto(b *testing.B) {
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(props); j++ {
			prop := props[j]
			switch prop.kind {
			case types.ProtoIntKind:
				buf.WriteString(strconv.FormatInt(int64(prop.value.(int)), 10))
			case types.ProtoStringKind:
				buf.WriteString(prop.value.(string))
			case types.ProtoFloat32Kind:
				buf.WriteString(strconv.FormatFloat(prop.value.(float64), 'E', -1, 32))
			case types.ProtoBoolKind:
				buf.WriteString(strconv.FormatBool(prop.value.(bool)))
			}
			buf.WriteString(",")
		}
	}
	//fmt.Println(buf.String())
}

func BenchmarkSwitchFrint(b *testing.B) {
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(props); j++ {
			prop := props[j]
			fmt.Fprint(buf, prop.value)
			//switch prop.kind {
			//case types.ProtoIntKind:
			//case types.ProtoStringKind:
			//	buf.WriteString(prop.value.(string))
			//case types.ProtoFloat32Kind:
			////fmt.Println(prop.value)
			//case types.ProtoBoolKind:
			////fmt.Println(prop.value)
			//}
			buf.WriteString(",")
		}
	}
}

func BenchmarkMarchal(b *testing.B) {
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(props); j++ {
			prop := props[j]
			m, _ := marshals[prop.kind]
			m(buf, prop.value)
			buf.WriteString(",")
		}
	}
	//fmt.Println(buf.String())
}
