package hardcore

import (
    "testing"
    "github.com/byorty/hardcore/types"
    "github.com/byorty/hardcore/orm/dao"
)

type UserRole1 int

func (u UserRole1) GetRawId() interface{} {
    return u
}

func (u UserRole1) GetId() int {
    return int(u)
}

func (u UserRole1) GetName() string {
    return userRoleMap1[u]
}

func (s UserRole1) DAO() types.IntEnumDAO {
    return userRoleDAO1
}

type UserRoleDAO1 struct {}

func (u UserRoleDAO1) GetList() []types.Named {
    return userRoleList1
}

func (u UserRoleDAO1) ById(id int) types.EnumScanner {
    return dao.NewEnumScanner(id, u)
}

func (u UserRoleDAO1) Eq(named types.Named, id interface{}) bool {
    return named.(UserRole1).GetId() == id
}

func (u UserRoleDAO1) Scan(src, dest types.Named) {
    enum := dest.(*UserRole1)
    (*enum) = src.(UserRole1)
}

const (
    LoggedUserRole1 UserRole1 = iota
    LoggedAdminRole1
)

var (
    userRoleDAO1 UserRoleDAO1
    userRoleList1 = []types.Named{
        LoggedUserRole1,
        LoggedAdminRole1,
    }
    userRoleMap1 = map[UserRole1]string{
        LoggedUserRole1: "Пользователь",
        LoggedAdminRole1: "Администратор",
    }
)

func TestEnum(t *testing.T) {
    var role UserRole1
    role.DAO().ById(1).One(&role)

    t.Log(role.GetId())
    t.Log(LoggedAdminRole1)
    if role != LoggedAdminRole1 {
        t.Fail()
    }
}