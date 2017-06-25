package prim

import (
	"github.com/byorty/hardcore/types"
	"strconv"
)

type EnumImpl struct {
	EntityImpl
}

func (e *EnumImpl) Import(rawValue interface{}) bool {
	return e.do(e, rawValue)
}

func (e *EnumImpl) ImportFromString(strValue string) bool {
	switch e.dest.(type) {
	case types.IntEnum, types.Int8Enum, types.Int16Enum, types.Int32Enum, types.Int64Enum:
		value, err := strconv.ParseInt(strValue, 10, e.bitSize)
		if err == nil {
			return e.importInt(value)
		} else {
			e.error = e.wrong
			return false
		}
	case types.UintEnum, types.Uint8Enum, types.Uint16Enum, types.Uint32Enum, types.Uint64Enum:
		value, err := strconv.ParseUint(strValue, 10, e.bitSize)
		if err == nil {
			return e.importUint(value)
		} else {
			e.error = e.wrong
			return false
		}
	case types.Float32Enum, types.Float64Enum:
		value, err := strconv.ParseFloat(strValue, e.bitSize)
		if err == nil {
			return e.importFloat(value)
		} else {
			e.error = e.wrong
			return false
		}
	case types.StringEnum:
		dest := e.dest.(types.StringEnum)
		dest.DAO().ById(strValue).One(dest)
		return e.isReceived(dest.GetId() != "")
	default:
		return false
	}
}

func (e EnumImpl) importInt(value int64) bool {
	switch dest := e.dest.(type) {
	case types.IntEnum:
		dest.DAO().ById(int(value)).One(dest)
		return e.isReceived(dest.GetId() != 0)
	case types.Int8Enum:
		dest.DAO().ById(int8(value)).One(dest)
		return e.isReceived(dest.GetId() != 0)
	case types.Int16Enum:
		dest.DAO().ById(int16(value)).One(dest)
		return e.isReceived(dest.GetId() != 0)
	case types.Int32Enum:
		dest.DAO().ById(int32(value)).One(dest)
		return e.isReceived(dest.GetId() != 0)
	case types.Int64Enum:
		dest.DAO().ById(value).One(dest)
		return e.isReceived(dest.GetId() != 0)
	default:
		return false
	}
}

func (e EnumImpl) importUint(value uint64) bool {
	switch dest := e.dest.(type) {
	case types.UintEnum:
		dest.DAO().ById(uint(value)).One(dest)
		return e.isReceived(dest.GetId() != 0)
	case types.Uint8Enum:
		dest.DAO().ById(uint8(value)).One(dest)
		return e.isReceived(dest.GetId() != 0)
	case types.Uint16Enum:
		dest.DAO().ById(uint16(value)).One(dest)
		return e.isReceived(dest.GetId() != 0)
	case types.Uint32Enum:
		dest.DAO().ById(uint32(value)).One(dest)
		return e.isReceived(dest.GetId() != 0)
	case types.Uint64Enum:
		dest.DAO().ById(value).One(dest)
		return e.isReceived(dest.GetId() != 0)
	default:
		return false
	}
}

func (e EnumImpl) importFloat(value float64) bool {
	switch dest := e.dest.(type) {
	case types.Float32Enum:
		dest.DAO().ById(float32(value)).One(dest)
		return e.isReceived(dest.GetId() != 0)
	case types.Float64Enum:
		dest.DAO().ById(value).One(dest)
		return e.isReceived(dest.GetId() != 0)
	default:
		return false
	}
}

func IntEnum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 0
	p.init(name)
	return p
}

func Int8Enum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 8
	p.init(name)
	return p
}

func Int16Enum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 16
	p.init(name)
	return p
}

func Int32Enum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 32
	p.init(name)
	return p
}

func Int64Enum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 64
	p.init(name)
	return p
}

func UintEnum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 0
	p.init(name)
	return p
}

func Uint8Enum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 8
	p.init(name)
	return p
}

func Uint16Enum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 16
	p.init(name)
	return p
}

func Uint32Enum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 32
	p.init(name)
	return p
}

func Uint64Enum(name string) types.Primitive {
	p := new(EnumImpl)
	p.bitSize = 64
	p.init(name)
	return p
}

func StringEnum(name string) types.Primitive {
	p := new(EnumImpl)
	p.init(name)
	return p
}