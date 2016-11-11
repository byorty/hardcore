package types

type ImportableProperty interface {
	GetName() string
	GetAliasName() string
	GetSetterName() string
	GetKind() string
	HasModelProperty() bool
	SetModelProperty(Property)
	GetModelProperty() Property
	GetProtoKind() string
	GetMethod() string
}

type ImporterEntity interface {
	Entity
	GetSource() string
	SetSourceEntity(Entity)
	GetSourceEntity() Entity
	SetProperties([]ImportableProperty)
	GetProperties() []ImportableProperty
}
