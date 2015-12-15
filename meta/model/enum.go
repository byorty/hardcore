package model

type Enum struct {
    Name string `xml:"name,attr"`
    Kind string `xml:"type,attr"`
    Constants []*Constant `xml:"constants>constant"`
    Filename string
    AutoFilename string
}