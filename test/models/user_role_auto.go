package models

import (
	"github.com/byorty/hardcore/orm/dao"
	"github.com/byorty/hardcore/types"
)

func (u UserRole) GetRawId() interface{} {
	return u.GetId()
}

func (u UserRole) GetId() int {
	return int(u)
}

func (u UserRole) GetName() string {
	return userRoleNames[u]
}

func (u UserRole) DAO() types.IntEnumDAO {
	return userRoleDao
}

type UserRoles []*UserRole

func (u UserRoles) Len() int {
	return len(u)
}

func (u UserRoles) Less(x, y int) bool {
	return u[x].GetId() < u[y].GetId()
}

func (u UserRoles) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u UserRoles) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u UserRoles) Get(x int) *UserRole {
	return u[x]
}

type UserRoleDao struct{}

func (u UserRoleDao) GetList() []types.Named {
	return userRoleList
}

func (u UserRoleDao) ById(id int) types.EnumScanner {
	return dao.NewEnumScanner(id, u)
}

func (u UserRoleDao) Eq(named types.Named, id interface{}) bool {
	return named.(UserRole).GetId() == id
}

func (u UserRoleDao) Scan(src, dest types.Named) {
	enum := dest.(*UserRole)
	(*enum) = src.(UserRole)
}

const (
	LoggedUserRole UserRole = iota
	LoggedAdminRole
)

var (
	userRoleDao  UserRoleDao
	userRoleList = []types.Named{
		LoggedUserRole,
		LoggedAdminRole,
	}
	userRoleNames = map[UserRole]string{
		LoggedUserRole:  "User",
		LoggedAdminRole: "Admin",
	}
)
