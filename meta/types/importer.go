package types

type ImportableProperty interface {
	GetName() string
	GetSetterName() string
	GetKind() string
	HasModelProperty() bool
	SetModelProperty(Property)
	GetModelProperty() Property
	GetProtoKind() string
}

type ImporterEntity interface {
	Entity
	GetSource() string
	SetSourceEntity(Entity)
	GetSourceEntity() Entity
	SetProperties([]ImportableProperty)
	GetProperties() []ImportableProperty
}
