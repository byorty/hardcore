package meta

import (
	"testing"
	_ "github.com/lib/pq"
	"github.com/byorty/hardcore/orm/db"
)

func TestBuilder1(t *testing.T) {
	sqlDb := db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/kapz?sslmode=disable&client_encoding=utf-8")
	if sqlDb == nil {
		t.Fail()
	}
	db.Pool().Add("default", sqlDb)

	RegisterEnums(
		&SomeStatus{},
		&UserRole{},
	)
	RegisterControllers(
		&SomeController{},
	)

	RegisterModels("default", &User{}, &UserAlter{})
	Build()
}
