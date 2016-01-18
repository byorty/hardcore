package prim
import "strconv"

type BoolImpl struct {
	Base
}

func (b *BoolImpl) Import(rawValue interface{}) bool {
	return b.do(b, rawValue)
}

func (b *BoolImpl) ImportFromString(strValue string) bool {
	dest := b.dest.(*bool)
	boolValue, err := strconv.ParseBool(strValue)
	if err == nil {
		(*dest) = boolValue
		return true
	} else {
		b.error = b.wrong
		return false
	}
}

func Bool(name string) *BoolImpl {
	p := new(BoolImpl)
	p.name = name
	return p
}
