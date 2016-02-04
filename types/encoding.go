package types

type ExportableProperty interface {
	GetName() string
	GetValue(interface{}) interface{}
}

type Exporter interface {
	Slice
	Get(int) ExportableProperty
	Add(ExportableProperty) Exporter
	SetExportable(interface{}) Exporter
	GetExportable() interface{}
}

type Encoder interface {
	Encode(Exporter) []byte
}
