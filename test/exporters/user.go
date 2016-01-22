package exporters

import (
	"github.com/byorty/hardcore/exporter"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

var (
	userProperties = []types.ExportedProperty{
		NewUserProperty("id", func(user *models.User) interface{} { return user.GetId()}),
		NewUserProperty("email", func(user *models.User) interface{} { return user.GetEmail()}),
		NewUserProperty("role", func(user *models.User) interface{} { return user.GetRole()}),
		NewUserProperty("registerDate", func(user *models.User) interface{} { return user.GetRegisterDate()}),
	}
)

type UserPropertyImpl struct {
	exporter.PropertyImpl
	closure func(*models.User) interface{}
}

func NewUserProperty(name string, closure func(user *models.User) interface{}) types.ExportedProperty {
	return &UserPropertyImpl{
		exporter.NewProperty("id"),
		closure,
	}
}

func (u UserPropertyImpl) GetValue() interface{} {
	return u.closure(u.GetPrototyped().(*models.User))
}

func NewUser() types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties(userProperties)
	return exp
}
