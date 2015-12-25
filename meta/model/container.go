package model
import (
"github.com/byorty/hardcore/meta/common"
"github.com/byorty/hardcore/meta/types"
)

type Container struct {
    common.Container
    Enums []*Enum `xml:"enum"`
    Models []*Model `xml:"model"`
}

func (c Container) GetContainerKind() types.ContainerKind {
    return types.EntityContainerKind
}