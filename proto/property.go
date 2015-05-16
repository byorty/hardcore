package proto

import "github.com/byorty/hardcore/types"

type Property struct {
	field    string
	kind     types.ProtoKind
	relation types.ProtoRelation
	required bool
	proto    types.Proto
}

func NewProperty(field string, kind types.ProtoKind, relation types.ProtoRelation, required bool) types.ProtoProperty {
	return newProperty(field, kind, relation, required)
}

func newProperty(field string, kind types.ProtoKind, relation types.ProtoRelation, required bool) *Property {
	return &Property{
		field   : field,
		kind    : kind,
		relation: relation,
		required: required,
	}
}

func NewRelationProperty(field string, kind types.ProtoKind, relation types.ProtoRelation, required bool, proto types.Proto) types.ProtoProperty {
	prop := newProperty(field, kind, relation, required)
	prop.proto = proto
	return prop
}

func (p Property) GetField() string {
	return p.field
}

func (p Property) GetKind() types.ProtoKind {
	return p.kind
}

func (p Property) GetRelation() types.ProtoRelation {
	return p.relation
}

func (p Property) GetProto() types.Proto {
	return p.proto
}

func (p Property) IsRequired() bool {
	return p.required
}
