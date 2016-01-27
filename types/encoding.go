package types

type ExportableProperty interface {
	GetName() string
	GetValue(Model) interface{}
}

type Exporter interface {
	Slice
	Get(int) ExportableProperty
	Add(ExportableProperty) Exporter
	SetExportable(Model) Exporter
	GetExportable() Model
}

type Encoder interface {
	Encode(Exporter) []byte
}
