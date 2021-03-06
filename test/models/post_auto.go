package models

import (
	"github.com/byorty/hardcore/orm/dao"
	"github.com/byorty/hardcore/pool"
	"github.com/byorty/hardcore/proto"
	"github.com/byorty/hardcore/query/criteria"
	"github.com/byorty/hardcore/query/expr"
	"github.com/byorty/hardcore/types"
)

type AutoPost struct {
	id          int64
	user        *User
	userId      int64
	name        string
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

func (p *Post) CommonDAO() types.ModelDAO {
	return p.DAO()
}

func (p *Post) KindDAO() types.Int64ModelDAO {
	return p.DAO()
}

func (p *Post) DAO() *PostDao {
	return PostDaoInst()
}

func (p *Post) Proto() types.Proto {
	return postProto
}

func (p Post) IsScanned() bool {
	return p.GetId() != 0
}

func (p Post) GetProtoKind() types.ProtoKind {
	return types.ProtoModelKind
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

func (p *Posts) CommonDAO() types.ModelDAO {
	return p.DAO()
}

func (p *Posts) KindDAO() types.Int64ModelDAO {
	return p.DAO()
}

func (p *Posts) DAO() *PostDao {
	return PostDaoInst()
}

func (p *Posts) Proto() types.Proto {
	return postProto
}

func (p Posts) IsScanned() bool {
	return p.Len() > 0 && p.Get(0).GetId() != 0
}

type AutoPostDao struct {
	dao.Int64Impl
}

func PostDaoInst() *PostDao {
	if postDao == nil {
		postDao = new(PostDao)
	}
	return postDao
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

func (p PostDao) ScanAll(rows interface{}, model interface{}) error {
	var err error
	items := model.(*Posts)
	item := new(Post)
	err = p.Scan(rows, item)
	(*items) = append((*items), item)
	return err
}

func (p PostDao) Scan(row interface{}, model interface{}) error {
	item := model.(*Post)
	return row.(types.DBScanner).Scan(
		&item.id,
		&item.userId,
		&item.name,
		&item.description,
	)
}

func (p *PostDao) Add(model *Post) {
	db := pool.DB().ByDAO(p)
	if db.SupportLastInsertId() {
		p.InsertStmt.Exec(
			model.userId,
			model.name,
			model.description,
		).One(model)
	} else if db.SupportReturningId() {
		p.InsertStmt.Custom(
			model.userId,
			model.name,
			model.description,
		).One(&model.id)
	}
}

func (p *PostDao) Save(model *Post) {
	p.UpdateStmt.Exec(
		model.userId,
		model.name,
		model.description,
		model.id,
	).One(model)
}

func (p *PostDao) Take(model *Post) {
	if model.IsScanned() {
		p.Save(model)
	} else {
		p.Add(model)
	}
}

func (p *PostDao) AutoInit(db types.DB) {
	p.ByIdStmt = db.Prepare(criteria.SelectByDAO(p).And(expr.Eq("id", nil)))
	//p.ByIdsStmt = db.Prepare(criteria.SelectByDAO(p).And(expr.In("id", nil)))
	p.InsertStmt = db.Prepare(criteria.InsertByDao(p))
	p.UpdateStmt = db.Prepare(criteria.UpdateByDAO(p).And(expr.Eq("id", nil)))
}

func postIdSetter(model interface{}, id interface{}) {
	model.(*Post).SetId(id.(int64))
}

func postIdGetter(model interface{}) interface{} {
	return model.(*Post).GetId()
}

func postUserSetter(model interface{}, user interface{}) {
	model.(*Post).SetUser(user.(*User))
}

func postUserGetter(model interface{}) interface{} {
	return model.(*Post).GetUser()
}

func postUserIdSetter(model interface{}, userId interface{}) {
	model.(*Post).SetUserId(userId.(int64))
}

func postUserIdGetter(model interface{}) interface{} {
	return model.(*Post).GetUserId()
}

func postNameSetter(model interface{}, name interface{}) {
	model.(*Post).SetName(name.(string))
}

func postNameGetter(model interface{}) interface{} {
	return model.(*Post).GetName()
}

func postDescriptionSetter(model interface{}, description interface{}) {
	model.(*Post).SetDescription(description.(string))
}

func postDescriptionGetter(model interface{}) interface{} {
	return model.(*Post).GetDescription()
}

var (
	postDao   *PostDao
	postProto = proto.New().
			Set("id", proto.NewProperty("id", types.ProtoInt64Kind, types.ProtoNoneRelation, true, postIdSetter, postIdGetter)).
			Set("user", proto.NewProperty("user", types.ProtoModelKind, types.ProtoOneToOneRelation, true, postUserSetter, postUserGetter)).
			Set("userId", proto.NewProperty("user_id", types.ProtoInt64Kind, types.ProtoNoneRelation, true, postUserIdSetter, postUserIdGetter)).
			Set("name", proto.NewProperty("name", types.ProtoStringKind, types.ProtoNoneRelation, true, postNameSetter, postNameGetter)).
			Set("description", proto.NewProperty("description", types.ProtoStringKind, types.ProtoNoneRelation, true, postDescriptionSetter, postDescriptionGetter))
)
