package decoder

import (
	"bytes"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"strconv"
	"time"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

const (
	jsonStartBrace         = '{'
	jsonEndBrace           = '}'
	jsonStartSquareBracket = '['
	jsonEndSquareBracket   = ']'
	jsonDoubleQuotes       = '"'
	jsonBackslash          = '\\'
	jsonComma              = ','
	jsonColon              = ':'
)

type decoderState int

const (
	startState decoderState = iota
	decodeModelState
	startDetectKeyState
	endDetectKeyState
	detectValue
)

func NewJson(data []byte) types.Decoder {
	return &JsonImpl{
		data: data,
		len:  len(data),
	}
}

type JsonImpl struct {
	data []byte
	len  int
}

func (j *JsonImpl) Decode(importer types.Importer) {
	var key string
	var start, quoteCount, deep int
	state := startState
	var prop types.ImportableProperty
	for i := 0; i < j.len; i++ {
		char := j.data[i]
		switch {
		case state == startState && char == jsonStartBrace:
			state = startDetectKeyState
		case state == startState && char == jsonStartSquareBracket:

		case state == startDetectKeyState && char == jsonDoubleQuotes:
			start = i + 1
			state = endDetectKeyState
		case state == endDetectKeyState && char == jsonDoubleQuotes:
			key = string(j.data[start:i])
			if existsProp, ok := importer.Get(key); ok {
				prop = existsProp
			}
		case state == endDetectKeyState && char == jsonColon && is.NotNil(prop):
			start = i + 1
			state = detectValue
		case state == detectValue && char == jsonDoubleQuotes && prop.GetProtoKind().IsString():
			quoteCount++
		case state == detectValue:
			kind := prop.GetProtoKind()
			isEndOfVal := char == jsonComma || char == jsonEndBrace || char == jsonEndSquareBracket
			switch {
			case (kind.IsNumber() || kind.IsBool() || kind.IsNotStringEnum()) && isEndOfVal:
				importer.Decode(key, j, bytes.TrimSpace(j.data[start:i]))
				prop = nil
				state = startDetectKeyState
			case (kind.IsString() || kind.IsStringEnum()) && quoteCount == 2 && isEndOfVal:
				value := bytes.TrimSpace(j.data[start:i])
				value = bytes.Trim(j.data[start:i], `"`)
				importer.Decode(key, j, value)
				prop = nil
				quoteCount = 0
				state = startDetectKeyState
			case (kind.IsModel() && char == jsonStartBrace) || (kind.IsSlice() && char == jsonStartSquareBracket):
				deep++
			case (kind.IsModel() && char == jsonEndBrace) || (kind.IsSlice() && char == jsonEndSquareBracket):
				deep--
			case (kind.IsModel() || kind.IsSlice()) && deep == 0 && isEndOfVal:
				importer.Decode(key, j, bytes.TrimSpace(j.data[start:i]))
				prop = nil
				state = startDetectKeyState
			}
		}
	}
}

func unquoteBytes(s []byte) (t []byte, ok bool) {
	// Check for unusual characters. If there are none,
	// then no unquoting is needed, so return a slice of the
	// original bytes.
	r := 0
	for r < len(s) {
		c := s[r]
		if c == jsonBackslash || c == jsonDoubleQuotes || c < ' ' {
			break
		}
		if c < utf8.RuneSelf {
			r++
			continue
		}
		rr, size := utf8.DecodeRune(s[r:])
		if rr == utf8.RuneError && size == 1 {
			break
		}
		r += size
	}
	if r == len(s) {
		return s, true
	}

	b := make([]byte, len(s)+2*utf8.UTFMax)
	w := copy(b, s[0:r])
	for r < len(s) {
		// Out of room?  Can only happen if s is full of
		// malformed UTF-8 and we're replacing each
		// byte with RuneError.
		if w >= len(b)-2*utf8.UTFMax {
			nb := make([]byte, (len(b)+utf8.UTFMax)*2)
			copy(nb, b[0:w])
			b = nb
		}
		switch c := s[r]; {
		case c == jsonBackslash:
			r++
			if r >= len(s) {
				return
			}
			switch s[r] {
			default:
				return
			case jsonDoubleQuotes, jsonBackslash, '/', '\'':
				b[w] = s[r]
				r++
				w++
			case 'b':
				b[w] = '\b'
				r++
				w++
			case 'f':
				b[w] = '\f'
				r++
				w++
			case 'n':
				b[w] = '\n'
				r++
				w++
			case 'r':
				b[w] = '\r'
				r++
				w++
			case 't':
				b[w] = '\t'
				r++
				w++
			case 'u':
				r--
				rr := getu4(s[r:])
				if rr < 0 {
					return
				}
				r += 6
				if utf16.IsSurrogate(rr) {
					rr1 := getu4(s[r:])
					if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
						// A valid pair; consume.
						r += 6
						w += utf8.EncodeRune(b[w:], dec)
						break
					}
					// Invalid surrogate; fall back to replacement rune.
					rr = unicode.ReplacementChar
				}
				w += utf8.EncodeRune(b[w:], rr)
			}

		// Quote, control characters are invalid.
		case c == jsonDoubleQuotes, c < ' ':
			return

		// ASCII
		case c < utf8.RuneSelf:
			b[w] = c
			r++
			w++

		// Coerce to well-formed UTF-8.
		default:
			rr, size := utf8.DecodeRune(s[r:])
			r += size
			w += utf8.EncodeRune(b[w:], rr)
		}
	}
	return b[0:w], true
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.
func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != jsonBackslash || s[1] != 'u' {
		return -1
	}
	r, err := strconv.ParseUint(string(s[2:6]), 16, 64)
	if err != nil {
		return -1
	}
	return rune(r)
}

func (j *JsonImpl) DecodeInt(value []byte) int {
	return int(j.decodeInt(value, 0))
}

func (j *JsonImpl) DecodeInt8(value []byte) int8 {
	return int8(j.decodeInt(value, 8))
}

func (j *JsonImpl) DecodeInt16(value []byte) int16 {
	return int16(j.decodeInt(value, 16))
}

func (j *JsonImpl) DecodeInt32(value []byte) int32 {
	return int32(j.decodeInt(value, 32))
}

func (j *JsonImpl) DecodeInt64(value []byte) int64 {
	return j.decodeInt(value, 64)
}

func (j *JsonImpl) decodeInt(value []byte, bitSize int) int64 {
	intValue, err := strconv.ParseInt(string(value), 10, bitSize)
	if is.Nil(err) {
		return intValue
	} else {
		scope.App().GetLogger().Error(err)
		return 0
	}
}

func (j *JsonImpl) DecodeUint(value []byte) uint {
	return uint(j.decodeUint(value, 0))
}

func (j *JsonImpl) DecodeUint8(value []byte) uint8 {
	return uint8(j.decodeUint(value, 8))
}

func (j *JsonImpl) DecodeUint16(value []byte) uint16 {
	return uint16(j.decodeUint(value, 16))
}

func (j *JsonImpl) DecodeUint32(value []byte) uint32 {
	return uint32(j.decodeUint(value, 32))
}

func (j *JsonImpl) DecodeUint64(value []byte) uint64 {
	return j.decodeUint(value, 64)
}

func (j *JsonImpl) decodeUint(value []byte, bitSize int) uint64 {
	uintValue, err := strconv.ParseUint(string(value), 10, bitSize)
	if is.Nil(err) {
		return uintValue
	} else {
		scope.App().GetLogger().Error(err)
		return 0
	}
}

func (j *JsonImpl) DecodeFloat32(value []byte) float32 {
	return float32(j.decodeFloat(value, 32))
}

func (j *JsonImpl) DecodeFloat64(value []byte) float64 {
	return j.decodeFloat(value, 64)
}

func (j *JsonImpl) decodeFloat(value []byte, bitSize int) float64 {
	floatValue, err := strconv.ParseFloat(string(value), bitSize)
	if is.Nil(err) {
		return floatValue
	} else {
		scope.App().GetLogger().Error(err)
		return 0
	}
}

func (j *JsonImpl) DecodeString(value []byte) string {
	return string(value)
}

func (j *JsonImpl) DecodeBool(value []byte) bool {
	boolValue, err := strconv.ParseBool(string(value))
	if is.Nil(err) {
		return boolValue
	} else {
		scope.App().GetLogger().Error(err)
		return false
	}
}

func (j *JsonImpl) DecodeTime(value []byte) time.Time {
	timeValue, err := time.Parse("2006-12-21 15:04:05", string(value))
	if is.Nil(err) {
		return timeValue
	} else {
		scope.App().GetLogger().Error(err)
		return time.Time{}
	}
}
