package exporter

import (
	"github.com/byorty/hardcore/types"
)

type ByteImpl struct {
	buf []byte
}

func NewBytes(buf []byte) types.Exporter {
	exp := new(ByteImpl)
	exp.buf = buf
	return exp
}

func (b *ByteImpl) Len() int {
	return 0
}

func (b *ByteImpl) Get(i int) types.ExportableProperty {
	return nil
}

func (b *ByteImpl) GetProtoKind() types.ProtoKind {
	return types.ProtoByteKind
}

func (b *ByteImpl) Export(i int, encoder types.Encoder) {
	encoder.EncodeBytes(b.buf)
}
