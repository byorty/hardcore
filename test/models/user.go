package models

import (
	"github.com/byorty/hardcore/proto"
	"github.com/byorty/hardcore/types"
	"time"
	"github.com/byorty/hardcore/orm/dao"
)

type UserRole int

func (u UserRole) GetRawId() interface{} {
	return u
}

func (u UserRole) GetId() int {
	return int(u)
}

func (u UserRole) GetName() string {
	return userRoleMap[u]
}

func (u UserRole) DAO() types.IntEnumDAO {
	return userRoleDao
}

type UserRoles []*UserRole

func (u UserRoles) Get(i int) *UserRole {
	return u[i]
}

func (u UserRoles) Len() int {
	return len(u)
}

type UserRoleDao struct {}

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
	userRoleDao UserRoleDao
	userRoleList = []types.Named{
		LoggedUserRole,
		LoggedAdminRole,
	}
	userRoleMap = map[UserRole]string{
		LoggedUserRole: "Пользователь",
		LoggedAdminRole: "Администратор",
	}
)

type User struct {
	AutoUser
}

type Users []*User

type UserDao struct {
	AutoUserDao
}

type AutoUser struct {
	id int64
	email string
	password string
	role *UserRole
	roleId int
	registerDate time.Time
	posts Posts
}

func (u User) GetId() int64 {
	return u.id
}

func (u *User) SetId(id int64) *User {
	u.id = id
	return u
}
func (u User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) *User {
	u.email = email
	return u
}
func (u User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) *User {
	u.password = password
	return u
}
func (u User) GetRole() *UserRole {
	if u.role == nil {
		var role UserRole
		role.DAO().ById(u.roleId).One(&role)
		u.role = &role
	}
	return u.role
}

func (u *User) SetRole(role *UserRole) *User {
	u.role = role
	u.SetRoleId(role.GetId())
	return u
}
func (u User) GetRoleId() int {
	return u.roleId
}

func (u *User) SetRoleId(roleId int) *User {
	u.roleId = roleId
	return u
}
func (u User) GetRegisterDate() time.Time {
	return u.registerDate
}

func (u *User) SetRegisterDate(registerDate time.Time) *User {
	u.registerDate = registerDate
	return u
}
func (u User) GetPosts() Posts {
	if u.posts == nil {
		dao.NewInt64OneToMany("user").ById(u.GetId()).All(&u.posts)
	}
	return u.posts
}

func (u *User) SetPosts(posts Posts) *User {
	u.posts = posts
	return u
}

func(u *User) CommonDAO() types.ModelDAO {
	return userDao
}

func(u *User) DAO() UserDao {
	return userDao
}

func (u *User) Proto() types.Proto {
	return userProto
}

func (u Users) Get(i int) *User {
	return u[i]
}

func (u Users) Len() int {
	return len(u)
}

func(u *Users) CommonDAO() types.ModelDAO {
	return userDao
}

func(u *Users) DAO() UserDao {
	return userDao
}

func (u *Users) Proto() types.Proto {
	return userProto
}

type AutoUserDao struct {
	dao.Int64Impl
}

func (u UserDao) GetDB() string {
	return "default"
}

func (u UserDao) GetTable() string {
	return "user"
}

func (u UserDao) Proto() types.Proto {
	return userProto
}

func (u UserDao) ScanAll(rows interface{}, model interface{}) {
	items := model.(*Users)
	item := new(User)
	u.Scan(rows, item)
	(*items) = append((*items), item)
}

func (u UserDao) Scan(row interface{}, model interface{}) {
	item := model.(*User)
	row.(types.SqlModelScanner).Scan(
		&item.id,
		&item.email,
		&item.password,
		&item.roleId,
		&item.registerDate,
	)
}

type UserIdSetter func(*User, int64) *User

func (u UserIdSetter) Call(model interface{}, id interface{}) {
	u(model.(*User), id.(int64))
}

type UserIdGetter func(*User) int64

func (u UserIdGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserEmailSetter func(*User, string) *User

func (u UserEmailSetter) Call(model interface{}, email interface{}) {
	u(model.(*User), email.(string))
}

type UserEmailGetter func(*User) string

func (u UserEmailGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserPasswordSetter func(*User, string) *User

func (u UserPasswordSetter) Call(model interface{}, password interface{}) {
	u(model.(*User), password.(string))
}

type UserPasswordGetter func(*User) string

func (u UserPasswordGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserRoleSetter func(*User, *UserRole) *User

func (u UserRoleSetter) Call(model interface{}, role interface{}) {
	u(model.(*User), role.(*UserRole))
}

type UserRoleGetter func(*User) *UserRole

func (u UserRoleGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserRoleIdSetter func(*User, int) *User

func (u UserRoleIdSetter) Call(model interface{}, roleId interface{}) {
	u(model.(*User), roleId.(int))
}

type UserRoleIdGetter func(*User) int

func (u UserRoleIdGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserRegisterDateSetter func(*User, time.Time) *User

func (u UserRegisterDateSetter) Call(model interface{}, registerDate interface{}) {
	u(model.(*User), registerDate.(time.Time))
}

type UserRegisterDateGetter func(*User) time.Time

func (u UserRegisterDateGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

type UserPostsSetter func(*User, Posts) *User

func (u UserPostsSetter) Call(model interface{}, posts interface{}) {
	u(model.(*User), posts.(Posts))
}

type UserPostsGetter func(*User) Posts

func (u UserPostsGetter) Call(model interface{}) interface{} {
	return u(model.(*User))
}

var (
	userIdSetter UserIdSetter = (*User).SetId
	userIdGetter UserIdGetter = (*User).GetId
	userEmailSetter UserEmailSetter = (*User).SetEmail
	userEmailGetter UserEmailGetter = (*User).GetEmail
	userPasswordSetter UserPasswordSetter = (*User).SetPassword
	userPasswordGetter UserPasswordGetter = (*User).GetPassword
	userRoleSetter UserRoleSetter = (*User).SetRole
	userRoleGetter UserRoleGetter = (*User).GetRole
	userRoleIdSetter UserRoleIdSetter = (*User).SetRoleId
	userRoleIdGetter UserRoleIdGetter = (*User).GetRoleId
	userRegisterDateSetter UserRegisterDateSetter = (*User).SetRegisterDate
	userRegisterDateGetter UserRegisterDateGetter = (*User).GetRegisterDate
	userPostsSetter UserPostsSetter = (*User).SetPosts
	userPostsGetter UserPostsGetter = (*User).GetPosts
	userDao UserDao
	userProto = proto.New().
		Set("id", proto.NewProperty("id", types.ProtoInt64Kind, types.ProtoNoneRelation, true, userIdSetter, userIdGetter)).
		Set("email", proto.NewProperty("email", types.ProtoStringKind, types.ProtoNoneRelation, true, userEmailSetter, userEmailGetter)).
		Set("password", proto.NewProperty("password", types.ProtoStringKind, types.ProtoNoneRelation, true, userPasswordSetter, userPasswordGetter)).
		Set("role", proto.NewProperty("role", types.ProtoEnumKind, types.ProtoOneToOneRelation, true, userRoleSetter, userRoleGetter)).
		Set("roleId", proto.NewProperty("role_id", types.ProtoIntKind, types.ProtoNoneRelation, true, userRoleIdSetter, userRoleIdGetter)).
		Set("registerDate", proto.NewProperty("register_date", types.ProtoTimeKind, types.ProtoNoneRelation, false, userRegisterDateSetter, userRegisterDateGetter)).
		Set("posts", proto.NewProperty("posts", types.ProtoModelKind, types.ProtoOneToManyRelation, true, userPostsSetter, userPostsGetter))
)
