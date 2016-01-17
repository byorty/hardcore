package controller

import (
"github.com/byorty/hardcore/meta/common"
"github.com/byorty/hardcore/meta/types"
)

type Controller struct {
    common.Entity
	Route string `xml:"route,attr"`
	Actions []*Action `xml:"actions>action"`
	actions []types.Action
}

func (c Controller) GetEntityKind() types.EntityKind {
    return types.ControllerEntityKind
}

func (c Controller) GetRoute() string {
	return c.Route
}

func (c Controller) GetActions() []types.Action {
	return c.actions
}

func (c *Controller) SetActions(actions []types.Action) {
	c.actions = actions
}