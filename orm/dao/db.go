package dao

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/query/criteria"
	"github.com/byorty/hardcore/query/expr"
)

type BaseModelImpl struct {}

func (b BaseModelImpl) Save(model types.Model) {
	idProperty := model.Proto().GetByName("id")
	getter := idProperty.GetGetter()
	criteria.Update().And(expr.Eq("id", getter.Call(model))).One(model)
}

func (b BaseModelImpl) Update(query types.Query, model types.Model) {
	currentDb := db.Pool().ByDAO(model.CommonDAO())
	currentDb.Exec(query, model.CommonDAO(), model)
}

func (b BaseModelImpl) Add(model types.Model) {
	criteria.Insert().One(model)
}

func (b BaseModelImpl) Insert(query types.Query, model types.Model) {
	currentDb := db.Pool().ByDAO(model.CommonDAO())
	if currentDb.SupportLastInsertId() {
		currentDb.Exec(query, model.CommonDAO(), model)
	} else if currentDb.SupportReturningId() {
		var id int
		currentDb.Custom(query, &id)
		setter := model.Proto().GetByName("id").GetSetter()
		setter.Call(model, id)
	}
}

func (b BaseModelImpl) All(query types.Query, models types.Model) {
	dao := models.CommonDAO()
	currentDb := db.Pool().ByDAO(dao)
	currentDb.Query(query, dao, models)
}

func (b BaseModelImpl) One(query types.Query, model types.Model) {
	dao := model.CommonDAO()
	currentDb := db.Pool().ByDAO(dao)
	currentDb.QueryRow(query, dao, model)
}

func (b BaseModelImpl) Custom(dao types.ModelDAO, query types.Query, items ...interface{}) {
	currentDb := db.Pool().ByDAO(dao)
	currentDb.Custom(query, items...)
}

type IntModelImpl struct {
	BaseModelImpl
}

func (i IntModelImpl) ById(id int) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i IntModelImpl) ByIds(ids []int) types.ModelScanner {
	return criteria.Select().And(expr.In("id", ids))
}
