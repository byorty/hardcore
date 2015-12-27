package model

import (
    "github.com/byorty/hardcore/meta/common"
    "github.com/byorty/hardcore/meta/types"
)

const DefaultIdentifierKind = "int"
const DefaultIdentifierName = "id"

type Model struct {
    common.Entity
    Source string `xml:"source,attr"`
    Identifier *Property `xml:"properties>identifier"`
    Properties []*Property `xml:"properties>property"`
    Pattern types.Pattern `xml:"pattern"`
    properties []types.Property
}

func (m Model) GetEntityKind() types.EntityKind {
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

func (m Model) GetPattern() types.Pattern  {
	return m.Pattern
}
