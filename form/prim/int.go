package prim

import "strconv"

type Int struct {
	Base
}

func (i *Int) Import(rawValue interface{}) bool {
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

func (i *Int) setDest(value int) {
	dest := i.dest.(*int)
	(*dest) = value
}

func Int(name string) *Int {
	p := new(Int)
	p.name = name
	return p
}

