package exporter

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	properties []types.ExportableProperty
	exportable interface{}
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
	return exportedProp
}

func (b *BaseImpl) Add(property types.ExportableProperty) types.Exporter {
	b.properties = append(b.properties, property)
	return b
}

func (b *BaseImpl) SetExportable(exportable interface{}) types.Exporter {
	b.exportable = exportable
	return b
}

func (b *BaseImpl) GetExportable() interface{} {
	return b.exportable
}