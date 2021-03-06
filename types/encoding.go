package types

import "time"

type ExportableProperty interface {
	ProtoEntity
	GetName() string
}

type ScalarExporter interface {
	Export(int, Encoder)
	Len() int
}

type Exporter interface {
	ProtoEntity
	ScalarExporter
	Get(int) ExportableProperty
}

type Encoder interface {
	EncodeInt(int)
	EncodeInt8(int8)
	EncodeInt16(int16)
	EncodeInt32(int32)
	EncodeInt64(int64)
	EncodeUint(uint)
	EncodeUint8(uint8)
	EncodeUint16(uint16)
	EncodeUint32(uint32)
	EncodeUint64(uint64)
	EncodeFloat32(float32)
	EncodeFloat64(float64)
	EncodeString(string)
	EncodeBool(bool)
	EncodeTime(time.Time)
	EncodeBytes([]byte)
	EncodeModel(Exporter)
	Encode(Exporter) []byte
}
