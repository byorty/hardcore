package exporter

import "github.com/byorty/hardcore/types"

//type ExportedProperty interface {
//	GetName() string
//	GetProtoKind() ProtoKind
//	GetValue() interface{}
//	SetPrototyped(Prototyped) ExportedProperty
//}

type Property struct {
	name string
	kind types.ProtoKind
	prototyped types.Prototyped
}

func (p Property) GetName() string {
	return p.name
}

//func (p Property) GetProtoKind() types.ProtoKind {
//
//}

func (p *Property) SetPrototyped(prototyped types.Prototyped) types.ExportedProperty {
	p.prototyped = prototyped
	return p
}