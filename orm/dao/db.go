package dao

import (
	"github.com/byorty/hardcore/pool"
	"github.com/byorty/hardcore/query/criteria"
	"github.com/byorty/hardcore/query/expr"
	"github.com/byorty/hardcore/slice"
	"github.com/byorty/hardcore/types"
)

type BaseImpl struct {
	ByIdStmt   types.DBStatement
	ByIdsStmt  types.DBStatement
	InsertStmt types.DBStatement
	UpdateStmt types.DBStatement
}

func (b BaseImpl) Update(query types.Query, model types.Model) {
	currentDb := pool.DB().ByDAO(model.CommonDAO())
	currentDb.Exec(query).One(model)
}

func (b BaseImpl) Insert(query types.Query, model types.Model) {
	currentDb := pool.DB().ByDAO(model.CommonDAO())
	if currentDb.SupportLastInsertId() {
		currentDb.Exec(query).One(model)
	} else if currentDb.SupportReturningId() {
		var id int64
		currentDb.Custom(query).One(&id)
		setter := model.Proto().GetByName("id").GetSetter()
		setter(model, id)
	}
}

func (b BaseImpl) SelectAll(query types.Query, models types.Model) {
	dao := models.CommonDAO()
	currentDb := pool.DB().ByDAO(dao)
	currentDb.Query(query).All(models)
}

func (b BaseImpl) SelectOne(query types.Query, model types.Model) {
	dao := model.CommonDAO()
	currentDb := pool.DB().ByDAO(dao)
	currentDb.QueryRow(query).One(model)
}

func (b BaseImpl) Custom(dao types.ModelDAO, query types.Query, items ...interface{}) {
	currentDb := pool.DB().ByDAO(dao)
	currentDb.Custom(query).One(items...)
}

type IntImpl struct {
	BaseImpl
}

func (i IntImpl) ById(id int) types.DBRow {
	return i.ByIdStmt.QueryRow(id)
}

func (i IntImpl) ByIds(ids []int) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewIntBy(ids)))
}

type Int8Impl struct {
	BaseImpl
}

func (i Int8Impl) ById(id int8) types.DBRow {
	return i.ByIdStmt.QueryRow(id)
}

func (i Int8Impl) ByIds(ids []int8) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewInt8By(ids)))
}

type Int16Impl struct {
	BaseImpl
}

func (i Int16Impl) ById(id int16) types.DBRow {
	return i.ByIdStmt.QueryRow(id)
}

func (i Int16Impl) ByIds(ids []int16) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewInt16By(ids)))
}

type Int32Impl struct {
	BaseImpl
}

func (i Int32Impl) ById(id int32) types.DBRow {
	return i.ByIdStmt.QueryRow(id)
}

func (i Int32Impl) ByIds(ids []int32) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewInt32By(ids)))
}

type Int64Impl struct {
	BaseImpl
}

func (i Int64Impl) ById(id int64) types.DBRow {
	return i.ByIdStmt.QueryRow(id)
}

func (i Int64Impl) ByIds(ids []int64) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewInt64By(ids)))
}

type UintImpl struct {
	BaseImpl
}

func (u UintImpl) ById(id uint) types.DBRow {
	return u.ByIdStmt.QueryRow(id)
}

func (i UintImpl) ByIds(ids []uint) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUintBy(ids)))
}

type Uint8Impl struct {
	BaseImpl
}

func (u Uint8Impl) ById(id uint8) types.DBRow {
	return u.ByIdStmt.QueryRow(id)
}

func (i Uint8Impl) ByIds(ids []uint8) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUint8By(ids)))
}

type Uint16Impl struct {
	BaseImpl
}

func (u Uint16Impl) ById(id uint16) types.DBRow {
	return u.ByIdStmt.QueryRow(id)
}

func (i Uint16Impl) ByIds(ids []uint16) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUint16By(ids)))
}

type Uint32Impl struct {
	BaseImpl
}

func (u Uint32Impl) ById(id uint32) types.DBRow {
	return u.ByIdStmt.QueryRow(id)
}

func (i Uint32Impl) ByIds(ids []uint32) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUint32By(ids)))
}

type Uint64Impl struct {
	BaseImpl
}

func (u Uint64Impl) ById(id uint64) types.DBRow {
	return u.ByIdStmt.QueryRow(id)
}

func (i Uint64Impl) ByIds(ids []uint64) types.ModelScanner {
	return criteria.Select().And(expr.In("id", slice.NewUint64By(ids)))
}
