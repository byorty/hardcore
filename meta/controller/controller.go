package controller

import "github.com/byorty/hardcore/meta/common"

type Controller struct {
	Name string `xml:"name,attr"`
	Route string `xml:"route,attr"`
	Extends []*common.Extend `xml:"extends>extend"`
	Actions []Action `xml:"actions>action"`
    Imports []string
}

func (c Controller) HasParents() bool {
    return len(c.Extends) > 0
}