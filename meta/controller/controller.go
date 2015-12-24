package controller

import (
"github.com/byorty/hardcore/meta/common"
"github.com/byorty/hardcore/meta/types"
)

type Controller struct {
    common.Entity
	Route string `xml:"route,attr"`
	Actions []Action `xml:"actions>action"`
    Imports []string
    Filename string
    AutoFilename string
}

func (c Controller) GetKind() types.EntityKind {
    return types.ControllerEntityKind
}