package query

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/db"
)

type CriteriaImpl struct {
	dao    types.DAO
	proto  types.Proto
	conditions []types.Logic
	args []interface{}
}

func Criteria() types.Criteria {
	return &CriteriaImpl{
		conditions: make([]types.Logic, 0),
		args: make([]interface{}, 0),
	}
}

func (c *CriteriaImpl) Add(condition types.Logic) types.Criteria {
	c.conditions = append(c.conditions, condition)
	c.args = append(c.args, condition.GetRight())
	return c
}

func (c *CriteriaImpl) One(model types.StraightMappingModel) {
	c.dao = model.CommonDAO()
	c.proto = model.Proto()
	currentDb := db.Pool().ByDAO(c.dao)
	currentDb.QueryRow(c, c.dao, model)
}

func (c *CriteriaImpl) All(model types.StraightMappingModel) {
	c.dao = model.CommonDAO()
	c.proto = model.Proto()
	currentDb := db.Pool().ByDAO(c.dao)
	currentDb.Query(c, c.dao, model)
}

func (c *CriteriaImpl) ToNative() interface{} {
	writer := db.Pool().ByDAO(c.dao).GetQueryWriter()
	writer.SetProto(c.proto)
	writer.SetTable(c.dao.GetTable())
	writer.SetConditions(c.conditions)
	return writer.WriteSelect()
}

func (c *CriteriaImpl) GetArgs() []interface{} {
	return c.args
}
