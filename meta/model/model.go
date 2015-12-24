package model

import (
    "github.com/byorty/hardcore/meta/common"
    "github.com/byorty/hardcore/meta/types"
)

type Model struct {
    common.Entity
    Source string `xml:"source,attr"`
    Identifier *Property `xml:"properties>identifier"`
    Properties []*Property `xml:"properties>property"`
    Pattern Pattern `xml:"pattern"`
    properties []types.Property
}

func (m Model) GetKind() types.EntityKind {
    return types.ModelEntityKind
}

func (m Model) GetSource() string {
    return m.Source
}

func (m *Model) SetProperties(properties []types.Property) {
    m.properties = properties
}

func (m Model) GetProperties() []types.Property {
    return m.properties
}
