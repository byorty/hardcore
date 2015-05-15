package meta

import (
	"testing"
)

type User struct {
	Id             int         `orm:"auto;pk"`
	Email          string      `orm:"unique"`
	Password       string
	Firstname      string      `orm:"null"`
	Lastname       string      `orm:"null"`
//	RoleId         int         `orm:"index"                       json:"-"`
//	Role           *UserRole   `orm:"-"                           json:"role"`
//	StatusId       int         `orm:"index"                       json:"-"`
//	Status         *UserStatus `orm:"-"`
//	RegisterDate   time.Time   `orm:"auto_now_add;type(datetime)"`
//	LoginDate      time.Time   `orm:"type(datetime);null"`
//	LoginIP        string      `orm:"null"`
//	LoginUserAgent string      `orm:"null"`
}

func TestModels(t *testing.T) {

}
