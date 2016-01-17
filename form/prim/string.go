package prim

type StringImpl struct {
	Base
}

func (s *StringImpl) Import(rawValue interface{}) bool {
	return s.do(s, rawValue)
}

func (s *StringImpl) ImportFromString(strValue string) bool {
	dest := s.dest.(*string)
	(*dest) = strValue
	return true
}

func String(name string) *StringImpl {
	p := new(StringImpl)
	p.name = name
	return p
}
