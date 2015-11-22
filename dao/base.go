package dao

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/db"
	"github.com/byorty/hardcore/query/criteria"
	"github.com/byorty/hardcore/expr"
)

type Base struct {}

func (b Base) Save(model types.StraightMappingModel) {
	idProperty := model.Proto().GetByName("id")
	getter := idProperty.GetGetter()
	criteria.Update().And(expr.Eq("id", getter.Call(model))).One(model)
}

func (b Base) Update(query types.Query, model types.StraightMappingModel) {
	currentDb := db.Pool().ByDAO(model.DAO())
	currentDb.Exec(query, model.DAO(), model)
}

func (b Base) Add(model types.StraightMappingModel) {
	criteria.Insert().One(model)
}

func (b Base) Insert(query types.Query, model types.StraightMappingModel) {
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

func (b Base) All(query types.Query, models types.StraightMappingModel) {
	dao := models.DAO()
	currentDb := db.Pool().ByDAO(dao)
	currentDb.Query(query, dao, models)
}

func (b Base) One(query types.Query, model types.StraightMappingModel) {
	dao := model.DAO()
	currentDb := db.Pool().ByDAO(dao)
	currentDb.QueryRow(query, dao, model)
}

func (b Base) Custom(dao types.DAO, query types.Query, items ...interface{}) {
	currentDb := db.Pool().ByDAO(dao)
	currentDb.Custom(query, items...)
}

func (b Base) ById(id int) types.SelectCriteria {
	return criteria.Select().And(expr.Eq("id", id))
}

func (b Base) ByIds(ids []int) types.SelectCriteria {
	return criteria.Select().And(expr.In("id", ids))
}

//func (b Base) Customs(dao types.DAO, query types.Query, items ...[]interface{}) {
//
//}
