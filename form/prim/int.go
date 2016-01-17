package prim

import "strconv"

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
		switch i.bitSize {
		case 0:
			dest := i.dest.(*int)
			(*dest) = int(value)

		case 8:
			dest := i.dest.(*int8)
			(*dest) = int8(value)

		case 16:
			dest := i.dest.(*int16)
			(*dest) = int16(value)

		case 32:
			dest := i.dest.(*int32)
			(*dest) = int32(value)

		case 64:
			dest := i.dest.(*int64)
			(*dest) = int64(value)

		}
		return true
	} else {
		i.error = i.wrong
		return false
	}
}

func Int(name string) *IntImpl {
	p := new(IntImpl)
	p.bitSize = 0
	p.name = name
	return p
}

func Int8(name string) *IntImpl {
	p := new(IntImpl)
	p.bitSize = 8
	p.name = name
	return p
}

func Int16(name string) *IntImpl {
	p := new(IntImpl)
	p.bitSize = 16
	p.name = name
	return p
}

func Int32(name string) *IntImpl {
	p := new(IntImpl)
	p.bitSize = 32
	p.name = name
	return p
}

func Int64(name string) *IntImpl {
	p := new(IntImpl)
	p.bitSize = 64
	p.name = name
	return p
}

