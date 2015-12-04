package controller

type Container struct {
	Package string `xml:"package,attr"`
	Route string `xml:"route,attr"`
	Controllers []Controller `xml:"controller"`
}
