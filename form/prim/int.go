package prim

import (
"strconv"
"github.com/byorty/hardcore/types"
)

type IntImpl struct {
	Base
	bitSize int
}

func (i *IntImpl) Import(rawValue interface{}) bool {
	return i.do(i, rawValue)
}

func (i *IntImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseInt(strValue, 10, i.bitSize)
	if err == nil {
		switch dest := i.dest.(type) {
		case *int: (*dest) = int(value)
		case *int8: (*dest) = int8(value)
		case *int16: (*dest) = int16(value)
		case *int32: (*dest) = int32(value)
		case *int64: (*dest) = int64(value)
		default: return false
		}
		return true
	} else {
		i.error = i.wrong
		return false
	}
}

func Int(name string) types.Primitive {
	p := new(IntImpl)
	p.bitSize = 0
	p.name = name
	return p
}

func Int8(name string) types.Primitive {
	p := new(IntImpl)
	p.bitSize = 8
	p.name = name
	return p
}

func Int16(name string) types.Primitive {
	p := new(IntImpl)
	p.bitSize = 16
	p.name = name
	return p
}

func Int32(name string) types.Primitive {
	p := new(IntImpl)
	p.bitSize = 32
	p.name = name
	return p
}

func Int64(name string) types.Primitive {
	p := new(IntImpl)
	p.bitSize = 64
	p.name = name
	return p
}

