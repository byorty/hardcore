package prim

import (
	"github.com/byorty/hardcore/types"
	"strconv"
)

type UintImpl struct {
	BaseImpl
	bitSize int
}

func (u *UintImpl) Import(rawValue interface{}) bool {
	return u.do(u, rawValue)
}

func (u *UintImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseUint(strValue, 10, u.bitSize)
	if err == nil {
		switch dest := u.dest.(type) {
		case *uint:
			(*dest) = uint(value)
		case *uint8:
			(*dest) = uint8(value)
		case *uint16:
			(*dest) = uint16(value)
		case *uint32:
			(*dest) = uint32(value)
		case *uint64:
			(*dest) = value
		default:
			return false
		}
		return true
	} else {
		if u.IsRequired() {
			u.error = u.wrong
			return false
		} else {
			return true
		}
	}
}

func Uint(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 0
	p.init(name)
	return p
}

func Uint8(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 8
	p.init(name)
	return p
}

func Uint16(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 16
	p.init(name)
	return p
}

func Uint32(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 32
	p.init(name)
	return p
}

func Uint64(name string) types.Primitive {
	p := new(UintImpl)
	p.bitSize = 64
	p.init(name)
	return p
}
