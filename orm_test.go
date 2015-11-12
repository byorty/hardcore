package hardcore

import (
	"testing"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/byorty/hardcore/db"
	"github.com/byorty/hardcore/proto"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/query"
	"github.com/byorty/hardcore/expr"
	"time"
	"github.com/byorty/hardcore/dao"
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
	userRoles = []types.IntEnum{
		GuestUserRole,
		LoggedUserRole,
	}
	userRoleNames = map[UserRole]string{
		GuestUserRole: "Guest",
		LoggedUserRole: "User",
	}
)

type AutoUser struct {
	Id     int
	Email   string
	Password   string
	Role UserRole
	RegisterDate time.Time
}

func(u *AutoUser) CommonDAO() types.DAO {
	return userDAO
}

func (u *AutoUser) GetRole() UserRole {
	return u.Role
}

type Users []*User

func(u *Users) CommonDAO() types.DAO {
	return userDAO
}

func (u *Users) Proto() types.Proto {
	return userProto
}

type User struct {
	AutoUser
}

func(u *User) DAO() UserDAO {
	return userDAO
}

func (u *User) Proto() types.Proto {
	return userProto
}

type UserDAO struct {
	dao.Sql
}

func (u UserDAO) GetProto() types.Proto {
	return userProto
}

func (t UserDAO) GetDB() string {
	return "default"
}

func (t UserDAO) GetTable() string {
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
	err := row.(types.SqlModelScanner).Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.RegisterDate)
	if err != nil {
		fmt.Println(err)
	}
}

var (
	userDAO UserDAO
	userProto = proto.New().
		Set("Id", proto.NewProperty("id", types.ProtoBasicKind, types.ProtoNoneRelation, true)).
		Set("Email", proto.NewProperty("email", types.ProtoBasicKind, types.ProtoNoneRelation, true)).
		Set("Password", proto.NewProperty("password", types.ProtoBasicKind, types.ProtoNoneRelation, true)).
		Set("Role", proto.NewProperty("role_id", types.ProtoEnumKind, types.ProtoOneToOneRelation, true)).
		Set("RegisterDate", proto.NewProperty("register_date", types.ProtoBasicKind, types.ProtoNoneRelation, true))
)

func TestDB(t *testing.T) {
	sqlDb := db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/kapz?sslmode=disable&client_encoding=utf-8")
	if sqlDb == nil {
		t.Fail()
	}
	db.Pool().
		Add("default", sqlDb)

	user := new(User)
	query.Criteria().Add(expr.Eq("Id", 1)).One(user)
	t.Log(user)
	if user.Id != 1 {
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
	query.Criteria().Add(expr.Eq("Id", 1)).All(&users)
	t.Log(users)
	if existsUsers[0].Id != users[0].Id {
		t.Fail()
	}
}
