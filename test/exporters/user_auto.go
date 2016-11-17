package exporters

import (
	"github.com/byorty/hardcore/exporter"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

type _UserImpl struct {
	item  *models.User
	items models.Users
	props []_UserPropertyImpl
	kind  types.ProtoKind
}

func NewUser(model *models.User) types.Exporter {
	return _newUser(model, userProperties)
}

func _newUser(model *models.User, props []_UserPropertyImpl) types.Exporter {
	exp := new(_UserImpl)
	exp.props = props
	exp.item = model
	exp.kind = types.ProtoModelKind
	return exp
}

func (u _UserImpl) Len() int {
	if is.Eq(u.kind, types.ProtoModelKind) {
		return len(u.props)
	} else {
		return len(u.items)
	}
}

func (u _UserImpl) Get(i int) types.ExportableProperty {
	return u.props[i]
}

func (u _UserImpl) GetProtoKind() types.ProtoKind {
	return u.kind
}

func (u _UserImpl) Export(i int, encoder types.Encoder) {
	if is.Eq(u.kind, types.ProtoModelKind) {
		u.props[i].closure(u.item, encoder)
	} else {
		encoder.Encode(_newUser(u.items[i], userProperties))
	}
}

func NewUsers(users models.Users) types.Exporter {
	exp := new(_UserImpl)
	exp.props = userProperties
	exp.items = users
	exp.kind = types.ProtoModelSliceKind
	return exp
}

type _UserPropertyImpl struct {
	exporter.PropertyImpl
	closure func(*models.User, types.Encoder)
}

func newUserProperty(name string, kind types.ProtoKind, closure func(*models.User, types.Encoder)) _UserPropertyImpl {
	return _UserPropertyImpl{
		exporter.NewPropertyWithKind(name, kind),
		closure,
	}
}
