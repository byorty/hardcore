package api

import (
    "github.com/byorty/hardcore/types"
    "github.com/byorty/hardcore/test/models"
    "github.com/byorty/hardcore/view"
    "github.com/byorty/hardcore/test/exporters"
)

type User struct { 
}

func NewUser() types.ActionController {
	// implement me
    return new(User)
}


func (u *User) List(page int) types.EncodeView {

	return nil
}

func (u *User) View(user *models.User) types.EncodeView {

	return view.Encode(
        exporters.NewUser(user),
    )
}

