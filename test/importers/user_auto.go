package importers

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

type _UserImpl struct {
	model *models.User
	props map[string]_UserPropertyImpl
}

func (u _UserImpl) Get(key string) (types.ImportableProperty, bool) {
	prop, ok := u.props[key]
	return prop, ok
}

func (u _UserImpl) Decode(key string, decoder types.Decoder, value []byte) {
	u.props[key].closure(u.model, decoder, value)
}

type _UserPropertyImpl struct {
	kind    types.ProtoKind
	closure func(*models.User, types.Decoder, []byte)
}

func (u _UserPropertyImpl) GetProtoKind() types.ProtoKind {
	return u.kind
}

func newUserProperty(kind types.ProtoKind, closure func(*models.User, types.Decoder, []byte)) _UserPropertyImpl {
	return _UserPropertyImpl{
		kind,
		closure,
	}
}

func NewUser(user *models.User) types.Importer {
	imp := new(_UserImpl)
	imp.model = user
	imp.props = _userProperties
	return imp
}
