package view
import "github.com/byorty/hardcore/types"

type StringImpl struct {
	BaseImpl
	str string
}

func String(str string) types.View {
	return &StringImpl{
		str: str,
	}
}

func (s *StringImpl) Render() {
	rw := s.scope.GetWriter()
	rw.Write([]byte(s.str))
}




