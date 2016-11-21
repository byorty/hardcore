package importers

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

type _MsgpackImpl struct {
	model *models.User
	props map[string]_MsgpackPropertyImpl
}

func (m _MsgpackImpl) Get(key string) (types.ImportableProperty, bool) {
	prop, ok := m.props[key]
	return prop, ok
}

func (m _MsgpackImpl) Decode(key string, decoder types.Decoder, value []byte) {
	m.props[key].closure(m.model, decoder, value)
}

type _MsgpackPropertyImpl struct {
	kind    types.ProtoKind
	closure func(*models.User, types.Decoder, []byte)
}

func (m _MsgpackPropertyImpl) GetProtoKind() types.ProtoKind {
	return m.kind
}

func newMsgpackProperty(kind types.ProtoKind, closure func(*models.User, types.Decoder, []byte)) _MsgpackPropertyImpl {
	return _MsgpackPropertyImpl{
		kind,
		closure,
	}
}

func NewMsgpack(user *models.User) types.Importer {
	imp := new(_MsgpackImpl)
	imp.model = user
	imp.props = _msgpackProperties
	return imp
}
