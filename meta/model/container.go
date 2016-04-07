package model

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Container struct {
	common.Container
	Enums    Enums  `xml:"enum"`
	Models   Models `xml:"model"`
	entities []types.Entity
}

func (c Container) GetContainerKind() types.ContainerKind {
	return types.EntityContainerKind
}

func (c Container) Len() int {
	return len(c.entities)
}

func (c Container) Get(i int) types.Entity {
	return c.entities[i]
}

func (c *Container) Init(env types.Environment) {
	c.Container.Init(env)
	c.entities = make([]types.Entity, 0)
	c.merge(c.Models)
	c.merge(c.Enums)
}

func (c *Container) merge(slice types.EntitySlice) {
	for i := 0; i < slice.Len(); i++ {
		c.entities = append(c.entities, slice.Get(i))
	}
}

type Containers []*Container

func (c *Containers) Add(container types.Container) {
	(*c) = append((*c), container.(*Container))
}

func (c Containers) Get(i int) types.Container {
	return c[i]
}

func (c Containers) Len() int {
	return len(c)
}
