package test

import "github.com/byorty/hardcore/test/models"

type UserExporterProperty struct {
	name string
	closure func(*models.User) interface{}
}

type UserExporterImpl []UserExporterProperty



func (u *UserExporterImpl) Add(name string, closure func(*models.User) interface{}) *UserExporterImpl {
	(*u) = append((*u), UserExporterProperty{name, closure})
	return u
}

var userExporter = make(UserExporterImpl).
	Add("id", func(user *models.User) interface{} { return user.GetId()}).
	Add("email", func(user *models.User) interface{} { return user.GetEmail()}).
	Add("role", func(user *models.User) interface{} { return user.GetRole().GetId()}).
	Add("registerDate", func(user *models.User) interface{} { return user.GetRegisterDate()})