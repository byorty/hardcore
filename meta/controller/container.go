package controller

import "github.com/byorty/hardcore/meta/common"

type Container struct {
	common.Container
	Route string `xml:"route,attr"`
	Controllers []*Controller `xml:"controller"`
}

func (c Container) Eq(container *Container) bool {
	return c.Package == container.Package
}
