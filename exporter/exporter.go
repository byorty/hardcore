package exporter

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	properties []types.ExportableProperty
	prototyped types.Prototyped
}

func (b *BaseImpl) SetProperties(properties []types.ExportableProperty) {
	b.properties = properties
}

func (b BaseImpl) Len() int {
	return len(b.properties)
}

func (b BaseImpl) Less(x, y int) bool {
	return b.properties[x].GetName() < b.properties[y].GetName()
}

func (b BaseImpl) Swap(x, y int) {
	b.properties[x], b.properties[y] = b.properties[y], b.properties[x]
}

func (b BaseImpl) GetRaw(x int) interface{} {
	return b.Get(x)
}

func (b BaseImpl) Get(x int) types.ExportableProperty {
	exportedProp := b.properties[x]
	exportedProp.SetPrototyped(b.prototyped)
	return exportedProp
}

func (b *BaseImpl) Add(property types.ExportableProperty) types.Exporter {
	b.properties = append(b.properties, property)
	return b
}

func (b *BaseImpl) Export(prototyped types.Prototyped) types.Exporter {
	b.prototyped = prototyped
	return b
}