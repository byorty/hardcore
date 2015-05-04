package criteria

import (
	"fmt"
	"github.com/byorty/hardcore/types"
)

type criteria struct {
	dao types.DAO
}

func New(dao types.DAO) types.Criteria {
	return &criteria{
		dao: dao,
	}
}

func (c *criteria) Add(conditions ...types.Logic) types.Criteria {
	for _, condition := range conditions {
		fmt.Println(condition)
	}
	return c
}

func (c *criteria) One(model interface{}) {

}

func (c *criteria) All(model interface{}) {

}
