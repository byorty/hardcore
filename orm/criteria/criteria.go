package criteria

import (
	"github.com/byorty/hardcore/types"
	"fmt"
	"strings"
	"github.com/byorty/hardcore/orm/db"
)

type CriteriaImpl struct {
	dao    types.DAO
	proto  types.Proto
	conditions []types.Logic
	args []interface{}
}

func New(dao types.DAO) types.Criteria {
	return &CriteriaImpl{
		dao: dao,
		conditions: make([]types.Logic, 0),
	}
}

func (c *CriteriaImpl) Add(condition types.Logic) types.Criteria {
	c.conditions = append(c.conditions, condition)
	return c
}

func (c *CriteriaImpl) One(model types.StraightMappingModel) {
	c.proto = model.Proto()
	currentDb := db.Pool().ByDAO(c.dao)
	currentDb.QueryRow(c, c.dao, model)
}

func (c *CriteriaImpl) All(model interface{}) {

}

func (c *CriteriaImpl) ToNative() interface{} {
	fields := make([]string, len(c.proto.GetSlice()))
	for i, field := range c.proto.GetSlice() {
		fields[i] = fmt.Sprintf("\"%s\".\"%s\"", c.dao.GetTable(), field.GetField())
	}
	query := fmt.Sprintf("SELECT %s FROM \"%s\"", strings.Join(fields, ", "), c.dao.GetTable())

	conditionsLen := len(c.conditions)
	if conditionsLen > 0 {
		conditions := make([]string, conditionsLen)
		c.args = make([]interface{}, conditionsLen)
		for i, condition := range c.conditions {
			field := c.proto.GetByName(condition.GetLeft())
			conditions[i] = fmt.Sprintf(
				"\"%s\".\"%s\" %s $%d",
				c.dao.GetTable(),
				field.GetField(),
				condition.GetLogic(),
				i + 1,
			)
			c.args[i] = condition.GetRight()
		}
		query += fmt.Sprintf(" WHERE %s", strings.Join(conditions, " "))
	}
	return query
}

func (c *CriteriaImpl) GetArgs() []interface{} {
	return c.args
}
