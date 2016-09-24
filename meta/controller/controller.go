package controller

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Controller struct {
	common.Entity
	Initializer
	Route   string         `xml:"path,attr"`
	Scheme  Scheme         `xml:"scheme,attr"`
	Kind    types.ControllerKind `xml:"type,attr"`
	Actions []*Action      `xml:"actions>action"`
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

	actions := make([]types.Action, 0)
	for _, action := range c.Actions {
		params := make([]types.ActionParam, 0)
		for _, param := range action.Params {
			params = append(params, param)
		}
		if len(params) == 0 {
			var param *Param
			if c.Kind.IsWebsocket() {
				param = &Param{
					Name:     "ws",
					Required: true,
					Source:   "",
					Kind:     WebsocketScopeKind,
				}
			} else {
				param = &Param{
					Name:     "rs",
					Required: true,
					Source:   "",
					Kind:     RequestScopeKind,
				}
			}
			params = append(params, param)
		}
		action.setController(c)
		action.SetParams(params)
		actions = append(actions, action)
		action.Params = nil
	}
	c.SetActions(actions)
	c.Actions = nil
}

func (c Controller) GetScheme() Scheme {
	if len(c.Scheme) == 0 {
		c.Scheme = HttpScheme
	}
	return c.Scheme
}

func (c Controller) GetKind() types.ControllerKind {
	if len(c.Kind) == 0 {
		c.Kind = types.HttpControllerKind
	}
	return c.Kind
}

type Controllers []*Controller

func (c Controllers) Len() int {
	return len(c)
}

func (c Controllers) Get(i int) types.Entity {
	return c[i]
}
