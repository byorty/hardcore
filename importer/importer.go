package importer

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	properties types.ImportableProperties
	importable interface{}
}

func (b BaseImpl) Get(key string) (types.ImportableProperty, bool) {
	property, ok := b.properties[key]
	return property, ok
}

func (b *BaseImpl) SetProperties(properties types.ImportableProperties) {
	b.properties = properties
}

func (b *BaseImpl) SetImportable(importable interface{}) {
	b.importable = importable
}

func (b BaseImpl) GetImportable() interface{} {
	return b.importable
}
