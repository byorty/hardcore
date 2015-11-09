package orm

import (
	"testing"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/proto"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/orm/criteria"
	"github.com/byorty/hardcore/orm/criteria/cond"
	"time"
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

func(u *User) CommonDAO() types.DAO {
	return userDAO
}

//func (a *AutoUser) GetStatus() *UserStatus {
//	if a.Status == nil {
//		a.Status = new(UserStatus).ById(a.statusId).(*UserStatus)
//	}
//	return a.Status;
//}

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

func (t UserDAO) Scan(row interface{}, model interface{}) {
	user := model.(*User)
	err := row.(*sql.Row).Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.RegisterDate)
	if err != nil {
		fmt.Println(err)
	}
}

//const (
//	selectUserById = `SELECT "id", "email", "status_id", "role_id" FROM "user" WHERE "id" = $1`
//)

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
	currentDb := db.Pool().ByDAO(user.DAO())
	fmt.Println(currentDb)
	fmt.Println(sqlDb)

//	currentDb.QueryRow(db.RawQuery{selectUserById, 1}, user.DAO(), user)
	criteria.New(user.DAO()).Add(cond.Eq("Id", 1)).One(user)

//	t.Log(criteria.New(user.DAO()).Add(cond.Eq("Id", 1)).ToNative())
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
}
