package view

import "github.com/byorty/hardcore/types"

type StringImpl struct {
	BaseImpl
	str string
}

func String(str string) types.StringView {
	return &StringImpl{
		str: str,
	}
}

func (s *StringImpl) SetStatus(status int) types.StringView {
	s.status = status
	return s
}

func (s *StringImpl) Render() {
	s.BaseImpl.Render()
	rw := s.scope.GetWriter()
	rw.Write([]byte(s.str))
}
