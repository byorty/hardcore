package exporters

import (
	"github.com/byorty/hardcore/exporter"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

var (
	userProperties = []types.ExportedProperty{
		&UserProperty{exporter.NewPropertyByKind("id", types.ProtoInt64Kind), func(user *models.User) interface{} { return user.GetId()}},
		&UserProperty{exporter.NewPropertyByKind("email", types.ProtoStringKind), func(user *models.User) interface{} { return user.GetEmail()}},
		&UserProperty{exporter.NewPropertyByKind("role", types.ProtoEnumKind), func(user *models.User) interface{} { return user.GetRole()}},
		&UserProperty{exporter.NewPropertyByKind("registerDate", types.ProtoTimeKind), func(user *models.User) interface{} { return user.GetRegisterDate()}},
	}
)

type UserProperty struct {
	exporter.PropertyImpl
	closure func(*models.User) interface{}
}

func (u UserProperty) GetValue() interface{} {
	return u.closure(u.GetPrototyped().(*models.User))
}

type User struct {
	exporter.BaseImpl
}

func NewUser() *User {
	user.SetProperties(userProperties)
	return user
}

func (u *User) Set(name string, closure func(user *models.User) interface{}) *User {
	u.Add(&UserProperty{exporter.NewProperty(name),closure})
	return u
}

var(
	user = new(User)
)
