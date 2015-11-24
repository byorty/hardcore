package prim

type StringPrimitiveImpl struct {
	BasePrimitive
}

func (s *StringPrimitiveImpl) Import(rawValue interface{}) bool {
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

func (s *StringPrimitiveImpl) setDest(value string) {
	dest := s.dest.(*string)
	(*dest) = value
}

func String(name string) *StringPrimitiveImpl {
	p := new(StringPrimitiveImpl)
	p.name = name
	return p
}
