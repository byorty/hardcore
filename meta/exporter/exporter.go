package exporter

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
)

type Exporter struct {
	Source     string      `xml:"source,attr"`
	Properties []*Property `xml:"properties>property"`
	common.Entity
	entity types.Entity
	props  []types.ExportableProperty
}

func (e Exporter) GetEntityKind() types.EntityKind {
	return types.ExporterEntityKind
}

func (e Exporter) GetSource() string {
	return e.Source
}

func (e *Exporter) SetSourceEntity(entity types.Entity) {
	e.entity = entity
}

func (e Exporter) GetSourceEntity() types.Entity {
	return e.entity
}

func (e *Exporter) SetProperties(props []types.ExportableProperty) {
	e.props = props
}

func (e Exporter) GetProperties() []types.ExportableProperty {
	return e.props
}

type Exporters []*Exporter

func (e Exporters) Len() int {
	return len(e)
}

func (e Exporters) Get(i int) types.Entity {
	return e[i]
}
