package types

type ExportedProperty interface {
	GetName() string
	GetProtoKind() ProtoKind
	GetValue() interface{}
	SetPrototyped(Prototyped) ExportedProperty
}

type Exporter interface {
	Slice
	Get(int) ExportedProperty
	Add(ExportedProperty) Exporter
	Export(Prototyped) Exporter
}
