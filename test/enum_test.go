package test

import (
	"github.com/byorty/hardcore/test/models"
	"testing"
)

func TestEnum(t *testing.T) {
	var role models.UserRole
	role.DAO().ById(int(models.LoggedUserRole)).One(&role)

	t.Log(role)
	if role != models.LoggedUserRole {
		t.Fail()
	}
}
