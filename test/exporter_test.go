package test

import "github.com/byorty/hardcore/test/models"



type UserExportedProperty struct {
	user *models.User
	name string
	closure func(*models.User) interface{}
}

func (u UserExportedProperty) GetName() string {
	return u.name
}

func (u UserExportedProperty) GetValue() interface{} {
	return u.closure(u.user)
}

type UserExporterImpl struct {
	properties []*UserExportedProperty
}

func (u UserExporterImpl) Len() int {
	return len(u.properties)
}

func (u UserExporterImpl) Less(x, y int) bool {
	return u.properties[x] < u.properties[y]
}

func (u UserExporterImpl) Swap(x, y int) {
	u.properties[x], u.properties[y] = u.properties[y], u.properties[x]
}

func (u UserExporterImpl) GetRaw(x int) interface{} {
	return u.properties[x]
}

func (u UserExporterImpl) Get(x int) string {
	return u.properties[x]
}



//func (u *UserExporterImpl) Add(name string, closure func(*models.User) interface{}) *UserExporterImpl {
//	(*u) = append((*u), UserExporterProperty{name, closure})
//	return u
//}
//
//var userExporter = make(UserExporterImpl).
//	Add("id", func(user *models.User) interface{} { return user.GetId()}).
//	Add("email", func(user *models.User) interface{} { return user.GetEmail()}).
//	Add("role", func(user *models.User) interface{} { return user.GetRole().GetId()}).
//	Add("registerDate", func(user *models.User) interface{} { return user.GetRegisterDate()})