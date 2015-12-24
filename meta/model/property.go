package model

import "github.com/byorty/hardcore/meta/types"

type Property struct {
    Name string `xml:"name,attr"`
    Kind string `xml:"type,attr"`
    Size int `xml:"size,attr"`
    Required bool `xml:"required,attr"`
    Relation types.Relation `xml:"relation,attr"`
    upperName string
    entity types.Entity
}

func (p Property) GetName() string {
    return p.Name
}

func (p Property) GetKind() string {
    return p.Kind
}

func (p Property) IsRequired() bool {
    return p.Required
}

func (p Property) GetRelation() types.Relation {
    return p.Relation
}

func (p *Property) SetUpperName(upperName string) {
    p.upperName = upperName
}

func (p Property) GetUpperName() string {
    return p.upperName
}

func (p *Property) SetEntity(entity types.Entity) {
    p.entity = entity
}

func (p Property) GetEntity() types.Entity {
    return p.entity
}

func (p Property) HasRelation() bool {
    return p.Relation != types.NoneRelation
}