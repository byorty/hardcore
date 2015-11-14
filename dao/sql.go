package dao

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/db"
	"github.com/byorty/hardcore/query/criteria"
	"github.com/byorty/hardcore/expr"
	"github.com/byorty/hardcore/query"
)

type Sql struct {}

func (s Sql) Save(model types.StraightMappingModel) {

}

func (s Sql) getWriter(model types.StraightMappingModel) types.QueryWriter {
	dao := model.DAO()
	writer := db.Pool().ByDAO(dao).GetQueryWriter()
	writer.SetProto(model.Proto())
	writer.SetTable(dao.GetTable())
	return writer
}

func (s Sql) Insert(model types.StraightMappingModel, args ...interface{}) {
	writer := s.getWriter(model)
	writer.SetArgs(args)

	currentDb := db.Pool().ByDAO(model.DAO())
	currentDb.Exec(query.Sql{writer.WriteInsert(), args}, model.DAO(), model)
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
