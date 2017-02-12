package types

import "time"

type ImportableProperty interface {
	ProtoEntity
}

type Importer interface {
	ProtoEntity
	Get(string) (ImportableProperty, bool)
	Decode(string, Decoder, []byte)
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
	Decode(Importer)
}
