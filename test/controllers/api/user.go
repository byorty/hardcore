package api

import (
    "github.com/byorty/hardcore/types"
    "github.com/byorty/hardcore/test/models"
)

type User struct { 
}

func NewUser() types.ActionController {
	// implement me
    return nil
}


func (u *User) List(page int) types.EncodeView {

	return nil
}

func (u *User) View(user *models.User) types.EncodeView {

	return nil
}

