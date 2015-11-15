package dao

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/db"
	"github.com/byorty/hardcore/query/criteria"
	"github.com/byorty/hardcore/expr"
)

type Sql struct {}

func (s Sql) Save(model types.StraightMappingModel) {
	idProperty := model.Proto().GetByName("id")
	getter := idProperty.GetGetter()
	criteria.Update().And(expr.Eq("id", getter.Call(model))).One(model)
}

func (s Sql) Update(query types.Query, model types.StraightMappingModel) {
	currentDb := db.Pool().ByDAO(model.DAO())
	currentDb.Exec(query, model.DAO(), model)
}

func (s Sql) Add(model types.StraightMappingModel) {
	criteria.Insert().One(model)
}

func (s Sql) Insert(query types.Query, model types.StraightMappingModel) {
	currentDb := db.Pool().ByDAO(model.DAO())
	if currentDb.SupportLastInsertId() {
		currentDb.Exec(query, model.DAO(), model)
	} else if currentDb.SupportReturningId() {
		var id int
		currentDb.Custom(query, &id)
		setter := model.Proto().GetByName("id").GetSetter()
		setter.Call(model, id)
	}
}

func (s Sql) All(query types.Query, models types.StraightMappingModel) {
	dao := models.DAO()
	currentDb := db.Pool().ByDAO(dao)
	currentDb.Query(query, dao, models)
}

func (s Sql) One(query types.Query, model types.StraightMappingModel) {
	dao := model.DAO()
	currentDb := db.Pool().ByDAO(dao)
	currentDb.QueryRow(query, dao, model)
}

func (s Sql) Custom(dao types.DAO, query types.Query, items ...interface{}) {
	currentDb := db.Pool().ByDAO(dao)
	currentDb.Custom(query, items...)
}

func (s Sql) ById(id int) types.SelectCriteria {
	return criteria.Select().And(expr.Eq("id", id))
}

func (s Sql) ByIds(ids []int) types.SelectCriteria {
	return criteria.Select().And(expr.In("id", ids))
}

//func (s Sql) Customs(dao types.DAO, query types.Query, items ...[]interface{}) {
//
//}
