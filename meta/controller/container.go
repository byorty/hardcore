package controller

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Container struct {
	common.Container
	Initializer
	Route       string      `xml:"path,attr"`
	Controllers Controllers `xml:"controller"`
}

func (c Container) GetContainerKind() types.ContainerKind {
	return types.ControllerContainerKind
}

func (c Container) GetRoute() string {
	if len(c.Route) == 0 {
		c.Route = "/"
	}
	return c.Route
}

func (c Container) Len() int {
	return len(c.GetEntities())
}

func (c Container) Get(i int) types.Entity {
	return c.GetEntities()[i]
}

func (c *Container) Init(env types.Environment) {
	c.Container.Init(env)

	c.Merge(c.Controllers)
	c.Merge(c.Befores)
	c.Merge(c.Afters)
}

type Containers []*Container

func (c Containers) Get(i int) types.Container {
	return c[i]
}

func (c Containers) Len() int {
	return len(c)
}
