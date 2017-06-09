package models

import (
	"github.com/byorty/hardcore/types"
)

type User struct {
	AutoUser
}

func (u *User) HasPosts() bool {
	return u.posts != nil
}

type Users []*User

type UserDao struct {
	AutoUserDao
}

func (u *UserDao) Init(db types.DB) {}
