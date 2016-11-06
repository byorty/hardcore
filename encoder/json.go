package encoder

import (
	"bytes"
	"github.com/byorty/hardcore/types"
	"sync"
	"time"
	"strconv"
	"unicode/utf8"
)

var (
	jsonHex = "0123456789abcdef"
	jsonStartBrace = []byte("{")
	jsonEndBrace = []byte("}")
	jsonQuotes = []byte("\"")
	jsonColon = []byte(":")
	jsonComma = []byte(",")
	jsonStartSquareBracket = []byte("[")
	jsonEndSquareBracket = []byte("]")
	jsonNull = []byte("null")
	jsonBuf            = new(bytes.Buffer)
	jsonMutex          = new(sync.Mutex)
)

type JsonImpl struct {
	buf *bytes.Buffer
}

func NewJson() types.Encoder {
	return &JsonImpl{
		buf: jsonBuf,
	}
}

func (j *JsonImpl) One(exporter types.Exporter) []byte {
	jsonMutex.Lock()
	j.EncodeModel(exporter)
	buf := j.buf.Bytes()
	j.buf.Reset()
	jsonMutex.Unlock()
	return buf
}

func (j *JsonImpl) All(exporter types.SliceExporter) []byte {
	jsonMutex.Lock()
	j.EncodeSlice(exporter)
	buf := j.buf.Bytes()
	j.buf.Reset()
	jsonMutex.Unlock()
	return buf
}

func (j *JsonImpl) EncodeModel(exporter types.Exporter) {
	propsLen := exporter.Len()
	lastIndex := propsLen - 1
	j.buf.Write(jsonStartBrace)
	for x := 0; x < propsLen; x++ {
		prop := exporter.Get(x)
		j.buf.Write(jsonQuotes)
		j.buf.WriteString(prop.GetName())
		j.buf.Write(jsonQuotes)
		j.buf.Write(jsonColon)
		exporter.Export(x, j)
		if x < lastIndex {
			j.buf.Write(jsonComma)
		}
	}
	j.buf.Write(jsonEndBrace)
}

func (j *JsonImpl) EncodeSlice(exporter types.SliceExporter) {
	sliceLen := exporter.Len()
	sliceLastIndex := sliceLen - 1
	j.buf.Write(jsonStartSquareBracket)
	for x := 0; x < sliceLen; x++ {
		exporter.ExportItem(x, j)
		if x < sliceLastIndex {
			j.buf.Write(jsonComma)
		}
	}
	j.buf.Write(jsonEndSquareBracket)
}

func (j *JsonImpl) EncodeInt(value int) {
	j.EncodeInt64(int64(value))
}

func (j *JsonImpl) EncodeInt8(value int8) {
	j.EncodeInt64(int64(value))
}

func (j *JsonImpl) EncodeInt16(value int16) {
	j.EncodeInt64(int64(value))
}

func (j *JsonImpl) EncodeInt32(value int32) {
	j.EncodeInt64(int64(value))
}

func (j *JsonImpl) EncodeInt64(value int64) {
	j.buf.WriteString(strconv.FormatInt(value, 10))
}

func (j *JsonImpl) EncodeUint(value uint) {
	j.EncodeUint64(uint64(value))
}

func (j *JsonImpl) EncodeUint8(value uint8) {
	j.EncodeUint64(uint64(value))
}

func (j *JsonImpl) EncodeUint16(value uint16) {
	j.EncodeUint64(uint64(value))
}

func (j *JsonImpl) EncodeUint32(value uint32) {
	j.EncodeUint64(uint64(value))
}

func (j *JsonImpl) EncodeUint64(value uint64) {
	j.buf.WriteString(strconv.FormatUint(value, 10))
}

func (j *JsonImpl) EncodeFloat32(value float32) {
	j.encodeFloat(float64(value), 32)
}

func (j *JsonImpl) EncodeFloat64(value float64) {
	j.encodeFloat(value, 64)
}

func (j *JsonImpl) encodeFloat(value float64, bitSize int) {
	j.buf.WriteString(strconv.FormatFloat(value, 'E', -1, bitSize))
}

func (j *JsonImpl) EncodeString(value string) {
	j.buf.Write(jsonQuotes)
	start := 0
	for x := 0; x < len(value); {
		if y := value[x]; y < utf8.RuneSelf {
			if 0x20 <= y && y != '\\' && y != '"' && y != '<' && y != '>' && y != '&' {
				x++
				continue
			}
			if start < x {
				j.buf.WriteString(value[start:x])
			}
			switch y {
			case '\\', '"':
				j.buf.WriteByte('\\')
				j.buf.WriteByte(y)
			case '\n':
				j.buf.WriteByte('\\')
				j.buf.WriteByte('n')
			case '\r':
				j.buf.WriteByte('\\')
				j.buf.WriteByte('r')
			case '\t':
				j.buf.WriteByte('\\')
				j.buf.WriteByte('t')
			default:
				j.buf.WriteString(`\u00`)
				j.buf.WriteByte(jsonHex[y >>4])
				j.buf.WriteByte(jsonHex[y &0xF])
			}
			x++
			start = x
			continue
		}
		c, size := utf8.DecodeRuneInString(value[x:])
		if c == utf8.RuneError && size == 1 {
			if start < x {
				j.buf.WriteString(value[start:x])
			}
			j.buf.WriteString(`\ufffd`)
			x += size
			start = x
			continue
		}
		if c == '\u2028' || c == '\u2029' {
			if start < x {
				j.buf.WriteString(value[start:x])
			}
			j.buf.WriteString(`\u202`)
			j.buf.WriteByte(jsonHex[c&0xF])
			x += size
			start = x
			continue
		}
		x += size
	}
	if start < len(value) {
		j.buf.WriteString(value[start:])
	}
	j.buf.Write(jsonQuotes)
}

func (j *JsonImpl) EncodeBool(value bool) {
	j.buf.WriteString(strconv.FormatBool(value))
}

func (j *JsonImpl) EncodeTime(value time.Time) {
	j.EncodeString(value.Format("2006-01-02 15:04:05"))
}

func (j *JsonImpl) EncodeBytes(value []byte) {
	j.buf.Write(value)
}
