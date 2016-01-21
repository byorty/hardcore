package prim

import (
"strconv"
"github.com/byorty/hardcore/types"
)

type UintImpl struct {
	Base
	bitSize int
}

func (u *UintImpl) Import(rawValue interface{}) bool {
	return u.do(u, rawValue)
}

func (u *UintImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseUint(strValue, 10, u.bitSize)
	if err == nil {
		switch dest := u.dest.(type) {
		case *uint: (*dest) = uint(value)
		case *uint8: (*dest) = uint8(value)
		case *uint16: (*dest) = uint16(value)
		case *uint32: (*dest) = uint32(value)
		case *uint64: (*dest) = uint64(value)
		default: return false
		}
		return true
	} else {
		u.error = u.wrong
		return false
	}
}

func Uint(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 0
	p.name = name
	return p
}

func Uint8(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 8
	p.name = name
	return p
}

func Uint16(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 16
	p.name = name
	return p
}

func Uint32(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 32
	p.name = name
	return p
}

func Uint64(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 64
	p.name = name
	return p
}
