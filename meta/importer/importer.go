package importer

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Importer struct {
	Source     string      `xml:"source,attr"`
	Properties []*Property `xml:"properties>property"`
	common.Entity
	entity types.Entity
	props  []types.ImportableProperty
}

func (i Importer) GetEntityKind() types.EntityKind {
	return types.ImporterEntityKind
}

func (i Importer) GetSource() string {
	return i.Source
}

func (i *Importer) SetSourceEntity(entity types.Entity) {
	i.entity = entity
}

func (i Importer) GetSourceEntity() types.Entity {
	return i.entity
}

func (i *Importer) SetProperties(props []types.ImportableProperty) {
	i.props = props
}

func (i Importer) GetProperties() []types.ImportableProperty {
	return i.props
}

func (i *Importer) Init(container types.Container) {
	i.Entity.Init(container)

	props := make([]types.ImportableProperty, 0)
	for _, prop := range i.Properties {
		if len(prop.Kind) == 0 {
			prop.Kind = "string"
		}
		props = append(props, prop)
	}
	i.SetProperties(props)
	i.Properties = nil
}

type Importers []*Importer

func (i Importers) Len() int {
	return len(i)
}

func (i Importers) Get(j int) types.Entity {
	return i[j]
}
