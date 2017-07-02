package exporter

import "github.com/byorty/hardcore/types"

type PropertyImpl struct {
	name    string
	kind    types.ProtoKind
}

func NewProperty(name string) PropertyImpl {
	return PropertyImpl{name: name}
}

func NewPropertyWithKind(name string, kind types.ProtoKind) PropertyImpl {
	return PropertyImpl{name: name, kind: kind}
}

func (p PropertyImpl) GetName() string {
	return p.name
}

func (p PropertyImpl) GetProtoKind() types.ProtoKind {
	return p.kind
}
