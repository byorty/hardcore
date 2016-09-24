package types

type ImportablePropertyKind string

type ImportableProperty interface {
	GetName() string
	GetSetterName() string
	GetKind() ImportablePropertyKind
	HasModelProperty() bool
	SetModelProperty(Property)
	GetModelProperty() Property
}

type ImporterEntity interface {
	Entity
	GetSource() string
	SetSourceEntity(Entity)
	GetSourceEntity() Entity
	SetProperties([]ImportableProperty)
	GetProperties() []ImportableProperty
}
