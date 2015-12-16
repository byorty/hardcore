package model

type Constant struct {
    Name string `xml:"name,attr"`
    Value string `xml:"value,attr"`
    Label string `xml:",chardata"`
}

func (c Constant) IsWriteKind(i int) bool {
    return i == 0
}

func (c Constant) HasValue() bool {
    return len(c.Value) > 0
}