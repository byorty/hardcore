package exporters

import (
	"github.com/byorty/hardcore/exporter"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

type _MsgpackImpl struct {
	item  *models.User
	items models.Users
	props []_MsgpackPropertyImpl
	kind  types.ProtoKind
}

func NewMsgpack(model *models.User) types.Exporter {
	return _newMsgpack(model, msgpackProperties)
}

func _newMsgpack(model *models.User, props []_MsgpackPropertyImpl) types.Exporter {
	exp := new(_MsgpackImpl)
	exp.props = props
	exp.item = model
	exp.kind = types.ProtoModelKind
	return exp
}

func (m _MsgpackImpl) Len() int {
	if is.Eq(m.kind, types.ProtoModelKind) {
		return len(m.props)
	} else {
		return m.items.Len()
	}
}

func (m _MsgpackImpl) Get(i int) types.ExportableProperty {
	return m.props[i]
}

func (m _MsgpackImpl) GetProtoKind() types.ProtoKind {
	return m.kind
}

func (m _MsgpackImpl) Export(i int, encoder types.Encoder) {
	if is.Eq(m.kind, types.ProtoModelKind) {
		m.props[i].closure(m.item, encoder)
	} else {
		encoder.Encode(_newMsgpack(m.items.Get(i), msgpackProperties))
	}
}

func NewMsgpacks(users models.Users) types.Exporter {
	exp := new(_MsgpackImpl)
	exp.props = msgpackProperties
	exp.items = users
	exp.kind = types.ProtoModelSliceKind
	return exp
}

type _MsgpackPropertyImpl struct {
	exporter.PropertyImpl
	closure func(*models.User, types.Encoder)
}

func newMsgpackProperty(name string, kind types.ProtoKind, closure func(*models.User, types.Encoder)) _MsgpackPropertyImpl {
	return _MsgpackPropertyImpl{
		exporter.NewPropertyWithKind(name, kind),
		closure,
	}
}
