package prim

import "strconv"

type FloatImpl struct {
	Base
	bitSize int
}

func (f *FloatImpl) Import(rawValue interface{}) bool {
	return f.do(f, rawValue)
}

func (f *FloatImpl) ImportFromString(strValue string) bool {
	value, err := strconv.ParseFloat(strValue, f.bitSize)
	if err == nil {
		switch f.bitSize {
		case 32:
			dest := f.dest.(*float32)
			(*dest) = float32(value)

		case 64:
			dest := f.dest.(*float64)
			(*dest) = value

		}
		return true
	} else {
		f.error = f.wrong
		return false
	}
}

func Float32(name string) *FloatImpl {
	p := new(FloatImpl)
	p.bitSize = 32
	p.name = name
	return p
}

func Float64(name string) *FloatImpl {
	p := new(FloatImpl)
	p.bitSize = 64
	p.name = name
	return p
}
