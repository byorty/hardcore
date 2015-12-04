package controller

import "github.com/byorty/hardcore/meta/common"

type Param struct {
	Name string `xml:"name,attr"`
	Required bool `xml:"required"`
	From From `xml:"from,attr"`
	Kind common.Kind `xml:"type,attr"`
}
