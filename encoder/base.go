package encoder

import (
	"bytes"
	"unicode/utf8"
	"strconv"
	"github.com/byorty/hardcore/types"
	"time"
)

var (
	hex = "0123456789abcdef"
)

type BaseImpl struct {
	buf *bytes.Buffer
	exporter types.Exporter
}

func (b *BaseImpl) encodeString(value string) {
	start := 0
	for i := 0; i < len(value); {
		if j := value[i]; j < utf8.RuneSelf {
			if 0x20 <= j && j != '\\' && j != '"' && j != '<' && j != '>' && j != '&' {
				i++
				continue
			}
			if start < i {
				b.buf.WriteString(value[start:i])
			}
			switch j {
			case '\\', '"':
				b.buf.WriteByte('\\')
				b.buf.WriteByte(j)
			case '\n':
				b.buf.WriteByte('\\')
				b.buf.WriteByte('n')
			case '\r':
				b.buf.WriteByte('\\')
				b.buf.WriteByte('r')
			case '\t':
				b.buf.WriteByte('\\')
				b.buf.WriteByte('t')
			default:
				// This encodes bytes < 0x20 except for \n and \r,
				// as well as <, > and &. The latter are escaped because they
				// can lead to security holes when user-controlled strings
				// are rendered into JSON and served to some browsers.
				b.buf.WriteString(`\u00`)
				b.buf.WriteByte(hex[j>>4])
				b.buf.WriteByte(hex[j&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRuneInString(value[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				b.buf.WriteString(value[start:i])
			}
			b.buf.WriteString(`\ufffd`)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				b.buf.WriteString(value[start:i])
			}
			b.buf.WriteString(`\u202`)
			b.buf.WriteByte(hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(value) {
		b.buf.WriteString(value[start:])
	}
}

func (b *BaseImpl) encodeInt64(value int64) {
	b.buf.WriteString(strconv.FormatInt(value, 10))
}

func (b *BaseImpl) encodeUint64(value uint64) {
	b.buf.WriteString(strconv.FormatUint(value, 10))
}

func (b *BaseImpl) encodeFloat32(value float32) {
	b.encodeFloat(float64(value), 32)
}

func (b *BaseImpl) encodeFloat64(value float64) {
	b.encodeFloat(value, 64)
}

func (b *BaseImpl) encodeFloat(value float64, bitSize int) {
	b.buf.WriteString(strconv.FormatFloat(value, 'E', -1, bitSize))
}

func (b *BaseImpl) encodeBool(value bool) {
	b.buf.WriteString(strconv.FormatBool(value))
}

func (b *BaseImpl) encodeTime(value time.Time) {
	b.buf.WriteString(value.Format("2006-01-02 15:04:05"))
}



