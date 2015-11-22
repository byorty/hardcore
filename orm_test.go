package hardcore

import (
	"testing"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/byorty/hardcore/db"
	"github.com/byorty/hardcore/proto"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/query/criteria"
	"time"
	"github.com/byorty/hardcore/dao"
	"github.com/byorty/hardcore/proj"
)

type UserRole int

func (u *UserRole) ById(id int) {
	existsId := UserRole(id)
	if _, ok := userRoleNames[existsId]; ok {
		(*u) = existsId
	}
}

func (u UserRole) GetId() int {
	return int(u)
}

func (u UserRole) GetName() string {
	return userRoleNames[u];
}

const (
	GuestUserRole  UserRole = iota
	LoggedUserRole
)

var (
	userRoles = []types.Enum{
		GuestUserRole,
		LoggedUserRole,
	}
	userRoleNames = map[UserRole]string{
		GuestUserRole: "Guest",
		LoggedUserRole: "User",
	}
)

type AutoUser struct {
	id     int
	email   string
	password   string
	role UserRole
	registerDate time.Time
}

func(u *AutoUser) DAO() types.DAO {
	return userDAO
}

func (u *AutoUser) SetId(id int) {
	u.id = id
}

func (u *AutoUser) GetId() int {
	return u.id
}

func (u *AutoUser) SetEmail(email string) {
	u.email = email
}

func (u *AutoUser) GetEmail() string {
	return u.email
}

func (u *AutoUser) SetPassword(password string) {
	u.password = password
}

func (u *AutoUser) GetPassword() string {
	return u.password
}

func (u *AutoUser) SetRole(role UserRole) {
	u.role = role
}

func (u *AutoUser) GetRole() UserRole {
	return u.role
}

func (u *AutoUser) SetRegisterDate(registerDate time.Time) {
	u.registerDate = registerDate
}

func (u *AutoUser) GetRegisterDate() time.Time {
	return u.registerDate
}

type Users []*User

func(u *Users) DAO() types.DAO {
	return userDAO
}

func (u *Users) Proto() types.Proto {
	return userProto
}

type User struct {
	AutoUser
}

func (u *User) Proto() types.Proto {
	return userProto
}

type UserDAO struct {
	dao.Base
}

func (u UserDAO) Proto() types.Proto {
	return userProto
}

func (u UserDAO) GetDB() string {
	return "default"
}

func (u UserDAO) GetTable() string {
	return "user"
}

func (u UserDAO) ScanAll(rows interface{}, model interface{}) {
	users := model.(*Users)
	user := new(User)
	u.Scan(rows, user)
	(*users) = append((*users), user)
}

func (t UserDAO) Scan(row interface{}, model interface{}) {
	user := model.(*User)
	err := row.(types.SqlModelScanner).Scan(&user.id, &user.email, &user.password, &user.role, &user.registerDate)
	if err != nil {
		fmt.Println(err)
	}
}

var (
	userIdSetter UserIdSetter = (*User).SetId
	userIdGetter UserIdGetter = (*User).GetId
	userEmailSetter UserEmailSetter = (*User).SetEmail
	userEmailGetter UserEmailGetter =(*User).GetEmail
	userPasswordSetter UserPasswordSetter = (*User).SetPassword
	userPasswordGetter UserPasswordGetter =(*User).GetPassword
	userUserRoleSetter UserRoleSetter = (*User).SetRole
	userUserRoleGetter UserRoleGetter =(*User).GetRole
	userRegisterDateSetter UserRegisterDateSetter = (*User).SetRegisterDate
	userRegisterDateGetter UserRegisterDateGetter =(*User).GetRegisterDate
	userDAO UserDAO
	userProto = proto.New().
		Set("id", proto.NewProperty("id", types.ProtoBasicKind, types.ProtoNoneRelation, true, userIdSetter, userIdGetter)).
		Set("email", proto.NewProperty("email", types.ProtoBasicKind, types.ProtoNoneRelation, true, userEmailSetter, userEmailGetter)).
		Set("password", proto.NewProperty("password", types.ProtoBasicKind, types.ProtoNoneRelation, true, userPasswordSetter, userPasswordGetter)).
		Set("role", proto.NewProperty("role_id", types.ProtoEnumKind, types.ProtoOneToOneRelation, true, userUserRoleSetter, userUserRoleGetter)).
		Set("registerDate", proto.NewProperty("register_date", types.ProtoBasicKind, types.ProtoNoneRelation, true, userRegisterDateSetter, userRegisterDateGetter))
)

func TestDB(t *testing.T) {
	sqlDb := db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/kapz?sslmode=disable&client_encoding=utf-8")
	if sqlDb == nil {
		t.Fail()
	}
	db.Pool().
		Add("default", sqlDb)

	user := new(User)
	user.DAO().ById(1).One(user)
	t.Log(user)
	if user.id != 1 {
		t.Fail()
	}

	var role UserRole
	role.ById(int(LoggedUserRole))

	t.Log(role)
	if role != LoggedUserRole {
		t.Fail()
	}

	existsUsers := Users{user}
	var users Users
	users.DAO().ByIds([]int{1, 2, 3}).All(&users)
	t.Log(users)
	if len(users) > 0 && existsUsers[0].id != users[0].id {
		t.Fail()
	}

	var count int
	criteria.SelectByDAO(user.DAO()).Add(proj.Count("id")).Custom(&count)
	t.Log(count)
	if count == 0 {
		t.Fail()
	}

	newUser := new(User)
	newUser.SetEmail(fmt.Sprintf("%v@qwerty.com", time.Now().UnixNano()))
	newUser.SetPassword("12345")
	newUser.SetRole(LoggedUserRole)
	newUser.SetRegisterDate(time.Now())
	newUser.DAO().Add(newUser)

	oldPassword := newUser.GetPassword()
	t.Log(newUser)
	t.Log(oldPassword)
	if newUser.GetId() == 0 {
		t.Fail()
	}

	newUser.SetPassword("0987654321")
	newUser.DAO().Save(newUser)

	t.Log(newUser)
	t.Log(newUser.GetPassword())
	if oldPassword == newUser.GetPassword() {
		t.Fail()
	}
}

type UserIdSetter func(*User, int)

func (u UserIdSetter) Call(model interface{}, id interface{}) {
	u(model.(*User), id.(int))
}

type UserIdGetter func(*User) int

func (u UserIdGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserEmailSetter func(*User, string)

func (u UserEmailSetter) Call(model interface{}, id interface{}) {
	u(model.(*User), id.(string))
}

type UserEmailGetter func(*User) string

func (u UserEmailGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserPasswordSetter func(*User, string)

func (u UserPasswordSetter) Call(model interface{}, id interface{}) {
	u(model.(*User), id.(string))
}

type UserPasswordGetter func(*User) string

func (u UserPasswordGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserRoleSetter func(*User, UserRole)

func (u UserRoleSetter) Call(model interface{}, id interface{}) {
	u(model.(*User), id.(UserRole))
}

type UserRoleGetter func(*User) UserRole

func (u UserRoleGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserRegisterDateSetter func(*User, time.Time)

func (u UserRegisterDateSetter) Call(model interface{}, id interface{}) {
	u(model.(*User), id.(time.Time))
}

type UserRegisterDateGetter func(*User) time.Time

func (u UserRegisterDateGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

