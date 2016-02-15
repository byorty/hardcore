package models

import (
	"github.com/byorty/hardcore/types"
)

type User struct {
	AutoUser
}

type Users []*User

type UserDao struct {
	AutoUserDao
}

func (u *UserDao) Init(db types.DB) {}
