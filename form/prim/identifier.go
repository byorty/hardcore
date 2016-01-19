package prim

import (
	"strconv"
	"github.com/byorty/hardcore/types"
)

type IdentifierImpl struct {
	Base
	bitSize int
}

type IntIdentifierImpl struct {
	IdentifierImpl
}

func (i *IntIdentifierImpl) Import(rawValue interface{}) bool {
	return i.do(i, rawValue)
}

func (i *IntIdentifierImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseInt(strValue, 10, i.bitSize)
	if err == nil {
		switch i.bitSize {
		case 0:
			dest := i.dest.(types.IntModel)
			dest.KindDAO().ById(int(value)).One(dest)

		case 8:
			dest := i.dest.(types.Int8Model)
			dest.KindDAO().ById(int8(value)).One(dest)

		case 16:
			dest := i.dest.(types.Int16Model)
			dest.KindDAO().ById(int16(value)).One(dest)

		case 32:
			dest := i.dest.(types.Int32Model)
			dest.KindDAO().ById(int32(value)).One(dest)

		case 64:
			dest := i.dest.(types.Int64Model)
			dest.KindDAO().ById(value).One(dest)

		}
		return true
	} else {
		i.error = i.wrong
		return false
	}
}

func IntIdentifier(name string) *IntIdentifierImpl {
	p := new(IntIdentifierImpl)
	p.bitSize = 0
	p.name = name
	return p
}

func Int8Identifier(name string) *IntIdentifierImpl {
	p := new(IntIdentifierImpl)
	p.bitSize = 8
	p.name = name
	return p
}

func Int16Identifier(name string) *IntIdentifierImpl {
	p := new(IntIdentifierImpl)
	p.bitSize = 16
	p.name = name
	return p
}

func Int32Identifier(name string) *IntIdentifierImpl {
	p := new(IntIdentifierImpl)
	p.bitSize = 32
	p.name = name
	return p
}

func Int64Identifier(name string) *IntIdentifierImpl {
	p := new(IntIdentifierImpl)
	p.bitSize = 64
	p.name = name
	return p
}

type UintIdentifierImpl struct {
	IdentifierImpl
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
