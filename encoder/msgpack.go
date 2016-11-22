package encoder

import (
	"bytes"
	"encoding/binary"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/types"
	"math"
	"sync"
	"time"
)

var (
	msgBuf   = new(bytes.Buffer)
	msgMutex = new(sync.Mutex)
)

type MsgpackImpl struct {
	buf *bytes.Buffer
}

func NewMsgpack() types.Encoder {
	return &MsgpackImpl{
		buf: msgBuf,
	}
}

func (m *MsgpackImpl) EncodeInt(value int) {
	if !m.encodeFixInt(value) {
		if is.LtEqInt(value, math.MaxInt8) {
			m.encodeUint8(types.MsgpackInt8, uint8(value))
		} else if is.LtEqInt(value, math.MaxInt16) {
			m.encodeUint16(types.MsgpackInt16, uint16(value))
		} else if is.LtEqInt(value, math.MaxInt32) {
			m.encodeUint32(types.MsgpackInt32, uint32(value))
		} else {
			m.encodeUint64(types.MsgpackInt64, uint64(value))
		}
	}
}

func (m *MsgpackImpl) EncodeInt8(value int8) {
	if !m.encodeFixInt(int(value)) {
		m.encodeUint8(types.MsgpackInt8, uint8(value))
	}
}

func (m *MsgpackImpl) EncodeInt16(value int16) {
	if !m.encodeFixInt(int(value)) {
		m.encodeUint16(types.MsgpackInt16, uint16(value))
	}
}

func (m *MsgpackImpl) EncodeInt32(value int32) {
	if !m.encodeFixInt(int(value)) {
		m.encodeUint32(types.MsgpackInt32, uint32(value))
	}
}

func (m *MsgpackImpl) EncodeInt64(value int64) {
	if !m.encodeFixInt(int(value)) {
		m.encodeUint64(types.MsgpackInt64, uint64(value))
	}
}

func (m *MsgpackImpl) encodeFixInt(value int) bool {
	if (is.GtInt(value, math.MinInt8) && is.LtInt(value, 0)) || (is.GtEqInt(value, 0) && is.LtInt(value, math.MaxInt8)) {
		m.buf.WriteByte(byte(value))
		return true
	} else {
		return false
	}
}

func (m *MsgpackImpl) EncodeUint(value uint) {
	if is.LtEqUint(value, uint(math.MaxUint8)) {
		m.encodeUint8(types.MsgpackUint8, uint8(value))
	} else if is.LtEqUint(value, uint(math.MaxUint16)) {
		m.encodeUint16(types.MsgpackUint16, uint16(value))
	} else if is.LtEqUint(value, uint(math.MaxUint32)) {
		m.encodeUint32(types.MsgpackUint32, uint32(value))
	} else {
		m.encodeUint64(types.MsgpackUint64, uint64(value))
	}
}

func (m *MsgpackImpl) EncodeUint8(value uint8) {
	m.encodeUint8(types.MsgpackUint8, value)
}

func (m *MsgpackImpl) encodeUint8(kind byte, value uint8) {
	m.buf.WriteByte(kind)
	m.buf.WriteByte(byte(value))
}

func (m *MsgpackImpl) EncodeUint16(value uint16) {
	m.encodeUint16(types.MsgpackUint16, value)
}

func (m *MsgpackImpl) encodeUint16(kind byte, value uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, value)
	m.buf.WriteByte(kind)
	m.buf.Write(buf)
}

func (m *MsgpackImpl) EncodeUint32(value uint32) {
	m.encodeUint32(types.MsgpackUint32, value)
}

func (m *MsgpackImpl) encodeUint32(kind byte, value uint32) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, value)
	m.buf.WriteByte(kind)
	m.buf.Write(buf)
}

func (m *MsgpackImpl) EncodeUint64(value uint64) {
	m.encodeUint64(types.MsgpackUint64, value)
}

func (m *MsgpackImpl) encodeUint64(kind byte, value uint64) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, value)
	m.buf.WriteByte(kind)
	m.buf.Write(buf)
}

func (m *MsgpackImpl) EncodeFloat32(value float32) {
	m.encodeUint32(types.MsgpackFloat32, math.Float32bits(value))
}

func (m *MsgpackImpl) EncodeFloat64(value float64) {
	m.encodeUint64(types.MsgpackFloat64, math.Float64bits(value))
}

func (m *MsgpackImpl) EncodeString(value string) {
	len := len(value)
	if is.LtEqInt(len, types.MsgpackFixRawLen) {
		m.writeHeaderFix(types.MsgpackFixRaw, len)
	} else if is.LtInt(len, types.MsgpackMax16Bit) {
		m.writeHeader16(types.MsgpackStr16, len)
	} else {
		m.writeHeader32(types.MsgpackStr32, len)
	}
	m.buf.Write([]byte(value))
}

func (m *MsgpackImpl) EncodeBool(value bool) {
	if value {
		m.buf.WriteByte(types.MsgpackTrue)
	} else {
		m.buf.WriteByte(types.MsgpackFalse)
	}
}

func (m *MsgpackImpl) EncodeTime(value time.Time) {
	m.EncodeString(value.Format(types.TimeFormat))
}

func (m *MsgpackImpl) EncodeBytes(value []byte) {
	len := len(value)
	if is.LtEqInt(len, types.MsgpackFixRawLen) {
		m.writeHeaderFix(types.MsgpackFixRaw, len)
	} else if is.LtInt(len, types.MsgpackMax16Bit) {
		m.writeHeader16(types.MsgpackBin16, len)
	} else {
		m.writeHeader32(types.MsgpackBin32, len)
	}
	m.buf.Write(value)
}

func (m *MsgpackImpl) EncodeModel(exporter types.Exporter) {
	propsLen := exporter.Len()
	if is.LtInt(propsLen, types.MsgpackMaxFixMap) {
		m.writeHeaderFix(types.MsgpackFixMap, propsLen)
	} else if is.LtInt(propsLen, types.MsgpackMax16Bit) {
		m.writeHeader16(types.MsgpackMap16, propsLen)
	} else {
		m.writeHeader32(types.MsgpackMap32, propsLen)
	}
	for x := 0; x < propsLen; x++ {
		prop := exporter.Get(x)
		m.EncodeString(prop.GetName())
		exporter.Export(x, m)
	}
}

func (m *MsgpackImpl) writeHeaderFix(kind byte, len int) {
	m.buf.WriteByte(kind | byte(len))
}

func (m *MsgpackImpl) writeHeader16(kind byte, len int) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(len))
	m.buf.WriteByte(kind)
	m.buf.Write(buf)
}

func (m *MsgpackImpl) writeHeader32(kind byte, len int) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(len))
	m.buf.WriteByte(kind)
	m.buf.Write(buf)
}

func (m *MsgpackImpl) EncodeSlice(exporter types.Exporter) {

}

func (m *MsgpackImpl) Encode(exporter types.Exporter) []byte {
	msgMutex.Lock()
	if exporter.GetProtoKind().IsSlice() {
		m.EncodeSlice(exporter)
	} else if exporter.GetProtoKind().IsModel() {
		m.EncodeModel(exporter)
	} else {
		exporter.Export(0, m)
	}
	buf := m.buf.Bytes()
	m.buf.Reset()
	msgMutex.Unlock()
	return buf
}
