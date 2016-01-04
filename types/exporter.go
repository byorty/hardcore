package types

type ExportedProperty interface {
	GetName() string
	GetValue() interface{}
	GetProtoKind() ProtoKind
	SetProtoKind(ProtoKind)
	GetPrototyped() Prototyped
	SetPrototyped(Prototyped)
}

type Exporter interface {
	Slice
	Get(int) ExportedProperty
	Add(ExportedProperty) Exporter
	Export(Prototyped) Exporter
}
