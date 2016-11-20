package decoder

import (
	"github.com/byorty/hardcore/types"
	"time"
)

type MsgpackImpl struct {
	data []byte
	len  int
}

func NewMsgpack(data []byte) types.Encoder {
	return &MsgpackImpl{
		data: data,
		len:  len(data),
	}
}

func (m *MsgpackImpl) DecodeInt([]byte) int {

}

func (m *MsgpackImpl) DecodeInt8([]byte) int8 {

}

func (m *MsgpackImpl) DecodeInt16([]byte) int16 {

}

func (m *MsgpackImpl) DecodeInt32([]byte) int32 {

}

func (m *MsgpackImpl) DecodeInt64([]byte) int64 {

}

func (m *MsgpackImpl) DecodeUint([]byte) uint {

}

func (m *MsgpackImpl) DecodeUint8([]byte) uint8 {

}

func (m *MsgpackImpl) DecodeUint16([]byte) uint16 {

}

func (m *MsgpackImpl) DecodeUint32([]byte) uint32 {

}

func (m *MsgpackImpl) DecodeUint64([]byte) uint64 {

}

func (m *MsgpackImpl) DecodeFloat32([]byte) float32 {

}

func (m *MsgpackImpl) DecodeFloat64([]byte) float64 {

}

func (m *MsgpackImpl) DecodeString([]byte) string {

}

func (m *MsgpackImpl) DecodeBool([]byte) bool {

}

func (m *MsgpackImpl) DecodeTime([]byte) time.Time {

}

func (m *MsgpackImpl) Decode(importer types.Importer) {
	for i := 0; i < m.len; i++ {

	}
}
