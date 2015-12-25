package controller

import (
"github.com/byorty/hardcore/meta/common"
"github.com/byorty/hardcore/meta/types"
)

type Container struct {
	common.Container
	Route string `xml:"route,attr"`
	Controllers []*Controller `xml:"controller"`
}

func (c Container) GetContainerKind() types.ContainerKind {
    return types.ControllerContainerKind
}
