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

func (u UserRole) GetProtoKind() types.ProtoKind {
	return types.ProtoIntEnumKind
}

type UserRoles []UserRole

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

func (u UserRoles) Get(x int) UserRole {
	return u[x]
}

func (u UserRoleDao) ById(id int) types.EnumScanner {
	scanner := new(_UserRoleScanner)
	scanner.id = id
	return scanner
}

type _UserRoleScanner struct {
	dao.EnumScannerImpl
	dest *UserRole
	id   int
}

func (u *_UserRoleScanner) Scan(i int) {
	(*u.dest) = userRoleList.Get(i)
}

func (u *_UserRoleScanner) One(named types.Named) {
	u.dest = named.(*UserRole)
	u.Find(u, userRoleList)
}

func (u *_UserRoleScanner) Eq(i int) bool {
	return u.id == userRoleList.Get(i).GetId()
}

const (
	LoggedUserRole UserRole = iota + 1
	LoggedAdminRole
)

var (
	userRoleDao  UserRoleDao
	userRoleList = UserRoles{
		LoggedUserRole,
		LoggedAdminRole,
	}
	userRoleNames = map[UserRole]string{
		LoggedUserRole:  "User",
		LoggedAdminRole: "Admin",
	}
)
