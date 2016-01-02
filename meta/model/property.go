package model

import (
	"github.com/byorty/hardcore/meta/types"
	"regexp"
	"strings"
)

var (
	AZ09Regex = regexp.MustCompile(`([A-Z0-9])`)
)

type Property struct {
    Name string `xml:"name,attr"`
    Kind string `xml:"type,attr"`
    Size int `xml:"size,attr"`
    Required bool `xml:"required,attr"`
    Relation types.Relation `xml:"relation,attr"`
    upperName string
    entity types.Entity
	isSelfPackage bool
	relationProperty types.Property
	relationKind string
}

func (p Property) GetName() string {
    return p.Name
}

func (p Property) GetKind() string {
	if len(p.Kind) == 0 {
		p.Kind = DefaultPropertyKind
	}
    return p.Kind
}

func (p Property) GetField() string {
	return strings.ToLower(AZ09Regex.ReplaceAllString(p.GetName(), "_$1"))
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

func (p *Property) SetSelfPackage(isSelfPackage bool) {
	p.isSelfPackage = isSelfPackage
}

func (p Property) IsSelfPackage() bool {
	return p.isSelfPackage
}

func (p Property) GetDefineKind() string {
	switch p.GetRelation() {
	case types.OneToOneRelation:
		if p.IsSelfPackage() {
			return p.entity.GetPointerName()
		} else {
			return p.entity.GetPointerFullName()
		}
	case types.OneToManyRelation:
		if p.IsSelfPackage() {
			return p.entity.GetMultipleName()
		} else {
			return p.entity.GetFullMultipleName()
		}
	default: return p.GetKind()
	}
}

func (p Property) GetVariableKind() string {
	switch p.GetRelation() {
	case types.OneToOneRelation:
		if p.IsSelfPackage() {
			return p.entity.GetName()
		} else {
			return p.entity.GetFullName()
		}
	case types.OneToManyRelation:
		if p.IsSelfPackage() {
			return p.entity.GetMultipleName()
		} else {
			return p.entity.GetFullMultipleName()
		}
	default: return p.GetKind()
	}
}

func (p Property) GetProtoKind() string {
	if p.HasRelation() {
		if p.GetEntity().GetEntityKind().IsModel() {
			return "ProtoModelKind"
		} else {
			return "ProtoEnumKind"
		}
	} else {
		return "ProtoBasicKind"
	}
}

func (p *Property) SetRelationProperty(relationProperty types.Property) {
	p.relationProperty = relationProperty
}

func (p Property) GetRelationProperty() types.Property {
	return p.relationProperty
}
