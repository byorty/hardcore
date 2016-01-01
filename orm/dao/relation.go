package dao

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/query/criteria"
	"github.com/byorty/hardcore/query/expr"
)

type OneToMany struct {
	BaseModelImpl
	field string
}

type IntOneToMany struct {
	OneToMany
}

func NewIntOneToMany(field string) *IntOneToMany  {
	var dao IntOneToMany
	dao.field = field
	return &dao
}

func (o IntOneToMany) ById(id int) types.ModelScanner {
	return criteria.Select().And(expr.Eq(o.field, id))
}

func (o IntOneToMany) ByIds(ids []int) types.ModelScanner {
	return criteria.Select().And(expr.In(o.field, ids))
}


