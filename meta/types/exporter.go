package types

type ExportableProperty interface {
	HasGetter() bool
	SetHasGetter(bool)
	GetName() string
	GetAliasName() string
	GetGetterName() string
}

type ExporterEntity interface {
	Entity
	GetSource() string
	SetSourceEntity(Entity)
	GetSourceEntity() Entity
	SetProperties([]ExportableProperty)
	GetProperties() []ExportableProperty
}
