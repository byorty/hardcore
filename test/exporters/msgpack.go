package exporters

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
	"math"
)

var (
	msgpackProperties = []_MsgpackPropertyImpl{
		newMsgpackProperty("i", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeInt(math.MaxInt16)
		}),
		newMsgpackProperty("i8", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeInt8(int8(math.MaxInt8))
		}),
		newMsgpackProperty("i16", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeInt16(int16(math.MaxInt16))
		}),
		newMsgpackProperty("i32", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeInt32(int32(math.MaxInt32))
		}),
		newMsgpackProperty("i64", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeInt64(int64(math.MaxInt64))
		}),
		newMsgpackProperty("u", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeUint(uint(math.MaxUint16))
		}),
		newMsgpackProperty("u8", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeUint8(uint8(math.MaxUint8))
		}),
		newMsgpackProperty("u16", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeUint16(uint16(math.MaxUint16))
		}),
		newMsgpackProperty("u32", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeUint32(uint32(math.MaxUint32))
		}),
		newMsgpackProperty("u64", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeUint64(uint64(math.MaxUint64))
		}),
		newMsgpackProperty("f32", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeFloat32(float32(math.MaxFloat32))
		}),
		newMsgpackProperty("f64", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeFloat64(math.MaxFloat64)
		}),
		newMsgpackProperty("b", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeBool(true)
		}),
		newMsgpackProperty("s", types.ProtoUnkownKind, func(user *models.User, encoder types.Encoder) {
			encoder.EncodeString("test@test.com")
		}),
	}
)
