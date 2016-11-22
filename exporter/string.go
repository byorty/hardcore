package exporter

import (
	"github.com/byorty/hardcore/types"
)

type StringImpl struct {
	str  string
	kind types.ProtoKind
}

func NewString(str string) types.Exporter {
	exp := new(StringImpl)
	exp.str = str
	exp.kind = types.ProtoStringKind
	return exp
}

func (s *StringImpl) Len() int {
	return 0
}

func (s *StringImpl) Get(i int) types.ExportableProperty {
	return nil
}

func (s *StringImpl) GetProtoKind() types.ProtoKind {
	return s.kind
}

func (s *StringImpl) Export(i int, encoder types.Encoder) {
	encoder.EncodeString(s.str)
}
