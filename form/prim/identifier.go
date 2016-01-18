package prim

import (
	"strconv"
	"github.com/byorty/hardcore/types"
)

type IdentifierImpl struct {
	Base
	bitSize int
}

func (i *IdentifierImpl) Import(rawValue interface{}) bool {
	return i.do(i, rawValue)
}

func (i *IdentifierImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseInt(strValue, 10, i.bitSize)
	if err == nil {
		dest := i.dest.(types.Model)
		switch i.bitSize {
		case 0:
			dest.CommonDAO().(types.IntModelDAO).ById(int(value)).One(dest)

		case 8:
			dest.CommonDAO().(types.Int8ModelDAO).ById(int8(value)).One(dest)

		case 16:
			dest.CommonDAO().(types.Int16ModelDAO).ById(int16(value)).One(dest)

		case 32:
			dest.CommonDAO().(types.Int32ModelDAO).ById(int32(value)).One(dest)

		case 64:
			dest.CommonDAO().(types.Int64ModelDAO).ById(value).One(dest)

		}
		return true
	} else {
		i.error = i.wrong
		return false
	}
}

func IntIdentifier(name string) *IdentifierImpl {
	p := new(IdentifierImpl)
	p.bitSize = 0
	p.name = name
	return p
}

func Int8Identifier(name string) *IdentifierImpl {
	p := new(IdentifierImpl)
	p.bitSize = 8
	p.name = name
	return p
}

func Int16Identifier(name string) *IdentifierImpl {
	p := new(IdentifierImpl)
	p.bitSize = 16
	p.name = name
	return p
}

func Int32Identifier(name string) *IdentifierImpl {
	p := new(IdentifierImpl)
	p.bitSize = 32
	p.name = name
	return p
}

func Int64Identifier(name string) *IdentifierImpl {
	p := new(IdentifierImpl)
	p.bitSize = 64
	p.name = name
	return p
}
