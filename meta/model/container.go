package model
import "github.com/byorty/hardcore/meta/common"

type Container struct {
    common.Container
    Enums []*Enum `xml:"enum"`
    Models []*Model `xml:"model"`
}

func (c Container) Eq(container *Container) bool {
    return c.Package == container.Package
}