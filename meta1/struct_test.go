package meta

import (
	"testing"
	"reflect"
)

type S struct {
	A int          `sql:"pk;auto"`
	B string       `sql:"name:bbb;uniq"`
	C []*S
	D map[string]S
}

func TestStruct(t *testing.T) {
	info := AnalyzeStruct(S{})
	if info.Name != "S" {
		t.Fail()
	}
	if info.fullname != "github.com/byorty/hardcore/meta/S" {
		t.Fail()
	}
}

func TestStructPtr(t *testing.T) {
	info := AnalyzeStruct(&S{})
	if info.Name != "S" {
		t.Fail()
	}
	if info.fullname != "github.com/byorty/hardcore/meta/S" {
		t.Fail()
	}
}

func TestStructFields(t *testing.T) {
	info := AnalyzeStruct(&S{})
	if info.FieldMap["A"].Name != "A" {
		t.Fail()
	}
	if info.FieldMap["A"].Kind != reflect.Int {
		t.Fail()
	}
	if info.FieldMap["A"].Tag.Get("sql") != "pk;auto" {
		t.Fail()
	}

	if info.Fields[1].Name != "B" {
		t.Fail()
	}
	if info.Fields[1].Kind != reflect.String {
		t.Fail()
	}
	if info.Fields[1].Tag.Get("sql") != "name:bbb;uniq" {
		t.Fail()
	}
}
