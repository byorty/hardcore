package importers

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

var (
	TestInt            int
	TestInt8           int8
	TestInt16          int16
	TestInt32          int32
	TestInt64          int64
	TestUint           uint
	TestUint8          uint8
	TestUint16         uint16
	TestUint32         uint32
	TestUint64         uint64
	TestFloat32        float32
	TestFloat64        float32
	_msgpackProperties = map[string]_MsgpackPropertyImpl{
		"i": newMsgpackProperty(types.ProtoIntKind, func(user *models.User, decoder types.Decoder, value []byte) {
			TestInt = decoder.DecodeInt(value)
		}),
		"i8": newMsgpackProperty(types.ProtoInt8Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"i16": newMsgpackProperty(types.ProtoInt16Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"i32": newMsgpackProperty(types.ProtoInt32Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"i64": newMsgpackProperty(types.ProtoInt64Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"u": newMsgpackProperty(types.ProtoUintKind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"u8": newMsgpackProperty(types.ProtoUint8Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"u16": newMsgpackProperty(types.ProtoUint16Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"u32": newMsgpackProperty(types.ProtoUint32Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"u64": newMsgpackProperty(types.ProtoUint64Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"f32": newMsgpackProperty(types.ProtoFloat32Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"f64": newMsgpackProperty(types.ProtoFloat64Kind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"b": newMsgpackProperty(types.ProtoBoolKind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
		"s": newMsgpackProperty(types.ProtoStringKind, func(user *models.User, decoder types.Decoder, value []byte) {

		}),
	}
)
