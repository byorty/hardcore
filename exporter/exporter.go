package exporter

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	properties []types.ExportedProperty
	prototyped types.Prototyped
}

func (b *BaseImpl) SetProperties(properties []types.ExportedProperty) {
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

func (b BaseImpl) Get(x int) types.ExportedProperty {
	exportedProp := b.properties[x]
	exportedProp.SetPrototyped(b.prototyped)
//	exportedProp.SetProtoKind(b.prototyped.Proto().GetByName(exportedProp.GetName()).GetKind())
	return exportedProp
}

func (b *BaseImpl) Add(property types.ExportedProperty) types.Exporter {
	b.properties = append(b.properties, property)
	return b
}

func (b *BaseImpl) Export(prototyped types.Prototyped) types.Exporter {
	b.prototyped = prototyped
	return b
}