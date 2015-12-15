package model

type Constant struct {
    Name string `xml:"name,attr"`
    Value string `xml:"value,attr"`
    Label string `xml:",chardata"`
}