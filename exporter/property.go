package exporter

import "github.com/byorty/hardcore/types"

type PropertyImpl struct {
	name string
	prototyped types.Prototyped
}

func NewProperty(name string) PropertyImpl {
	return PropertyImpl{name: name}
}

func (p PropertyImpl) GetName() string {
	return p.name
}

func (p*PropertyImpl) GetPrototyped() types.Prototyped {
	return p.prototyped
}

func (p *PropertyImpl) SetPrototyped(prototyped types.Prototyped) {
	p.prototyped = prototyped
}