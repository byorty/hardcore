package exporter

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Container struct {
	common.Container
	Exporters Exporters `xml:"exporter"`
}

func (c Container) GetContainerKind() types.ContainerKind {
	return types.ExporterContainerKind
}

func (c Container) Len() int {
	return len(c.Exporters)
}

func (c Container) Get(i int) types.Entity {
	return c.Exporters[i]
}

func (c *Container) Init(env types.Environment) {
	c.Container.Init(env)

	c.Merge(c.Exporters)
}

type Containers []*Container

func (c Containers) Get(i int) types.Container {
	return c[i]
}

func (c Containers) Len() int {
	return len(c)
}
