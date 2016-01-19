package prim

import (
	"strconv"
	"github.com/byorty/hardcore/types"
)

type UintIdentifierImpl struct {
	Base
	bitSize int
}

func (u *UintIdentifierImpl) Import(rawValue interface{}) bool {
	return u.do(u, rawValue)
}

func (u *UintIdentifierImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseUint(strValue, 10, u.bitSize)
	if err == nil {
		switch u.bitSize {
		case 0:
			dest := u.dest.(types.UintModel)
			dest.KindDAO().ById(uint(value)).One(dest)

		case 8:
			dest := u.dest.(types.Uint8Model)
			dest.KindDAO().ById(uint8(value)).One(dest)

		case 16:
			dest := u.dest.(types.Uint16Model)
			dest.KindDAO().ById(uint16(value)).One(dest)

		case 32:
			dest := u.dest.(types.Uint32Model)
			dest.KindDAO().ById(uint32(value)).One(dest)

		case 64:
			dest := u.dest.(types.Uint64Model)
			dest.KindDAO().ById(value).One(dest)

		}
		return true
	} else {
		u.error = u.wrong
		return false
	}
}

func UintIdentifier(name string) *UintIdentifierImpl {
	p := new(UintIdentifierImpl)
	p.bitSize = 0
	p.name = name
	return p
}

func Uint8Identifier(name string) *UintIdentifierImpl {
	p := new(UintIdentifierImpl)
	p.bitSize = 8
	p.name = name
	return p
}

func Uint16Identifier(name string) *UintIdentifierImpl {
	p := new(UintIdentifierImpl)
	p.bitSize = 16
	p.name = name
	return p
}

func Uint32Identifier(name string) *UintIdentifierImpl {
	p := new(UintIdentifierImpl)
	p.bitSize = 32
	p.name = name
	return p
}

func Uint64Identifier(name string) *UintIdentifierImpl {
	p := new(UintIdentifierImpl)
	p.bitSize = 64
	p.name = name
	return p
}
