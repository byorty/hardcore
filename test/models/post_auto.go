package models

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/orm/dao"
	"github.com/byorty/hardcore/proto"
)

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

func(p *Post) KindDAO() types.Int64ModelDAO {
	return postDao
}

func(p *Post) DAO() PostDao {
	return postDao
}

func (p *Post) Proto() types.Proto {
	return postProto
}

func (p Posts) Len() int {
	return len(p)
}

func (p Posts) Less(x, y int) bool {
	return p[x].GetId() < p[y].GetId()
}

func (p Posts) Swap(x, y int) {
	p[x], p[y] = p[y], p[x]
}

func (p Posts) GetRaw(x int) interface{} {
	return p.Get(x)
}

func (p Posts) Get(x int) *Post {
	return p[x]
}

func(p *Posts) CommonDAO() types.ModelDAO {
	return postDao
}

func(p *Posts) KindDAO() types.Int64ModelDAO {
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
	row.(types.DBScanner).Scan(
		&item.id,
		&item.userId,
		&item.name,
		&item.description,
	)
}

func postIdSetter (model interface{}, id interface{}) {
	model.(*Post).SetId(id.(int64))
}

func postIdGetter (model interface{}) interface{} {
	return model.(*Post).GetId()
}

func postUserSetter (model interface{}, user interface{}) {
	model.(*Post).SetUser(user.(*User))
}

func postUserGetter (model interface{}) interface{} {
	return model.(*Post).GetUser()
}

func postUserIdSetter (model interface{}, userId interface{}) {
	model.(*Post).SetUserId(userId.(int64))
}

func postUserIdGetter (model interface{}) interface{} {
	return model.(*Post).GetUserId()
}

func postNameSetter (model interface{}, name interface{}) {
	model.(*Post).SetName(name.(string))
}

func postNameGetter (model interface{}) interface{} {
	return model.(*Post).GetName()
}

func postDescriptionSetter (model interface{}, description interface{}) {
	model.(*Post).SetDescription(description.(string))
}

func postDescriptionGetter (model interface{}) interface{} {
	return model.(*Post).GetDescription()
}

var (
	postDao PostDao
	postProto = proto.New().
		Set("id", proto.NewProperty("id", types.ProtoInt64Kind, types.ProtoNoneRelation, true, postIdSetter, postIdGetter)).
		Set("user", proto.NewProperty("user", types.ProtoModelKind, types.ProtoOneToOneRelation, true, postUserSetter, postUserGetter)).
		Set("userId", proto.NewProperty("user_id", types.ProtoInt64Kind, types.ProtoNoneRelation, true, postUserIdSetter, postUserIdGetter)).
		Set("name", proto.NewProperty("name", types.ProtoStringKind, types.ProtoNoneRelation, true, postNameSetter, postNameGetter)).
		Set("description", proto.NewProperty("description", types.ProtoStringKind, types.ProtoNoneRelation, true, postDescriptionSetter, postDescriptionGetter))
)
