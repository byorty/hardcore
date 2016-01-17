package prim

import "strconv"

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
		switch u.bitSize {
		case 0:
			dest := u.dest.(*uint)
			(*dest) = uint(value)

		case 8:
			dest := u.dest.(*uint8)
			(*dest) = uint8(value)

		case 16:
			dest := u.dest.(*uint16)
			(*dest) = uint16(value)

		case 32:
			dest := u.dest.(*uint32)
			(*dest) = uint32(value)

		case 64:
			dest := u.dest.(*uint64)
			(*dest) = uint64(value)

		}
		return true
	} else {
		u.error = u.wrong
		return false
	}
}

func Uint(name string) *UintImpl {
	p := new(UintImpl)
	p.bitSize = 0
	p.name = name
	return p
}

func Uint8(name string) *UintImpl {
	p := new(UintImpl)
	p.bitSize = 8
	p.name = name
	return p
}

func Uint16(name string) *UintImpl {
	p := new(UintImpl)
	p.bitSize = 16
	p.name = name
	return p
}

func Uint32(name string) *UintImpl {
	p := new(UintImpl)
	p.bitSize = 32
	p.name = name
	return p
}

func Uint64(name string) *UintImpl {
	p := new(UintImpl)
	p.bitSize = 64
	p.name = name
	return p
}
