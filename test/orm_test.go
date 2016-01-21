package test

import (
	"testing"
	"github.com/byorty/hardcore/query/criteria"
	"time"
	"github.com/byorty/hardcore/query/proj"
	"fmt"
	"github.com/byorty/hardcore/test/models"
)


func TestDB(t *testing.T) {
	var user models.User
	user.DAO().ById(1).One(&user)
	t.Log(user)
	if user.GetId() != 1 {
		t.Fail()
	}

	var role models.UserRole
	role.DAO().ById(int(models.LoggedUserRole))

	t.Log(role)
	if role != models.LoggedUserRole {
		t.Fail()
	}

//	existsUsers := Users{user}
//	var users Users
//	users.DAO().ByIds([]int64{1, 2, 3}).All(&users)
//	t.Log(users)
//	if len(users) > 0 && existsUsers[0].id != users[0].id {
//		t.Fail()
//	}

	var count int
	criteria.SelectByDAO(user.DAO()).Add(proj.Count("id")).Custom(&count)
	t.Log(count)
	if count == 0 {
		t.Fail()
	}

	newUser := new(models.User)
	newUser.SetEmail(fmt.Sprintf("%v@qwerty.com", time.Now().UnixNano()))
	newUser.SetPassword("12345")
	newUser.SetRole(&role)
	newUser.SetRegisterDate(time.Now())
	newUser.DAO().Add(newUser)

	oldPassword := newUser.GetPassword()
	t.Log(newUser)
	t.Log(oldPassword)
	if newUser.GetId() == 0 {
		t.Fail()
	}

	newUser.SetPassword("1234567890")
	newUser.DAO().Save(newUser)

	t.Log(newUser)
	t.Log(newUser.GetPassword())
	if oldPassword == newUser.GetPassword() {
		t.Fail()
	}
}
