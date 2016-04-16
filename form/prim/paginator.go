package prim

import (
	"github.com/byorty/hardcore/types"
	"strconv"
)

type PaginatorImpl struct {
	BaseImpl
}

func (p *PaginatorImpl) Import(rawValue interface{}) bool {
	return p.do(p, rawValue)
}

func (p *PaginatorImpl) ImportFromString(strValue string) bool {
	paginator := p.dest.(types.Paginator)
	value, err := strconv.Atoi(strValue)
	if err == nil {
		paginator.SetPage(value)
	}
	return true
}

func Paginator(name string) types.Primitive {
	p := new(PaginatorImpl)
	p.init(name)
	return p
}
