package exporter

import (
	"github.com/byorty/hardcore/types"
)

type ByteImpl struct {
	buf  []byte
	kind types.ProtoKind
}

func NewBytes(buf []byte) types.Exporter {
	exp := new(ByteImpl)
	exp.buf = buf
	exp.kind = types.ProtoStringKind
	return exp
}

func (b *ByteImpl) Len() int {
	return 0
}

func (b *ByteImpl) Get(i int) types.ExportableProperty {
	return nil
}

func (b *ByteImpl) GetProtoKind() types.ProtoKind {
	return b.kind
}

func (b *ByteImpl) Export(i int, encoder types.Encoder) {
	encoder.EncodeBytes(b.buf)
}
