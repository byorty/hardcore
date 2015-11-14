package proto

import "github.com/byorty/hardcore/types"

type Property struct {
	field    string
	kind     types.ProtoKind
	relation types.ProtoRelation
	required bool
	proto    types.Proto
	setter   types.Setter
	getter   types.Getter
}

func NewProperty(field string, kind types.ProtoKind, relation types.ProtoRelation, required bool, setter types.Setter, getter types.Getter) types.ProtoProperty {
	return newProperty(field, kind, relation, required, setter, getter)
}

func newProperty(field string, kind types.ProtoKind, relation types.ProtoRelation, required bool, setter types.Setter, getter types.Getter) *Property {
	return &Property{
		field   : field,
		kind    : kind,
		relation: relation,
		required: required,
		setter: setter,
		getter: getter,
	}
}

func NewRelationProperty(field string, kind types.ProtoKind, relation types.ProtoRelation, required bool, proto types.Proto, setter types.Setter, getter types.Getter) types.ProtoProperty {
	prop := newProperty(field, kind, relation, required, setter, getter)
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

func (p Property) GetSetter() types.Setter {
	return p.setter
}

func (p Property) GetGetter() types.Getter {
	return p.getter
}
