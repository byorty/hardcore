package model

import "github.com/byorty/hardcore/meta/common"

type Model struct {
    Name string `xml:"name,attr"`
    Source string `xml:"source,attr"`
    Extends []*common.Extend `xml:"extends>extend"`
    Filename string
    AutoFilename string
    Identifier *Property `xml:"properties>identifier"`
    Properties []*Property `xml:"properties>property"`
    Pattern Pattern `xml:"pattern"`
}
