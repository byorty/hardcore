package models

import (
	"github.com/byorty/hardcore/proto"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/orm/dao"
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
