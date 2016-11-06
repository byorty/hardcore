package types

import "time"

type ImportablePropertyKind int

const (
	ScalarImportablePropertyKind ImportablePropertyKind = iota
	ModelImportablePropertyKind
	EnumImportablePropertyKind
	SliceImportablePropertyKind
)

type ImportableProperty interface {
	ProtoEntity
	//SetValue(interface{}, interface{})
	//GetKind() ImportablePropertyKind
	//GetImporter() Importer
}

//type ImportableProperties map[string]ImportableProperty

type Importer interface {
	Get(string) (ImportableProperty, bool)
	Decode(string, Decoder, []byte)
	//Get(string) (ImportableProperty, bool)
	//SetProperties(ImportableProperties)
	//SetImportable(interface{})
	//GetImportable() interface{}
}

type Decoder interface {
	DecodeInt([]byte) int
	DecodeInt8([]byte) int8
	DecodeInt16([]byte) int16
	DecodeInt32([]byte) int32
	DecodeInt64([]byte) int64
	DecodeUint([]byte) uint
	DecodeUint8([]byte) uint8
	DecodeUint16([]byte) uint16
	DecodeUint32([]byte) uint32
	DecodeUint64([]byte) uint64
	DecodeFloat32([]byte) float32
	DecodeFloat64([]byte) float64
	DecodeString([]byte) string
	DecodeBool([]byte) bool
	DecodeTime([]byte) time.Time
	DecodeModel(Importer, []byte)
	//DecodeSlice(SliceExporter)
	One([]byte, Importer)
	//All(SliceExporter) []byte
	//Decode([]byte, Importer)
}
