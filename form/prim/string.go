package prim

type StringImpl struct {
	Base
}

func (s *StringImpl) Import(rawValue interface{}) bool {
	value, ok := rawValue.(string)
	if ok {
		if s.required {
			if len(value) > 0 {
				s.setDest(value)
				return true
			} else {
				s.error = s.missing
				return false
			}
		} else {
			s.setDest(value)
			return true
		}
	} else {
		s.error = s.wrong
		return ok
	}
}

func (s *StringImpl) setDest(value string) {
	dest := s.dest.(*string)
	(*dest) = value
}

func String(name string) *StringImpl {
	p := new(StringImpl)
	p.name = name
	return p
}
