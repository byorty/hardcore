package exporters

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/exporter"
	"github.com/byorty/hardcore/test/models"
)

type UserPropertyImpl struct {
	exporter.PropertyImpl
	closure func(*models.User) interface{}
}

func NewUserProperty(name string, closure func(user *models.User) interface{}) types.ExportableProperty {
	return &UserPropertyImpl{
		exporter.NewProperty(name),
		closure,
	}
}

func (u UserPropertyImpl) GetValue(model types.Model) interface{} {
	return u.closure(model.(*models.User))
}

func NewUser(user *models.User) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties(userProperties)
	exp.SetExportable(user)
	return exp
}

var (
	userProperties = []types.ExportableProperty{ 
		NewUserProperty("id", func(user *models.User) interface{} {
			return user.GetId()
		}),
		NewUserProperty("email", func(user *models.User) interface{} {
			return user.GetEmail()
		}),
		NewUserProperty("role", func(user *models.User) interface{} {
			return user.GetRole()
		}),
		NewUserProperty("registerDate", func(user *models.User) interface{} {
			return user.GetRegisterDate()
		}),
	}
)
