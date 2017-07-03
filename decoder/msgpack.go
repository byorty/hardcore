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
	char byte
}

func NewMsgpack(data []byte) types.Decoder {
	return &MsgpackImpl{
		data: data,
		len:  len(data),
	}
}

func (m *MsgpackImpl) DecodeInt(value []byte) int {
	if m.isPositiveFixInt(m.char) {
		return int(value[0])
	} else {
		switch m.char {
		case types.MsgpackInt8, types.MsgpackUint8:
			return int(value[0])
		case types.MsgpackInt16, types.MsgpackUint16:
			return int(m.DecodeUint16(value))
		case types.MsgpackInt32, types.MsgpackUint32:
			return int(m.DecodeUint32(value))
		case types.MsgpackInt64, types.MsgpackUint64:
			return int(m.DecodeUint64(value))
		default:
			return 0
		}
	}
}

func (m *MsgpackImpl) DecodeInt8(value []byte) int8 {
	return int8(value[0])
}

func (m *MsgpackImpl) DecodeInt16(value []byte) int16 {
	return int16(m.DecodeUint16(value))
}

func (m *MsgpackImpl) DecodeInt32(value []byte) int32 {
	return int32(m.DecodeUint32(value))
}

func (m *MsgpackImpl) DecodeInt64(value []byte) int64 {
	return int64(m.DecodeUint64(value))
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
	if m.isPositiveFixInt(m.char) {
		return uint(value[0])
	} else {
		switch m.char {
		case types.MsgpackUint8:
			return uint(value[0])
		case types.MsgpackUint16:
			return uint(m.DecodeUint16(value))
		case types.MsgpackUint32:
			return uint(m.DecodeUint32(value))
		case types.MsgpackUint64:
			return uint(m.DecodeUint64(value))
		default:
			return 0
		}
	}
}

func (m *MsgpackImpl) DecodeUint8(value []byte) uint8 {
	return uint8(value[0])
}

func (m *MsgpackImpl) DecodeUint16(value []byte) uint16 {
	return binary.BigEndian.Uint16(value)
}

func (m *MsgpackImpl) DecodeUint32(value []byte) uint32 {
	return binary.BigEndian.Uint32(value)
}

func (m *MsgpackImpl) DecodeUint64(value []byte) uint64 {
	return binary.BigEndian.Uint64(value)
}

func (m *MsgpackImpl) DecodeFloat32(value []byte) float32 {
	if m.isPositiveFixInt(m.char) {
		return float32(value[0])
	} else {
		return math.Float32frombits(m.DecodeUint32(value))
	}
}

func (m *MsgpackImpl) DecodeFloat64(value []byte) float64 {
	if m.isPositiveFixInt(m.char) {
		return float64(value[0])
	} else {
		return math.Float64frombits(m.DecodeUint64(value))
	}
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
			m.decodeValue(importer, 0, "")
		}
	}
}

func (m *MsgpackImpl) decodeModel(importer types.Importer) {
	var key string
	var numEls int
	state := startState
	for i := 0; i < m.len; i++ {
		m.char = m.data[i]
		switch {
		case state == startState && m.char >= types.MsgpackFixMap && m.char <= types.MsgpackFixRawMax:
			numEls = m.lownibble(m.char)
			state = startDetectKeyState

		case state == startState && m.char == types.MsgpackMap16:
			i++
			numEls = int(binary.BigEndian.Uint16(m.data[i : i+1]))
			state = startDetectKeyState
			i++

		case state == startState && m.char == types.MsgpackMap32:
			i++
			numEls = int(binary.BigEndian.Uint16(m.data[i : i+3]))
			i = i + 3
			state = startDetectKeyState

		case state == startDetectKeyState && numEls > 0:
			switch {
			case m.isFixRaw(m.char):
				keyBuf, l := m.readFixRaw(i, m.char)
				key = string(keyBuf)
				if _, ok := importer.Get(key); ok {
					i = i + l
					state = detectValue
				}
			}

		case state == detectValue:
			i, state = m.decodeValue(importer, i, key)
		}
	}
}

func (m *MsgpackImpl) decodeValue(importer types.Importer, i int, key string) (int, decoderState) {
	if m.isFixRaw(m.char) {
		buf, l := m.readFixRaw(i, m.char)
		importer.Decode(key, m, buf)
		return i + l, startDetectKeyState
	} else if m.isPositiveFixInt(m.char) {
		importer.Decode(key, m, []byte{m.char})
		return i, startDetectKeyState
	} else {
		switch m.char {
		//case types.MsgpackNil:
		case types.MsgpackTrue, types.MsgpackFalse:
			importer.Decode(key, m, []byte{m.char})
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

			importer.Decode(key, m, m.data[i+1:])
			return i + 1, startDetectKeyState

		case types.MsgpackInt16, types.MsgpackUint16:
			importer.Decode(key, m, m.data[i+1:])
			return i + 2, startDetectKeyState

		case types.MsgpackInt32, types.MsgpackUint32, types.MsgpackFloat32:
			importer.Decode(key, m, m.data[i+1:])
			return i + 4, startDetectKeyState

		case types.MsgpackInt64, types.MsgpackUint64, types.MsgpackFloat64:
			importer.Decode(key, m, m.data[i+1:])
			return i + 8, startDetectKeyState

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

func (m *MsgpackImpl) lownibble(u8 uint8) int {
	return int(u8 & 0xf)
}

// Get the five lowest bits
func (m *MsgpackImpl) lowfive(u8 uint8) int {
	return int(u8 & 0x1f)
}
