package importer

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Container struct {
	common.Container
	Importers Importers `xml:"importer"`
}

func (c Container) GetContainerKind() types.ContainerKind {
	return types.ImporterContainerKind
}

func (c Container) Len() int {
	return len(c.Importers)
}

func (c Container) Get(i int) types.Entity {
	return c.Importers[i]
}

func (c *Container) Init(env types.Environment) {
	c.Container.Init(env)

	c.Merge(c.Importers)
}

type Containers []*Container

func (c Containers) Get(i int) types.Container {
	return c[i]
}

func (c Containers) Len() int {
	return len(c)
}
