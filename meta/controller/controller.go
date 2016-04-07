package controller

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Controller struct {
	*common.Entity
	Route   string    `xml:"path,attr"`
	Actions []*Action `xml:"actions>action"`
	actions []types.Action
}

func (c Controller) GetEntityKind() types.EntityKind {
	return types.ControllerEntityKind
}

func (c Controller) GetRoute() string {
	if len(c.Route) == 0 {
		c.Route = "/"
	}
	return c.Route
}

func (c Controller) GetActions() []types.Action {
	return c.actions
}

func (c *Controller) SetActions(actions []types.Action) {
	c.actions = actions
}

func (c *Controller) Init(container types.Container) {
	c.Entity.Init(container)

	for _, action := range c.Actions {
		if len(action.Params) == 0 {
			action.Params = []*Param{
				&Param{
					Name:     "scope",
					Required: true,
					Source:   "",
					Kind:     RequestScopeKind,
				},
			}
		}
	}
}

type Controllers []*Controller

func (c Controllers) Len() int {
	return len(c)
}

func (c Controllers) Get(i int) types.Entity {
	return c[i]
}
