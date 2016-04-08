package model

import (
	"fmt"
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
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

func (m *Model) Init(container types.Container) {
	m.Entity.Init(container)

	isStraightMapping := m.GetPattern() == types.StraightMappingPattern
	hasntProperties := m.GetProperties() == nil

	if hasntProperties {
		m.SetProperties(make([]types.Property, 0))
	}

	if isStraightMapping {
		var kind string
		if m.Identifier == nil || (m.Identifier != nil && len(m.Identifier.Kind) == 0) {
			kind = DefaultIdentifierKind
		} else {
			kind = m.Identifier.Kind
		}
		prop := &Property{
			Name:     DefaultIdentifierName,
			Kind:     kind,
			Size:     0,
			Required: true,
			Relation: types.NoneRelation,
		}
		prop.SetUpperName(utils.UpperFirst(prop.GetName()))

		m.SetProperties(append(m.GetProperties(), prop))
	}
}

func (m *Model) PostInit(env types.Environment) {
	config := env.GetConfiguration()
	properties := m.GetProperties()

	for _, property := range m.Properties {
		property.SetUpperName(utils.UpperFirst(property.GetName()))

		if property.HasRelation() {
			relEntity := config.GetEntity(property.GetKind())
			if relEntity == nil {
				env.GetLogger().Error(
					"invalid type %s in property %s entity %s",
					property.GetKind(),
					property.GetName(),
					m.GetName(),
				)
			} else {
				property.SetSelfPackage(m.GetContainer().GetImport() == relEntity.GetContainer().GetImport())
				if !property.IsSelfPackage() {
					m.AddImport(relEntity.GetContainer().GetImport())
				}
				property.SetEntity(relEntity)
				properties = append(properties, property)

				if m.GetPattern() == types.StraightMappingPattern {
					relation := property.GetRelation()
					if relation.IsOneToOne() {
						var kind string
						if relEntity.GetEntityKind() == types.EnumEntityKind {
							enumEntity := relEntity.(*Enum)
							kind = string(enumEntity.GetKind())
						} else if relEntity.GetEntityKind() == types.ModelEntityKind {
							relModelEntity := relEntity.(types.ModelEntity)
							identifier := relModelEntity.GetProperties()[0]
							kind = identifier.GetKind()
						}

						prop := &Property{
							Name:     fmt.Sprintf("%sId", property.GetName()),
							Kind:     kind,
							Size:     0,
							Required: property.IsRequired(),
							Relation: types.NoneRelation,
						}
						prop.SetUpperName(utils.UpperFirst(prop.GetName()))
						properties = append(properties, prop)
					}
				}
			}
		} else {
			if property.GetKind() == TimeKind {
				m.AddImport(TimeImport)
			}
			properties = append(properties, property)
		}
	}

	m.SetProperties(properties)
}

type Models []*Model

func (m Models) Len() int {
	return len(m)
}

func (m Models) Get(i int) types.Entity {
	return m[i]
}
