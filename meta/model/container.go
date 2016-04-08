package model

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Container struct {
	common.Container
	Enums  Enums  `xml:"enum"`
	Models Models `xml:"model"`
}

func (c Container) GetContainerKind() types.ContainerKind {
	return types.EntityContainerKind
}

func (c Container) Len() int {
	return len(c.GetEntities())
}

func (c Container) Get(i int) types.Entity {
	return c.GetEntities()[i]
}

func (c *Container) Init(env types.Environment) {
	c.Container.Init(env)

	c.SetEntities(make([]types.Entity, 0))
	c.merge(c.Models)
	c.merge(c.Enums)
}

func (c *Container) merge(slice types.EntitySlice) {
	entities := c.GetEntities()
	for i := 0; i < slice.Len(); i++ {
		entities = append(entities, slice.Get(i))
	}
	c.SetEntities(entities)
}

type Containers []*Container

func (c Containers) Get(i int) types.Container {
	return c[i]
}

func (c Containers) Len() int {
	return len(c)
}
