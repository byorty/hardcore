package model

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
	"strings"
)

const DefaultIdentifierKind = "int64"
const DefaultIdentifierName = "id"
const TimeKind = "time.Time"
const TimeImport = "time"
const DefaultPropertyKind = "interface{}"

type Model struct {
	common.Entity
	Source     string        `xml:"source,attr"`
	Identifier *Property     `xml:"properties>identifier"`
	Properties []*Property   `xml:"properties>property"`
	Pattern    types.Pattern `xml:"pattern"`
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

func (m Model) GetPattern() types.Pattern {
	if len(m.Pattern) == 0 {
		m.Pattern = types.ValueObjectPattern
	}
	return m.Pattern
}

func (m Model) GetTable() string {
	return strings.ToLower(m.GetName()[0:1] + AZ09Regex.ReplaceAllString(m.GetName()[1:], "_$1"))
}

type Models []*Model

func (m Models) Len() int {
	return len(m)
}

func (m Models) Get(i int) types.Entity {
	return m[i]
}
