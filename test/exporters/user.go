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

func newUserProperty(name string, closure func(user *models.User) interface{}) types.ExportableProperty {
	return &UserPropertyImpl{
		exporter.NewProperty(name),
		closure,
	}
}

func (u UserPropertyImpl) GetValue(model interface{}) interface{} {
	return u.closure(model.(*models.User))
}

func NewUser(user *models.User) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties(userProperties)
	exp.SetExportable(user)
	return exp
}

func NewUsers(users models.Users) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties(userProperties)
	exp.SetExportable(users)
	return exp
}

var (
	userProperties = []types.ExportableProperty{ 
		newUserProperty("id", func(user *models.User) interface{} {
			return user.GetId()
		}),
		newUserProperty("email", func(user *models.User) interface{} {
			return user.GetEmail()
		}),
		newUserProperty("role", func(user *models.User) interface{} {
			return nil
		}),
		newUserProperty("registerDate", func(user *models.User) interface{} {
			return user.GetRegisterDate()
		}),
	}
)
