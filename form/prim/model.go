package prim

import (
	"strconv"
	"github.com/byorty/hardcore/types"
)

type EntityImpl struct {
	BaseImpl
	bitSize int
}

func (i EntityImpl) isReceived(received bool) bool {
	if i.IsRequired() {
		return received
	} else {
		return true
	}
}

type IntModelImpl struct {
	EntityImpl
}

func (i *IntModelImpl) Import(rawValue interface{}) bool {
	return i.do(i, rawValue)
}

func (i *IntModelImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseInt(strValue, 10, i.bitSize)
	if err == nil {
		var isReceived bool
		switch dest := i.dest.(type) {
		case types.IntModel:
			dest.KindDAO().ById(int(value)).One(dest)
			isReceived = i.isReceived(dest.GetId() != 0)
		case types.Int8Model:
			dest.KindDAO().ById(int8(value)).One(dest)
			isReceived = i.isReceived(dest.GetId() != 0)
		case types.Int16Model:
			dest.KindDAO().ById(int16(value)).One(dest)
			isReceived = i.isReceived(dest.GetId() != 0)
		case types.Int32Model:
			dest.KindDAO().ById(int32(value)).One(dest)
			isReceived = i.isReceived(dest.GetId() != 0)
		case types.Int64Model:
			dest.KindDAO().ById(value).One(dest)
			isReceived = i.isReceived(dest.GetId() != 0)
		default: isReceived = false
		}
		if i.IsRequired() {
			if isReceived {
				return true
			} else {
				i.error = i.wrong
				return false
			}
		} else {
			return true
		}
	} else {
		if i.IsRequired() {
			i.error = i.wrong
			return false
		} else {
			return true
		}
	}
}

func IntModel(name string) types.Primitive {
	p := new(IntModelImpl)
	p.bitSize = 0
	p.init(name)
	return p
}

func Int8Model(name string) types.Primitive {
	p := new(IntModelImpl)
	p.bitSize = 8
	p.init(name)
	return p
}

func Int16Model(name string) types.Primitive {
	p := new(IntModelImpl)
	p.bitSize = 16
	p.init(name)
	return p
}

func Int32Model(name string) types.Primitive {
	p := new(IntModelImpl)
	p.bitSize = 32
	p.init(name)
	return p
}

func Int64Model(name string) types.Primitive {
	p := new(IntModelImpl)
	p.bitSize = 64
	p.init(name)
	return p
}

type UintModelImpl struct {
	EntityImpl
}

func (u *UintModelImpl) Import(rawValue interface{}) bool {
	return u.do(u, rawValue)
}

func (u *UintModelImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseUint(strValue, 10, u.bitSize)
	if err == nil {
		switch dest := u.dest.(type) {
		case types.UintModel:
			dest.KindDAO().ById(uint(value)).One(dest)
			return u.isReceived(dest.GetId() != 0)
		case types.Uint8Model:
			dest.KindDAO().ById(uint8(value)).One(dest)
			return u.isReceived(dest.GetId() != 0)
		case types.Uint16Model:
			dest.KindDAO().ById(uint16(value)).One(dest)
			return u.isReceived(dest.GetId() != 0)
		case types.Uint32Model:
			dest.KindDAO().ById(uint32(value)).One(dest)
			return u.isReceived(dest.GetId() != 0)
		case types.Uint64Model:
			dest.KindDAO().ById(value).One(dest)
			return u.isReceived(dest.GetId() != 0)
		default: return false
		}
	} else {
		u.error = u.wrong
		return false
	}
}

func UintModel(name string) types.Primitive {
	p := new(UintModelImpl)
	p.bitSize = 0
	p.init(name)
	return p
}

func Uint8Model(name string) types.Primitive {
	p := new(UintModelImpl)
	p.bitSize = 8
	p.init(name)
	return p
}

func Uint16Model(name string) types.Primitive {
	p := new(UintModelImpl)
	p.bitSize = 16
	p.init(name)
	return p
}

func Uint32Model(name string) types.Primitive {
	p := new(UintModelImpl)
	p.bitSize = 32
	p.init(name)
	return p
}

func Uint64Model(name string) types.Primitive {
	p := new(UintModelImpl)
	p.bitSize = 64
	p.init(name)
	return p
}
