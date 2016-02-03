package controller

import (
"github.com/byorty/hardcore/meta/common"
"github.com/byorty/hardcore/meta/types"
)

type Container struct {
	common.Container
	Route string `xml:"path,attr"`
	Controllers []*Controller `xml:"controller"`
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
