package controller

type Container struct {
//	types.ControllerContainer
	Package string `xml:"package,attr"`
	Route string `xml:"route,attr"`
	Controllers []Controller `xml:"controller"`
}
