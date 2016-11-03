package types

import "time"

type ExportableProperty interface {
	HasName() bool
	GetName() string
	GetProtoKind() ProtoKind
}

type Exporter interface {
	Len() int
	Get(int) ExportableProperty
	GetProtoKind() ProtoKind
	Export(int, Encoder)
}

type SliceExporter interface {
	Len() int
	ExportItem(int, Encoder)
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
	EncodeSlice(Slice)
	One(Exporter) []byte
	All(SliceExporter) []byte
}
