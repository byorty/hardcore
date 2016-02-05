package prim

import (
"strconv"
"github.com/byorty/hardcore/types"
)

type FloatImpl struct {
	BaseImpl
	bitSize int
}

func (f *FloatImpl) Import(rawValue interface{}) bool {
	return f.do(f, rawValue)
}

func (f *FloatImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseFloat(strValue, f.bitSize)
	if err == nil {
		switch dest := f.dest.(type) {
		case *float32: (*dest) = float32(value)
		case *float64: (*dest) = value
		default: return false
		}
		return true
	} else {
		f.error = f.wrong
		return false
	}
}

func Float32(name string) types.Primitive {
	p := new(FloatImpl)
	p.bitSize = 32
	p.init(name)
	return p
}

func Float64(name string) types.Primitive {
	p := new(FloatImpl)
	p.bitSize = 64
	p.init(name)
	return p
}
