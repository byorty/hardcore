package controller

type Action struct {
	Name string `xml:"name,attr"`
	Method string `xml:"method,attr"`
	Params []*Param `xml:"params>param"`
}
