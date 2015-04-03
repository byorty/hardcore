package meta

import (
	"testing"
//	"fmt"
//	"github.com/byorty/hardcore"
	"fmt"
	"time"
)

var (
	context IContext
	handler interface{}
)

type ISome interface {
	Actions()
}

type ISomeFunc func() ISome

type A struct {

}

func NewA() *A {
	return new(A)
}

func (a *A) Actions() {

}

type B struct {

}

func NewB() *B {
	return new(B)
}

func (b *B) Actions() {

}

func NewSome() {

}



type IContext interface {
	Init()
}

//type IContextFunc func(IContext)()
type IContextFunc func()

func (i *IContextFunc) Boo() {
	fmt.Println("boo")
	fmt.Println(*i)
}

type Context struct {
	Name string
	I    int
}

func (c *Context) Init() {
	c.Name = time.Now().String()
	c.I++
}

func (c *Context) F() {
	fmt.Println("Context I:", c.I)
}

func TestBuilder1(t *testing.T) {
//	var o SomeStatuses
//	oPtr := &o
//	oPtr.All()
//	fmt.Println(oPtr)
//	fmt.Println(o)
//
//	oMap := NewSomeStatusMap()
//	oMap.All()
//	fmt.Println(oMap)
//	a(&o)
//	b := hardcore.Enum(&o).GetId()
//	fmt.Println("b :=", b)
//	var s SomeStatuses
//	fmt.Println(s.All())
	RegisterEnums(
		&SomeStatus{},
//		&SomeType{},
	)
	Build()

//	c(&Context{})
//	h((*Context).F)
//	h(*((*Context).F))
//	for i := 0;i < 10;i++ {
//		go call()
//	}

//	f := (*Context).F
//	f()

//	ff := (*Context).F
//	ff := (&Context{}).F
//	ff(&Context{})
//	fff()
//	fmt.Println("-->", ff)


//	time.Sleep(time.Second)
}

//func a(e hardcore.Enum) {}

func c(c IContext) {
	context = c
}

func h(h interface {}) {
	handler = h
}

func call() {
	context.Init()
	handler.(func(*Context)())(context.(*Context))
}
