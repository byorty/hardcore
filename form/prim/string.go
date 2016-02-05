package prim

import "github.com/byorty/hardcore/types"

type StringImpl struct {
	BaseImpl
}

func (s *StringImpl) Import(rawValue interface{}) bool {
	return s.do(s, rawValue)
}

func (s *StringImpl) ImportFromString(strValue string) bool {
	dest := s.dest.(*string)
	(*dest) = strValue
	return true
}

func String(name string) types.Primitive {
	p := new(StringImpl)
	p.init(name)
	return p
}
