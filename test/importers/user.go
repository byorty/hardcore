package importers

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

var (
	_userProperties = map[string]_UserPropertyImpl{
		"id": newUserProperty(types.ProtoInt64Kind, func(user *models.User, decoder types.Decoder, value []byte) {
			user.SetId(decoder.DecodeInt64(value))
		}),
		"email": newUserProperty(types.ProtoStringKind, func(user *models.User, decoder types.Decoder, value []byte) {
			user.SetEmail(decoder.DecodeString(value))
		}),
	}
)
