package dao

import (
	"github.com/byorty/hardcore/query/criteria"
	"github.com/byorty/hardcore/query/expr"
	"github.com/byorty/hardcore/slice"
	"github.com/byorty/hardcore/types"
)

type OneToMany struct {
	BaseImpl
	field string
}

type IntOneToMany struct {
	OneToMany
}

func NewIntOneToMany(field string) *IntOneToMany {
	var dao IntOneToMany
	dao.field = field
	return &dao
}

func (o IntOneToMany) ById(id int) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o IntOneToMany) ByIds(ids []int) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewIntsBy(ids)))
}

type Int8OneToMany struct {
	OneToMany
}

func NewInt8OneToMany(field string) *Int8OneToMany {
	var dao Int8OneToMany
	dao.field = field
	return &dao
}

func (o Int8OneToMany) ById(id int8) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o Int8OneToMany) ByIds(ids []int8) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewInt8sBy(ids)))
}

type Int16OneToMany struct {
	OneToMany
}

func NewInt16OneToMany(field string) *Int16OneToMany {
	var dao Int16OneToMany
	dao.field = field
	return &dao
}

func (o Int16OneToMany) ById(id int16) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o Int16OneToMany) ByIds(ids []int16) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewInt16sBy(ids)))
}

type Int32OneToMany struct {
	OneToMany
}

func NewInt32OneToMany(field string) *Int32OneToMany {
	var dao Int32OneToMany
	dao.field = field
	return &dao
}

func (o Int32OneToMany) ById(id int32) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o Int32OneToMany) ByIds(ids []int32) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewInt32sBy(ids)))
}

type Int64OneToMany struct {
	OneToMany
}

func NewInt64OneToMany(field string) *Int64OneToMany {
	var dao Int64OneToMany
	dao.field = field
	return &dao
}

func (o Int64OneToMany) ById(id int64) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o Int64OneToMany) ByIds(ids []int64) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewInt64sBy(ids)))
}

type UintOneToMany struct {
	OneToMany
}

func NewUintOneToMany(field string) *UintOneToMany {
	var dao UintOneToMany
	dao.field = field
	return &dao
}

func (o UintOneToMany) ById(id uint) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o UintOneToMany) ByIds(ids []uint) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewUintsBy(ids)))
}

type Uint8OneToMany struct {
	OneToMany
}

func NewUint8OneToMany(field string) *Uint8OneToMany {
	var dao Uint8OneToMany
	dao.field = field
	return &dao
}

func (o Uint8OneToMany) ById(id uint8) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o Uint8OneToMany) ByIds(ids []uint8) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewUint8sBy(ids)))
}

type Uint16OneToMany struct {
	OneToMany
}

func NewUint16OneToMany(field string) *Uint16OneToMany {
	var dao Uint16OneToMany
	dao.field = field
	return &dao
}

func (o Uint16OneToMany) ById(id uint16) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o Uint16OneToMany) ByIds(ids []uint16) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewUint16sBy(ids)))
}

type Uint32OneToMany struct {
	OneToMany
}

func NewUint32OneToMany(field string) *Uint32OneToMany {
	var dao Uint32OneToMany
	dao.field = field
	return &dao
}

func (o Uint32OneToMany) ById(id uint32) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o Uint32OneToMany) ByIds(ids []uint32) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewUint32sBy(ids)))
}

type Uint64OneToMany struct {
	OneToMany
}

func NewUint64OneToMany(field string) *Uint64OneToMany {
	var dao Uint64OneToMany
	dao.field = field
	return &dao
}

func (o Uint64OneToMany) ById(id uint64) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o Uint64OneToMany) ByIds(ids []uint64) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, slice.NewUint64sBy(ids)))
}
