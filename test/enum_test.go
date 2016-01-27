package test

import (
    "testing"
    "github.com/byorty/hardcore/test/models"
    "github.com/byorty/hardcore/types"
)

func TestEnum(t *testing.T) {
    var role models.UserRole
    role.DAO().ById(int(models.LoggedUserRole))

    t.Log(role)
    if role != models.LoggedUserRole {
        t.Fail()
    }
}