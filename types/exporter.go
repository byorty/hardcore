package types

type ExportableProperty interface {
	GetName() string
	GetValue() interface{}
	GetPrototyped() Prototyped
	SetPrototyped(Prototyped)
}

type Exporter interface {
	Slice
	Get(int) ExportableProperty
	Add(ExportableProperty) Exporter
	Export(Prototyped) Exporter
}
