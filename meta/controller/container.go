package controller

type Container struct {
	Package string `xml:"package,attr"`
	Route string `xml:"route,attr"`
	Controllers []Controller `xml:"controller"`
}

func (c Container) Eq(container Container) bool {
	return c.Package == container.Package
}
