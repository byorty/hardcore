package models

type User struct {
	AutoUser
}

type Users []*User

type UserDao struct {
	AutoUserDao
}
