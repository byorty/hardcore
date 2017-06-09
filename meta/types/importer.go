package types

type ImportableProperty interface {
	GetName() string
	GetAliasName() string
	GetGetterName() string
	GetSetterName() string
	GetKind() string
	HasModelProperty() bool
	SetModelProperty(Property)
	GetModelProperty() Property
	GetProtoKind() string
	GetMethod() string
	IsSlice() bool
	IsModelSlice() bool
	IsNotModelSlice() bool
	GetSliceConstruct() string
}

type ImporterEntity interface {
	Entity
	GetSource() string
	SetSourceEntity(Entity)
	GetSourceEntity() Entity
	SetProperties([]ImportableProperty)
	GetProperties() []ImportableProperty
}
