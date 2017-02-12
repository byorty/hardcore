package decoder

import (
	"encoding/binary"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/types"
	"math"
	"time"
)

type MsgpackImpl struct {
	data []byte
	len  int
}

func NewMsgpack(data []byte) types.Decoder {
	return &MsgpackImpl{
		data: data,
		len:  len(data),
	}
}

func (m *MsgpackImpl) DecodeInt(value []byte) int {
	if i, ok := m.decodeFixInt(value); ok {
		return int(i)
	} else {
		return int(m.DecodeInt32(value))
	}
}

func (m *MsgpackImpl) DecodeInt8(value []byte) int8 {
	return int8(value[0])
}

func (m *MsgpackImpl) DecodeInt16(value []byte) int16 {
	if i, ok := m.decodeFixInt(value); ok {
		return int16(i)
	} else {
		return int16(binary.BigEndian.Uint32(value))
	}
}

func (m *MsgpackImpl) DecodeInt32(value []byte) int32 {
	if i, ok := m.decodeFixInt(value); ok {
		return int32(i)
	} else {
		return int32(binary.BigEndian.Uint32(value))
	}
}

func (m *MsgpackImpl) DecodeInt64(value []byte) int64 {
	if i, ok := m.decodeFixInt(value); ok {
		return int64(i)
	} else {
		return int64(binary.BigEndian.Uint64(value))
	}
}

func (m *MsgpackImpl) decodeFixInt(value []byte) (byte, bool) {
	char := byte(value[0])
	if char >= types.MsgpackPositiveFixInt && char <= types.MsgpackPositiveFixIntMax {
		return char ^ types.MsgpackPositiveFixInt, true
	} else {
		return 0, false
	}
}

func (m *MsgpackImpl) DecodeUint(value []byte) uint {
	if i, ok := m.decodeFixInt(value); ok {
		return uint(i)
	} else {
		return uint(m.DecodeUint32(value))
	}
}

func (m *MsgpackImpl) DecodeUint8(value []byte) uint8 {
	return uint8(value[0])
}

func (m *MsgpackImpl) DecodeUint16(value []byte) uint16 {
	if i, ok := m.decodeFixInt(value); ok {
		return uint16(i)
	} else {
		return binary.BigEndian.Uint16(value)
	}
}

func (m *MsgpackImpl) DecodeUint32(value []byte) uint32 {
	if i, ok := m.decodeFixInt(value); ok {
		return uint32(i)
	} else {
		return binary.BigEndian.Uint32(value)
	}
}

func (m *MsgpackImpl) DecodeUint64(value []byte) uint64 {
	if i, ok := m.decodeFixInt(value); ok {
		return uint64(i)
	} else {
		return binary.BigEndian.Uint64(value)
	}
}

func (m *MsgpackImpl) DecodeFloat32(value []byte) float32 {
	return math.Float32frombits(m.DecodeUint32(value))
}

func (m *MsgpackImpl) DecodeFloat64(value []byte) float64 {
	return math.Float64frombits(m.DecodeUint64(value))
}

func (m *MsgpackImpl) DecodeString(value []byte) string {
	return string(value)
}

func (m *MsgpackImpl) DecodeBool(value []byte) bool {
	return value[0] == types.MsgpackTrue
}

func (m *MsgpackImpl) DecodeTime(value []byte) time.Time {
	if currentTime, err := time.Parse(types.TimeFormat, m.DecodeString(value)); is.Nil(err) {
		return currentTime
	} else {
		return time.Time{}
	}
}

func (m *MsgpackImpl) Decode(importer types.Importer) {
	if m.len > 0 {
		kind := importer.GetProtoKind()
		if kind.IsModel() {
			m.decodeModel(importer)
		} else if kind.IsBase() {
			m.decodeValue(importer, m.data[0], 0, "")
		}
	}
}

func (m *MsgpackImpl) decodeModel(importer types.Importer) {
	var key string
	var numEls int
	state := startState
	for i := 0; i < m.len; i++ {
		char := m.data[i]
		switch {
		case state == startState && char >= types.MsgpackFixMap && char <= types.MsgpackFixRawMax:
			numEls = m.lownibble(char)
			state = startDetectKeyState

		case state == startState && char == types.MsgpackMap16:
			i++
			numEls = int(binary.BigEndian.Uint16(m.data[i : i+1]))
			state = startDetectKeyState
			i++

		case state == startState && char == types.MsgpackMap32:
			i++
			numEls = int(binary.BigEndian.Uint16(m.data[i : i+3]))
			i = i + 3
			state = startDetectKeyState

		case state == startDetectKeyState && numEls > 0:
			switch {
			case m.isFixRaw(char):
				keyBuf, l := m.readFixRaw(i, char)
				key = string(keyBuf)
				if _, ok := importer.Get(key); ok {
					i = i + l
					state = detectValue
				}
			}

		case state == detectValue:
			i, state = m.decodeValue(importer, char, i, key)
		}
	}
}

func (m *MsgpackImpl) decodeValue(importer types.Importer, char byte, i int, key string) (int, decoderState) {
	if m.isFixRaw(char) {
		buf, l := m.readFixRaw(i, char)
		importer.Decode(key, m, buf)
		return i + l, startDetectKeyState
	} else if m.isPositiveFixInt(char) {
		importer.Decode(key, m, []byte{char})
		return i, startDetectKeyState
	} else {
		switch char {
		case types.MsgpackNil, types.MsgpackTrue, types.MsgpackFalse:
			importer.Decode(key, m, []byte{char})
			return i, startDetectKeyState

		case types.MsgpackBin8, types.MsgpackStr8:
			buf, l := m.readHeader8(i)
			importer.Decode(key, m, buf)
			return i + l, startDetectKeyState

		case types.MsgpackBin16, types.MsgpackStr16:
			buf, l := m.readHeader16(i)
			importer.Decode(key, m, buf)
			return i + l, startDetectKeyState

		case types.MsgpackBin32, types.MsgpackStr32:
			buf, l := m.readHeader32(i)
			importer.Decode(key, m, buf)
			return i + l, startDetectKeyState

		case types.MsgpackInt8, types.MsgpackUint8:
			buf, l := m.readNumber8(i)
			importer.Decode(key, m, buf)
			return i + l, startDetectKeyState

		case types.MsgpackInt16, types.MsgpackUint16:
			buf, l := m.readNumber16(i)
			importer.Decode(key, m, buf)
			return i + l, startDetectKeyState

		case types.MsgpackInt32, types.MsgpackUint32, types.MsgpackFloat32:
			buf, l := m.readNumber32(i)
			importer.Decode(key, m, buf)
			return i + l, startDetectKeyState

		case types.MsgpackInt64, types.MsgpackUint64, types.MsgpackFloat64:
			buf, l := m.readNumber64(i)
			importer.Decode(key, m, buf)
			return i + l, startDetectKeyState

		default:
			return i, startDetectKeyState
		}
	}
}

func (m *MsgpackImpl) isFixRaw(char byte) bool {
	return char >= types.MsgpackFixRaw && char <= types.MsgpackFixRawMax
}

func (m *MsgpackImpl) readFixRaw(i int, char byte) ([]byte, int) {
	length := m.getFixRawLen(char)
	start := i + 1
	return m.data[start : start+length], length
}

func (m *MsgpackImpl) getFixRawLen(char byte) int {
	return int(char ^ types.MsgpackFixRaw)
}

func (m *MsgpackImpl) isPositiveFixInt(char byte) bool {
	return char >= types.MsgpackPositiveFixInt && char <= types.MsgpackPositiveFixIntMax
}

func (m *MsgpackImpl) readHeader8(i int) ([]byte, int) {
	start := i + 1
	length := int(m.data[start])
	start += 1
	return m.data[start : start+length], length + 1
}

func (m *MsgpackImpl) readHeader16(i int) ([]byte, int) {
	start := i + 1
	length := int(binary.BigEndian.Uint16(m.data[start : start+2]))
	start += 2
	return m.data[start : start+length], length + 2
}

func (m *MsgpackImpl) readHeader32(i int) ([]byte, int) {
	start := i + 1
	length := int(binary.BigEndian.Uint32(m.data[start : start+4]))
	start += 4
	return m.data[start : start+length], length + 4
}

func (m *MsgpackImpl) readNumber8(i int) ([]byte, int) {
	return []byte{m.data[i+1]}, 1
}

func (m *MsgpackImpl) readNumber16(i int) ([]byte, int) {
	return m.data[i+1 : i+2], 2
}

func (m *MsgpackImpl) readNumber32(i int) ([]byte, int) {
	return m.data[i+1 : i+4], 4
}

func (m *MsgpackImpl) readNumber64(i int) ([]byte, int) {
	return m.data[i+1 : i+8], 8
}

func (m *MsgpackImpl) lownibble(u8 uint8) int {
	return int(u8 & 0xf)
}

// Get the five lowest bits
func (m *MsgpackImpl) lowfive(u8 uint8) int {
	return int(u8 & 0x1f)
}
