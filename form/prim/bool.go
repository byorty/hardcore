package prim

import (
"strconv"
"github.com/byorty/hardcore/types"
)

type BoolImpl struct {
	BaseImpl
}

func (b *BoolImpl) Import(rawValue interface{}) bool {
	return b.do(b, rawValue)
}

func (b *BoolImpl) ImportFromString(strValue string) bool {
	dest := b.dest.(*bool)
	boolValue, err := strconv.ParseBool(strValue)
	if err == nil {
		(*dest) = boolValue
		return true
	} else {
		b.error = b.wrong
		return false
	}
}

func Bool(name string) types.Primitive {
	p := new(BoolImpl)
	p.init(name)
	return p
}
