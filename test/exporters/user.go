package exporters

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

var (
	userProperties = []_UserPropertyImpl{
		newUserProperty("id", types.ProtoInt64Kind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeInt64(user.GetId())
		}),
		newUserProperty("email", types.ProtoStringKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeString(user.GetEmail())
		}),
		//newUserProperty("role", types.ProtoIntEnumKind, func(user *models.User, encoder types.Encoder) {
		//	encoder.EncodeInt(user.GetRoleId())
		//}),
		//newUserProperty("registerDate", types.ProtoTimeKind, func(user *models.User, encoder types.Encoder) {
		//	encoder.EncodeTime(user.GetRegisterDate())
		//}),
	}
)
