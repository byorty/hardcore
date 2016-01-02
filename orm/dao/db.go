package dao

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/query/criteria"
	"github.com/byorty/hardcore/query/expr"
	"github.com/byorty/hardcore/slice"
)

type BaseImpl struct {}

func (b BaseImpl) Save(model types.Model) {
	idProperty := model.Proto().GetByName("id")
	getter := idProperty.GetGetter()
	criteria.Update().And(expr.Eq("id", getter.Call(model))).One(model)
}

func (b BaseImpl) Update(query types.Query, model types.Model) {
	currentDb := db.Pool().ByDAO(model.CommonDAO())
	currentDb.Exec(query, model.CommonDAO(), model)
}

func (b BaseImpl) Add(model types.Model) {
	criteria.Insert().One(model)
}

func (b BaseImpl) Insert(query types.Query, model types.Model) {
	currentDb := db.Pool().ByDAO(model.CommonDAO())
	if currentDb.SupportLastInsertId() {
		currentDb.Exec(query, model.CommonDAO(), model)
	} else if currentDb.SupportReturningId() {
		var id int64
		currentDb.Custom(query, &id)
		setter := model.Proto().GetByName("id").GetSetter()
		setter.Call(model, id)
	}
}

func (b BaseImpl) All(query types.Query, models types.Model) {
	dao := models.CommonDAO()
	currentDb := db.Pool().ByDAO(dao)
	currentDb.Query(query, dao, models)
}

func (b BaseImpl) One(query types.Query, model types.Model) {
	dao := model.CommonDAO()
	currentDb := db.Pool().ByDAO(dao)
	currentDb.QueryRow(query, dao, model)
}

func (b BaseImpl) Custom(dao types.ModelDAO, query types.Query, items ...interface{}) {
	currentDb := db.Pool().ByDAO(dao)
	currentDb.Custom(query, items...)
}

type IntImpl struct {
	BaseImpl
}

func (i IntImpl) ById(id int) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i IntImpl) ByIds(ids []int) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewIntsBy(ids)))
}

type Int8Impl struct {
	BaseImpl
}

func (i Int8Impl) ById(id int8) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i Int8Impl) ByIds(ids []int8) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewInt8sBy(ids)))
}

type Int16Impl struct {
	BaseImpl
}

func (i Int16Impl) ById(id int16) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i Int16Impl) ByIds(ids []int16) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewInt16sBy(ids)))
}

type Int32Impl struct {
	BaseImpl
}

func (i Int32Impl) ById(id int32) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i Int32Impl) ByIds(ids []int32) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewInt32sBy(ids)))
}

type Int64Impl struct {
	BaseImpl
}

func (i Int64Impl) ById(id int64) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i Int64Impl) ByIds(ids []int64) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewInt64sBy(ids)))
}

type UintImpl struct {
	BaseImpl
}

func (i UintImpl) ById(id uint) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i UintImpl) ByIds(ids []uint) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUintsBy(ids)))
}

type Uint8Impl struct {
	BaseImpl
}

func (i Uint8Impl) ById(id uint8) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i Uint8Impl) ByIds(ids []uint8) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUint8sBy(ids)))
}

type Uint16Impl struct {
	BaseImpl
}

func (i Uint16Impl) ById(id uint16) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i Uint16Impl) ByIds(ids []uint16) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUint16sBy(ids)))
}

type Uint32Impl struct {
	BaseImpl
}

func (i Uint32Impl) ById(id uint32) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i Uint32Impl) ByIds(ids []uint32) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUint32sBy(ids)))
}

type Uint64Impl struct {
	BaseImpl
}

func (i Uint64Impl) ById(id uint64) types.ModelScanner {
	return criteria.Select().And(expr.Eq("id", id))
}

func (i Uint64Impl) ByIds(ids []uint64) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUint64sBy(ids)))
}
