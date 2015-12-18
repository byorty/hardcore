package model

import "github.com/byorty/hardcore/meta/common"

type Model struct {
    Name string `xml:"name,attr"`
    Extends []*common.Extend `xml:"extends>extend"`
    Filename string
    AutoFilename string
}
