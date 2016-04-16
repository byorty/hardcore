package models

import (
	"github.com/byorty/hardcore/orm/dao"
	"github.com/byorty/hardcore/proto"
	"github.com/byorty/hardcore/types"
	"time"
)

type AutoUser struct {
	id           int64
	email        string
	password     string
	role         *UserRole
	roleId       int
	registerDate time.Time
	posts        Posts
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
func (u User) GetRole() UserRole {
	if u.role == nil {
		var role UserRole
		role.DAO().ById(u.roleId).One(&role)
		u.role = &role
	}
	return *(u.role)
}

func (u *User) SetRole(role UserRole) *User {
	u.role = &role
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

func (u *User) CommonDAO() types.ModelDAO {
	return u.DAO()
}

func (u *User) KindDAO() types.Int64ModelDAO {
	return u.DAO()
}

func (u *User) DAO() *UserDao {
	return UserDaoInst()
}

func (u *User) Proto() types.Proto {
	return userProto
}

func (u User) IsScanned() bool {
	return u.GetId() != 0
}

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(x, y int) bool {
	return u[x].GetId() < u[y].GetId()
}

func (u Users) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Users) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Users) Get(x int) *User {
	return u[x]
}

func (u *Users) CommonDAO() types.ModelDAO {
	return u.DAO()
}

func (u *Users) KindDAO() types.Int64ModelDAO {
	return u.DAO()
}

func (u *Users) DAO() *UserDao {
	return UserDaoInst()
}

func (u *Users) Proto() types.Proto {
	return userProto
}

func (u Users) IsScanned() bool {
	return u.Len() > 0 && u.Get(0).GetId() != 0
}

type AutoUserDao struct {
	dao.Int64Impl
}

func UserDaoInst() *UserDao {
	if userDao == nil {
		userDao = new(UserDao)
	}
	return userDao
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

func (u UserDao) ScanAll(rows interface{}, model interface{}) error {
	var err error
	items := model.(*Users)
	item := new(User)
	err = u.Scan(rows, item)
	(*items) = append((*items), item)
	return err
}

func (u UserDao) Scan(row interface{}, model interface{}) error {
	item := model.(*User)
	return row.(types.DBScanner).Scan(
		&item.id,
		&item.email,
		&item.password,
		&item.roleId,
		&item.registerDate,
	)
}

func (u *UserDao) AutoInit(db types.DB) {

}

func userIdSetter(model interface{}, id interface{}) {
	model.(*User).SetId(id.(int64))
}

func userIdGetter(model interface{}) interface{} {
	return model.(*User).GetId()
}

func userEmailSetter(model interface{}, email interface{}) {
	model.(*User).SetEmail(email.(string))
}

func userEmailGetter(model interface{}) interface{} {
	return model.(*User).GetEmail()
}

func userPasswordSetter(model interface{}, password interface{}) {
	model.(*User).SetPassword(password.(string))
}

func userPasswordGetter(model interface{}) interface{} {
	return model.(*User).GetPassword()
}

func userRoleSetter(model interface{}, role interface{}) {
	model.(*User).SetRole(role.(UserRole))
}

func userRoleGetter(model interface{}) interface{} {
	return model.(*User).GetRole()
}

func userRoleIdSetter(model interface{}, roleId interface{}) {
	model.(*User).SetRoleId(roleId.(int))
}

func userRoleIdGetter(model interface{}) interface{} {
	return model.(*User).GetRoleId()
}

func userRegisterDateSetter(model interface{}, registerDate interface{}) {
	model.(*User).SetRegisterDate(registerDate.(time.Time))
}

func userRegisterDateGetter(model interface{}) interface{} {
	return model.(*User).GetRegisterDate()
}

func userPostsSetter(model interface{}, posts interface{}) {
	model.(*User).SetPosts(posts.(Posts))
}

func userPostsGetter(model interface{}) interface{} {
	return model.(*User).GetPosts()
}

var (
	userDao   *UserDao
	userProto = proto.New().
			Set("id", proto.NewProperty("id", types.ProtoInt64Kind, types.ProtoNoneRelation, true, userIdSetter, userIdGetter)).
			Set("email", proto.NewProperty("email", types.ProtoStringKind, types.ProtoNoneRelation, true, userEmailSetter, userEmailGetter)).
			Set("password", proto.NewProperty("password", types.ProtoStringKind, types.ProtoNoneRelation, true, userPasswordSetter, userPasswordGetter)).
			Set("role", proto.NewProperty("role", types.ProtoEnumKind, types.ProtoOneToOneRelation, true, userRoleSetter, userRoleGetter)).
			Set("roleId", proto.NewProperty("role_id", types.ProtoIntKind, types.ProtoNoneRelation, true, userRoleIdSetter, userRoleIdGetter)).
			Set("registerDate", proto.NewProperty("register_date", types.ProtoTimeKind, types.ProtoNoneRelation, false, userRegisterDateSetter, userRegisterDateGetter)).
			Set("posts", proto.NewProperty("posts", types.ProtoSliceKind, types.ProtoOneToManyRelation, true, userPostsSetter, userPostsGetter))
)
