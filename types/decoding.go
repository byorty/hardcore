package types

type ImportablePropertyKind int

const (
	ScalarImportablePropertyKind ImportablePropertyKind = iota
	ModelImportablePropertyKind
	EnumImportablePropertyKind
	SliceImportablePropertyKind
)

type ImportableProperty interface {
	SetValue(interface{}, interface{})
	GetKind() ImportablePropertyKind
	GetProtoKind() ProtoKind
	GetImporter() Importer
}

type ImportableProperties map[string]ImportableProperty

type Importer interface {
	Get(string) (ImportableProperty, bool)
	SetProperties(ImportableProperties)
	SetImportable(interface{})
	GetImportable() interface{}
}

type Decoder interface {
	Decode([]byte, Importer)
}
