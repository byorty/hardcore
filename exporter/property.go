package exporter

import "github.com/byorty/hardcore/types"

type PropertyImpl struct {
	name string
	kind types.ProtoKind
	prototyped types.Prototyped
}

func NewProperty(name string) PropertyImpl {
	return PropertyImpl{name: name}
}

func NewPropertyByKind(name string, kind types.ProtoKind) PropertyImpl {
	prop := NewProperty(name)
	prop.kind = kind
	return prop
}

func (p PropertyImpl) GetName() string {
	return p.name
}

func (p PropertyImpl) GetProtoKind() types.ProtoKind {
	return p.kind
}

func (p *PropertyImpl) SetProtoKind(kind types.ProtoKind) {
	p.kind = kind
}

func (p*PropertyImpl) GetPrototyped() types.Prototyped {
	return p.prototyped
}

func (p *PropertyImpl) SetPrototyped(prototyped types.Prototyped) {
	p.prototyped = prototyped
}