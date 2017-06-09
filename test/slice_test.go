package test

import (
	"testing"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/slice"
	"fmt"
)

func TestIntSlice(t *testing.T) {
	s := slice.NewInt()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	fmt.Println(s)

	if is.NotEq(s.Get(0), 1) {
		t.Fail()
	}
	if is.NotEq(s.Get(2), 3) {
		t.Fail()
	}
}
