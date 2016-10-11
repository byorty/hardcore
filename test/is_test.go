package test

import (
	"testing"
	"github.com/byorty/hardcore/is"
	"errors"
)

func TestEq(t *testing.T) {
	if is.Eq("a", "b") {
		t.Fail()
	}
	if is.NotEq("a", "a") {
		t.Fail()
	}
	if is.Eq(1, 2) {
		t.Fail()
	}
	if is.NotEq(1, 1) {
		t.Fail()
	}
	if is.Eq(1.1, 2.1) {
		t.Fail()
	}
	if is.NotEq(1.1, 1.1) {
		t.Fail()
	}
}

func TestNil(t *testing.T) {
	if is.Nil(errors.New("some text")) {
		t.Fail()
	}
	if is.NotNil(nil) {
		t.Fail()
	}
}

func TestStr(t *testing.T) {
	if is.EmptyStr("qwerty") {
		t.Fail()
	}
	if is.NotEmptyStr("") {
		t.Fail()
	}
}
