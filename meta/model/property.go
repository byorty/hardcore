package model

type Property struct {
    Name string `xml:"name,attr"`
    Kind string `xml:"type,attr"`
    Size int `xml:"size,attr"`
    Required bool `xml:"required,attr"`
    Relation Relation `xml:"relation,attr"`
    RelationName string
    RelationKind string
}

func (p Property) HasRelation() bool {
    return p.Relation != NoneRelation
}