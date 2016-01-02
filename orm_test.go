package hardcore

import (
	"testing"
	_ "github.com/lib/pq"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/proto"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/query/criteria"
	"time"
	"github.com/byorty/hardcore/orm/dao"
	"github.com/byorty/hardcore/query/proj"
	"fmt"
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
		Set("id", proto.NewProperty("id", types.ProtoBasicKind, types.ProtoNoneRelation, true, userIdSetter, userIdGetter)).
		Set("email", proto.NewProperty("email", types.ProtoBasicKind, types.ProtoNoneRelation, true, userEmailSetter, userEmailGetter)).
		Set("password", proto.NewProperty("password", types.ProtoBasicKind, types.ProtoNoneRelation, true, userPasswordSetter, userPasswordGetter)).
		Set("role", proto.NewProperty("role", types.ProtoEnumKind, types.ProtoOneToOneRelation, true, userRoleSetter, userRoleGetter)).
		Set("roleId", proto.NewProperty("role_id", types.ProtoBasicKind, types.ProtoNoneRelation, true, userRoleIdSetter, userRoleIdGetter)).
		Set("registerDate", proto.NewProperty("register_date", types.ProtoBasicKind, types.ProtoNoneRelation, false, userRegisterDateSetter, userRegisterDateGetter)).
		Set("posts", proto.NewProperty("posts", types.ProtoModelKind, types.ProtoOneToManyRelation, true, userPostsSetter, userPostsGetter))
)

type Post struct {
	AutoPost
}

type Posts []*Post

type PostDao struct {
	AutoPostDao
}

type AutoPost struct {
	id int64
	user *User
	userId int64
	name string
	description string
}

func (p Post) GetId() int64 {
	return p.id
}

func (p *Post) SetId(id int64) *Post {
	p.id = id
	return p
}
func (p Post) GetUser() *User {
	if p.user == nil {
		var user User
		user.DAO().ById(p.userId).One(&user)
		p.user = &user
	}
	return p.user
}

func (p *Post) SetUser(user *User) *Post {
	p.user = user
	p.SetUserId(user.GetId())
	return p
}
func (p Post) GetUserId() int64 {
	return p.userId
}

func (p *Post) SetUserId(userId int64) *Post {
	p.userId = userId
	return p
}
func (p Post) GetName() string {
	return p.name
}

func (p *Post) SetName(name string) *Post {
	p.name = name
	return p
}
func (p Post) GetDescription() string {
	return p.description
}

func (p *Post) SetDescription(description string) *Post {
	p.description = description
	return p
}

func(p *Post) CommonDAO() types.ModelDAO {
	return postDao
}

func(p *Post) DAO() PostDao {
	return postDao
}

func (p *Post) Proto() types.Proto {
	return postProto
}

func (p Posts) Get(i int) *Post {
	return p[i]
}

func (p Posts) Len() int {
	return len(p)
}

func(p *Posts) CommonDAO() types.ModelDAO {
	return postDao
}

func(p *Posts) DAO() PostDao {
	return postDao
}

func (p *Posts) Proto() types.Proto {
	return postProto
}

type AutoPostDao struct {
	dao.Int64Impl
}

func (p PostDao) GetDB() string {
	return "default"
}

func (p PostDao) GetTable() string {
	return "post"
}

func (p PostDao) Proto() types.Proto {
	return postProto
}

func (p PostDao) ScanAll(rows interface{}, model interface{}) {
	items := model.(*Posts)
	item := new(Post)
	p.Scan(rows, item)
	(*items) = append((*items), item)
}

func (p PostDao) Scan(row interface{}, model interface{}) {
	item := model.(*Post)
	row.(types.SqlModelScanner).Scan(
		&item.id,
		&item.userId,
		&item.name,
		&item.description,
	)
}

type PostIdSetter func(*Post, int64) *Post

func (p PostIdSetter) Call(model interface{}, id interface{}) {
	p(model.(*Post), id.(int64))
}

type PostIdGetter func(*Post) int64

func (p PostIdGetter) Call(model interface{}) interface{} {
	return p(model.(*Post))
}

type PostUserSetter func(*Post, *User) *Post

func (p PostUserSetter) Call(model interface{}, user interface{}) {
	p(model.(*Post), user.(*User))
}

type PostUserGetter func(*Post) *User

func (p PostUserGetter) Call(model interface{}) interface{} {
	return p(model.(*Post))
}

type PostUserIdSetter func(*Post, int64) *Post

func (p PostUserIdSetter) Call(model interface{}, userId interface{}) {
	p(model.(*Post), userId.(int64))
}

type PostUserIdGetter func(*Post) int64

func (p PostUserIdGetter) Call(model interface{}) interface{} {
	return p(model.(*Post))
}

type PostNameSetter func(*Post, string) *Post

func (p PostNameSetter) Call(model interface{}, name interface{}) {
	p(model.(*Post), name.(string))
}

type PostNameGetter func(*Post) string

func (p PostNameGetter) Call(model interface{}) interface{} {
	return p(model.(*Post))
}

type PostDescriptionSetter func(*Post, string) *Post

func (p PostDescriptionSetter) Call(model interface{}, description interface{}) {
	p(model.(*Post), description.(string))
}

type PostDescriptionGetter func(*Post) string

func (p PostDescriptionGetter) Call(model interface{}) interface{} {
	return p(model.(*Post))
}

var (
	postIdSetter PostIdSetter = (*Post).SetId
	postIdGetter PostIdGetter = (*Post).GetId
	postUserSetter PostUserSetter = (*Post).SetUser
	postUserGetter PostUserGetter = (*Post).GetUser
	postUserIdSetter PostUserIdSetter = (*Post).SetUserId
	postUserIdGetter PostUserIdGetter = (*Post).GetUserId
	postNameSetter PostNameSetter = (*Post).SetName
	postNameGetter PostNameGetter = (*Post).GetName
	postDescriptionSetter PostDescriptionSetter = (*Post).SetDescription
	postDescriptionGetter PostDescriptionGetter = (*Post).GetDescription
	postDao PostDao
	postProto = proto.New().
		Set("id", proto.NewProperty("id", types.ProtoBasicKind, types.ProtoNoneRelation, true, postIdSetter, postIdGetter)).
		Set("user", proto.NewProperty("user", types.ProtoModelKind, types.ProtoOneToOneRelation, true, postUserSetter, postUserGetter)).
		Set("userId", proto.NewProperty("user_id", types.ProtoBasicKind, types.ProtoNoneRelation, true, postUserIdSetter, postUserIdGetter)).
		Set("name", proto.NewProperty("name", types.ProtoBasicKind, types.ProtoNoneRelation, true, postNameSetter, postNameGetter)).
		Set("description", proto.NewProperty("description", types.ProtoBasicKind, types.ProtoNoneRelation, true, postDescriptionSetter, postDescriptionGetter))
)

func TestDB(t *testing.T) {
	sqlDb := db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/hardcore?sslmode=disable&client_encoding=utf-8")
	if sqlDb == nil {
		t.Fail()
	}
	db.Pool().
		Add("default", sqlDb)

	var user User
	user.DAO().ById(1).One(&user)
	t.Log(user)
	if user.id != 1 {
		t.Fail()
	}

	var role UserRole
	role.DAO().ById(int(LoggedUserRole))

	t.Log(role)
	if role != LoggedUserRole {
		t.Fail()
	}

//	existsUsers := Users{user}
//	var users Users
//	users.DAO().ByIds([]int64{1, 2, 3}).All(&users)
//	t.Log(users)
//	if len(users) > 0 && existsUsers[0].id != users[0].id {
//		t.Fail()
//	}

	var count int
	criteria.SelectByDAO(user.DAO()).Add(proj.Count("id")).Custom(&count)
	t.Log(count)
	if count == 0 {
		t.Fail()
	}

	newUser := new(User)
	newUser.SetEmail(fmt.Sprintf("%v@qwerty.com", time.Now().UnixNano()))
	newUser.SetPassword("12345")
	newUser.SetRole(&role)
	newUser.SetRegisterDate(time.Now())
	newUser.DAO().Add(newUser)

	oldPassword := newUser.GetPassword()
	t.Log(newUser)
	t.Log(oldPassword)
	if newUser.GetId() == 0 {
		t.Fail()
	}

	newUser.SetPassword("1234567890")
	newUser.DAO().Save(newUser)

	t.Log(newUser)
	t.Log(newUser.GetPassword())
	if oldPassword == newUser.GetPassword() {
		t.Fail()
	}
}
