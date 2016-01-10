package controller

//import "github.com/byorty/hardcore/meta/types"

type Action struct {
	Name string `xml:"name,attr"`
	Route string `xml:"route,attr"`
	Method string `xml:"method,attr"`
	Params []*Param `xml:"params>param"`
}

func (a Action) GetName() string {
	return a.Name
}

func (a Action) GetRoute() string {
	return a.Route
}

func (a Action) GetMethod() string {
	return a.Method
}

//func (a Action) GetParameters() []types.ActionParameter {
//	return a.Params
//}
