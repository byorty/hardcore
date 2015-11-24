package prim

import "strconv"

type IntPrimitiveImpl struct {
	BasePrimitive
}

func (i *IntPrimitiveImpl) Import(rawValue interface{}) bool {
	strValue, ok := rawValue.(string)
	if ok {
		value, err := strconv.Atoi(strValue)
		if err == nil {
			i.setDest(value)
			return true
		} else {
			i.error = i.wrong
			return false
		}
	} else {
		i.error = i.wrong
		return ok
	}
}

func (i *IntPrimitiveImpl) setDest(value int) {
	dest := i.dest.(*int)
	(*dest) = value
}

func Int(name string) *IntPrimitiveImpl {
	p := new(IntPrimitiveImpl)
	p.name = name
	return p
}

