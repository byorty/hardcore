package meta

import (
	"testing"
	"github.com/byorty/hardcore/orm/db"
	_ "github.com/lib/pq"
	"time"
//	"fmt"
	"database/sql"
)

const (
	USER_GUEST_ROLE    int = 1 << iota
	USER_CUSTOMER_ROLE
	USER_MANAGER_ROLE
)

type UserRole struct {
	Id   int    `enum:"USER_GUEST_ROLE, USER_CUSTOMER_ROLE, USER_MANAGER_ROLE"`
	Name string `enum:"Гость, Покупатель, Менеджер"`
}

type User struct {
	Id             int
	Email          string
	Password       string
	Role           *UserRole
	RegisterDate   time.Time
//	Firstname      string      `orm:"null"`
//	Lastname       string      `orm:"null"`
//	RoleId         int         `orm:"index"                       json:"-"`
//	StatusId       int         `orm:"index"                       json:"-"`
//	Status         *UserStatus `orm:"-"`
//	RegisterDate   time.Time   `orm:"auto_now_add;type(datetime)"`
//	LoginDate      time.Time   `orm:"type(datetime);null"`
//	LoginIP        string      `orm:"null"`
//	LoginUserAgent string      `orm:"null"`
}

type UserAlter struct {
	Id         int                `orm:"pk"`
	User       *User              `orm:"relation:oneToOne;required;"`
//	PropertyId int                `orm:"index"`
//	Property   *UserAlterProperty `orm:"-"`
//	Value      string             `orm:"null"`
//	Token      string             `orm:"null"`
//	CreateDate time.Time          `orm:"auto_now_add;type(datetime)"`
//	TryCount   int                `orm:"default(0)"`
}

const (
	selectUserById = `SELECT "id", "email", "password", "role_id", "register_date" FROM "user" WHERE "id" = $1`
	selectAlterByUserId = `SELECT "id", "user_id" FROM "user_alter" WHERE "user_id" = $1`
)

func TestModels(t *testing.T) {
	sqlDb := db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/kapz?sslmode=disable&client_encoding=utf-8")
	if sqlDb == nil {
		t.Fail()
	}
	db.Pool().Add("default", sqlDb)

	builder := new(modelBuilder)
	builder.Build(
		ModelByDb{
			dbName: "default",
			models: []interface{}{new(User), new(UserAlter)},
		},
	)

//	user := NewUser()
//	currentDb := db.Pool().ByDAO(user.DAO())
//	currentDb.QueryRow(db.RawQuery{selectUserById, 1}, user.DAO(), user)
//	fmt.Println(user, user.Role)
//
//	alter := NewUserAlter()
//	currentDb.QueryRow(db.RawQuery{selectAlterByUserId, 1}, alter.DAO(), alter)
//	fmt.Println(alter, alter.User)
}

func BenchmarkDB(b *testing.B) {
	sqlDb := db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/kapz?sslmode=disable&client_encoding=utf-8")

	db.Pool().Add("default", sqlDb)

//	for i := 0;i < 10;i++ {
//		user := NewUser()
//		currentDb := db.Pool().ByDAO(user.DAO())
//		currentDb.QueryRow(db.RawQuery{selectUserById, 1}, user.DAO(), user)
//	}
}

func BenchmarkDB2(b *testing.B) {
	currentDb, _ := sql.Open("postgres", "postgres://byorty:MK99rc@localhost:5432/kapz?sslmode=disable&client_encoding=utf-8")
	for i := 0;i < 10;i++ {
		user := &User{}
		row := currentDb.QueryRow(selectUserById, 1)
		row.Scan(&user.Id, &user.Email, &user.Password, &user.RegisterDate)
	}
}
