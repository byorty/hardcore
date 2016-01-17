package prim

import "strconv"

type IntImpl struct {
	Base
}

func (i *IntImpl) Import(rawValue interface{}) bool {
	strValue, ok := rawValue.(string)
	if ok {
		if len(strValue) == 0 {
			i.error = i.missing
			return false
		} else {
			value, err := strconv.Atoi(strValue)
			if err == nil {
				i.setDest(value)
				return true
			} else {
				i.error = i.wrong
				return false
			}
		}
	} else {
		i.error = i.wrong
		return ok
	}
}

func (i *IntImpl) setDest(value int) {
	dest := i.dest.(*int)
	(*dest) = value
}

func Int(name string) *IntImpl {
	p := new(IntImpl)
	p.name = name
	return p
}

