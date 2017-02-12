package importer

import "github.com/byorty/hardcore/types"

type ByteImpl struct {
	buf *[]byte
}

func NewBytes(buf *[]byte) types.Importer {
	imp := new(ByteImpl)
	imp.buf = buf
	return imp
}

func (b *ByteImpl) GetProtoKind() types.ProtoKind {
	return types.ProtoByteKind
}

func (b *ByteImpl) Get(key string) (types.ImportableProperty, bool) {
	return nil, true
}

func (b *ByteImpl) Decode(key string, decoder types.Decoder, value []byte) {
	(*b.buf) = []byte(decoder.DecodeString(value))
}
