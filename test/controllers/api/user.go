package api

import (
    "github.com/byorty/hardcore/types"
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

