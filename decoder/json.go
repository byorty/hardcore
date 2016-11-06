package decoder

import (
	"github.com/byorty/hardcore/types"
	"strconv"
	"github.com/byorty/hardcore/is"
	"unicode/utf8"
	"unicode/utf16"
	"unicode"
	"github.com/byorty/hardcore/scope"
	"errors"
	"bytes"
	"time"
)

const (
	invalidIndex = -1
	jsonSpace = '\x20'
	jsonStartBrace = '{'
	jsonEndBrace = '}'
	jsonStartSquareBracket = '['
	jsonEndSquareBracket = ']'
	jsonDoubleQuotes = '"'
	jsonBackslash = '\\'
	jsonComma = ','
	jsonColon = ':'

)

var (
	jsonTrue = []byte("true")
	jsonFalse = []byte("false")
	jsonNull = []byte("null")

	jsonEmptyBytes = []byte{}
)

var (
	jsonErrCloseQuote = errors.New("incorrect syntax - expect close quote")
	jsonErrValue = errors.New("incorrect syntax - expect value")
	jsonErrTrue = errors.New("incorrect syntax - 'true' expected")
	jsonErrFalse = errors.New("incorrect syntax - 'false' expected")
	jsonErrNull = errors.New("incorrect syntax - 'null' expected")
	jsonErrToken = errors.New("incorrect syntax - unrecognized token")
	jsonErrString = errors.New("incorrect syntax - string contains invalid characters")
	jsonErrDigit = errors.New("incorrect syntax - expected digit")
	jsonErrFractional = errors.New("incorrect syntax - expected fractional")
	jsonErrDigitInExponent = errors.New("incorrect syntax - expected digit in exponential")
)

const arr0Size int = 8

func NewJson() types.Decoder {
	return new(JsonImpl)
}

type JsonImpl struct {
	//deep int
	//cur string // current bytes
	off int    // current offset
	len int
}

func (j *JsonImpl) One(json []byte, importer types.Importer) {

	//j.deep = 0
	//j.decodeObjectOrArray(data, importer)
	j.decode(json, importer)
}

func (j *JsonImpl) decode(json []byte, importer types.Importer) {
	j.len = len(json)
	j.skipSpaces(json)
	if j.len <= j.off {
		scope.App().GetLogger().Error("incorrect syntax - expect value")
		return
	}
	switch json[j.off] {
	case jsonStartBrace:
		j.off++
		j.DecodeModel(importer, json)
	//case jsonStartSquareBracket:
	//	j.off++
	//	j.decodeSlice(json)
	//case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
	//	return j.decodeNumber(json)
	//case 't':
	//	if j.len >= j.off+4 && json[j.off:j.off+4] == jsonTrue {
	//		j.off += 4
	//		return jsonTrue, nil
	//	} else {
	//		return jsonEmptyBytes, errors.New("'true' expected")
	//	}
	//case 'f':
	//	if j.len >= j.off+5 && json[j.off:j.off+5] == jsonFalse {
	//		j.off += 5
	//		return jsonFalse, nil
	//	} else {
	//		return jsonEmptyBytes, errors.New("'false' expected")
	//	}
	//case 'n':
	//	if j.len >= j.off+4 && json[j.off:j.off+4] == jsonNull {
	//		j.off += 4
	//		return jsonNull, nil
	//	} else {
	//		return jsonEmptyBytes, errors.New("'null' expected")
	//	}
	default:
		scope.App().GetLogger().Error("incorrect syntax - unrecognized token")
		//return jsonEmptyBytes, errors.New("incorrect syntax - unrecognized token")
	}
}

func (j *JsonImpl) DecodeModel(importer types.Importer, json []byte) {
	for {
		j.skipSpaces(json)
		if j.len <= j.off {
			scope.App().GetLogger().Error("incorrect syntax - object")
			break
		}
		switch json[j.off] {
		case jsonEndBrace:
			j.off++
			break
		case jsonDoubleQuotes:
			j.off++
			rawKey, err := j.decodeString(json)
			if is.NotNil(err) {
				scope.App().GetLogger().Error(err)
				break
			}
			j.skipSpaces(json)
			if j.len > j.off && json[j.off] == jsonColon {
				j.off++
			} else {
				scope.App().GetLogger().Error("incorrect syntax - expect ':' after object key")
				break
			}
			key := string(rawKey)
			prop, ok := importer.Get(key)
			if ok {
				var value []byte
				var err error
				kind := prop.GetProtoKind()
				switch {
				case kind.IsString(): value, err = j.decodeString(json)
				case kind.IsNumber(): value, err = j.decodeNumber(json)
				//case kind.IsBool():
				}
				//value, err := j.decodeValue(json)

				if is.NotNil(err) {
					scope.App().GetLogger().Error(err)
					break
				}
				importer.Decode(key, j, value)
			}
			j.skipSpaces(json)
			if j.len > j.off && json[j.off] == jsonComma {
				j.off++
			}
		default:
			scope.App().GetLogger().Error("incorrect syntax - expect object key or incomplete object")
			break
		}
	}
}

//func (j *JsonImpl) decodeValue(json []byte) ([]byte, error) {
//	if j.len <= j.off {
//		return jsonEmptyBytes, jsonErrValue
//	}
//	switch json[j.off] {
//	case jsonDoubleQuotes:
//		j.off++
//		return j.decodeString(json)
//	//case jsonStartBrace:
//	//	j.off++
//		//j.DecodeModel(importer, json)
//	//case jsonStartSquareBracket:
//	//	j.off++
//	//	j.decodeSlice(json)
//	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
//		return j.decodeNumber(json)
//	case 't':
//		if j.len >= j.off+4 && json[j.off:j.off+4] == jsonTrue {
//			j.off += 4
//			return jsonTrue, nil
//		} else {
//			return jsonEmptyBytes, jsonErrTrue
//		}
//	case 'f':
//		if j.len >= j.off+5 && json[j.off:j.off+5] == jsonFalse {
//			j.off += 5
//			return jsonFalse, nil
//		} else {
//			return jsonEmptyBytes, jsonErrFalse
//		}
//	case 'n':
//		if j.len >= j.off+4 && json[j.off:j.off+4] == jsonNull {
//			j.off += 4
//			return jsonNull, nil
//		} else {
//			return jsonEmptyBytes, jsonErrNull
//		}
//	default:
//		return jsonEmptyBytes, jsonErrToken
//	}
//}

//func (j *JsonImpl) decodeSlice(json []byte) []interface{} {
//	var arr0 [arr0Size]interface{}
//
//	j.skipSpaces(json)
//
//	if j.len > j.off && json[j.off] == jsonEndSquareBracket {
//		j.off++
//		return []interface{}{}
//	}
//
//	arr0[0] = j.decode(json)
//	if j.err != nil {
//		return arr0[:1]
//	}
//
//	var i int = 1
//SMALL:
//	for j.len > j.off {
//		j.skipSpaces(json)
//		if j.len <= j.off {
//			//s.error("incorrect syntax - incomplete array")
//			return arr0[:i]
//		}
//		switch json[j.off] {
//		case jsonEndSquareBracket:
//			j.off++
//			return arr0[:i]
//		case jsonComma:
//			j.off++
//			j.skipSpaces(json)
//			if i >= arr0Size {
//				break SMALL
//			}
//		default:
//			//s.error("incorrect syntax - incomplete array")
//			return arr0[:i]
//		}
//
//		val := j.decode(json)
//		if j.err != nil {
//			return arr0[:i]
//		}
//		arr0[i] = val
//		i++
//
//		j.skipSpaces(json)
//	}
//
//	if j.len <= j.off {
//		//s.error("incorrect syntax - incomplete array")
//		return arr0[:i]
//	}
//
//	arr := arr0[:]
//
//	for j.len > j.off {
//		if json[j.off] == jsonEndSquareBracket {
//			j.off++
//			return arr
//		}
//
//		val := j.decode(json)
//		//if s.err != nil {
//		//	return arr
//		//}
//		arr = append(arr, val)
//
//		j.skipSpaces(json)
//		if j.len > j.off && json[j.off] == jsonComma {
//			j.off++
//			j.skipSpaces(json)
//		}
//	}
//
//	s.error("incorrect syntax - incomplete array")
//	return arr
//}

func (j *JsonImpl) decodeString(json []byte) ([]byte, error) {
	quotePos := bytes.IndexByte(json[j.off:], jsonDoubleQuotes)
	if quotePos < 0 {
		j.off = j.len
		return jsonEmptyBytes, jsonErrCloseQuote
	}
	quotePos += j.off
	// fast path
	val := json[j.off:quotePos]
	if bytes.IndexByte(val, jsonBackslash) < 0 {
		j.off = quotePos + 1
		return val, nil
	}

	// TODO(vovkasm): rewrite from zero
	// full decoding
	// - find end of string
	for json[quotePos-1] == jsonBackslash {
		n := bytes.IndexByte(json[quotePos+1:], jsonDoubleQuotes)
		if n < 0 {
			j.off = j.len
			return jsonEmptyBytes, jsonErrCloseQuote
		}
		quotePos += n + 1
	}
	// (from standard json package)
	result, ok := unquote(json[j.off:quotePos])
	j.off = quotePos + 1
	if ok {
		return result, nil
	} else {
		return jsonEmptyBytes, jsonErrString
	}
}

// unquote converts a quoted JSON string literal s into an actual string t.
// The rules are different than for Go, so cannot use strconv.Unquote.
func unquote(s []byte) ([]byte, bool) {
	t, ok := unquoteBytes(s)
	return t, ok
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

const charToNum64 int64 = 0x0F

func (j *JsonImpl) decodeNumber(json []byte) ([]byte, error) {
	var startPos = j.off
	//var signMul int64 = 1

	// sign
	if json[j.off] == '-' {
		//signMul = -1
		j.off++
	}

	// significand
	// - integer
	if j.len > j.off && json[j.off] >= '1' && json[j.off] <= '9' {
		j.off++
		for j.len > j.off && json[j.off] >= '0' && json[j.off] <= '9' {
			j.off++
		}
	} else if j.len > j.off && json[j.off] == '0' {
		j.off++
	} else {
		return jsonEmptyBytes, jsonErrDigit
	}

	// - fractional
	//var slowParsing bool
	if j.len > j.off && json[j.off] == '.' {
		//slowParsing = true
		j.off++
		if j.len > j.off && json[j.off] >= '0' && json[j.off] <= '9' {
			j.off++
			for j.len > j.off && json[j.off] >= '0' && json[j.off] <= '9' {
				j.off++
			}
		} else {
			return jsonEmptyBytes, jsonErrFractional
		}
	}

	// exponential
	if j.len > j.off && (json[j.off] == 'e' || json[j.off] == 'E') {
		//slowParsing = true
		j.off++
		if j.len > j.off && (json[j.off] == '+' || json[j.off] == '-') {
			j.off++
		}

		if j.len > j.off && json[j.off] >= '0' && json[j.off] <= '9' {
			j.off++
			for j.len > j.off && json[j.off] >= '0' && json[j.off] <= '9' {
				j.off++
			}
		} else {
			return jsonEmptyBytes, jsonErrDigitInExponent
		}
	}

	//if !slowParsing {
	//	if signMul < 0 {
	//		startPos++
	//	}
	//	var acc int64
	//	switch j.off - startPos {
	//	case 1:
	//		acc = int64(json[startPos]) & charToNum64
	//	case 2:
	//		acc = 10*(int64(json[startPos])&charToNum64) + (int64(json[startPos+1]) & charToNum64)
	//	case 3:
	//		acc = 100*(int64(json[startPos])&charToNum64) + 10*(int64(json[startPos+1])&charToNum64) + int64(json[startPos+2])&charToNum64
	//	case 4:
	//		acc = 1000*(int64(json[startPos])&charToNum64) + 100*(int64(json[startPos+1])&charToNum64) + 10*(int64(json[startPos+2])&charToNum64) + int64(json[startPos+3])&charToNum64
	//	default:
	//		i := j.off - 1
	//		var mul int64 = 1
	//		for i >= startPos {
	//			acc += mul * (int64(json[i]) & charToNum64)
	//			mul *= 10
	//			i--
	//		}
	//	}
	//	return float64(signMul * acc)
	//}

	//val, _ := strconv.ParseFloat(string(json[startPos:j.off]), 64)
	//if err != nil {
	//	s.err = err
	//}

	return json[startPos:j.off], nil
}

func (j *JsonImpl) skipSpaces(json []byte) {
	for j.len > j.off {
		if json[j.off] > jsonSpace {
			break
		} else if json[j.off] == jsonSpace || json[j.off] == '\x0A' || json[j.off] == '\x0D' || json[j.off] == '\x09' {
			j.off++
		} else {
			break
		}
	}
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

func (j *JsonImpl) DecodeUint(value []byte) uint {
	return 0
}

func (j *JsonImpl) DecodeUint8(value []byte) uint8 {
	return 0
}

func (j *JsonImpl) DecodeUint16(value []byte) uint16 {
	return 0
}

func (j *JsonImpl) DecodeUint32(value []byte) uint32 {
	return 0
}

func (j *JsonImpl) DecodeUint64(value []byte) uint64 {
	return 0
}

func (j *JsonImpl) DecodeFloat32(value []byte) float32 {
	return 0
}

func (j *JsonImpl) DecodeFloat64(value []byte) float64 {
	return 0
}

func (j *JsonImpl) DecodeString(value []byte) string {
	return string(value)
}

func (j *JsonImpl) DecodeBool(value []byte) bool {
	return true
}

func (j *JsonImpl) DecodeTime(value []byte) time.Time {
	return time.Now()
}

//func (j *JsonImpl) decodeObjectOrArray(data []byte, importer types.Importer) {
//	dataLen := len(data)
//	for i := 0; i < dataLen; i++ {
//		char := data[i]
//		switch char {
//		case jsonStartBrace:
//			j.deep++
//			objLen := dataLen - 2
//			j.decodeObject(data[i+1:dataLen-1], objLen, importer)
//			i += objLen
//		case jsonEndBrace:
//			j.deep--
//		}
//	}
//}

//func (j *JsonImpl) decodeObject(data []byte, dataLen int, importer types.Importer) {
//	var property types.ImportableProperty
//	start := invalidIndex
//	last := dataLen - 1
//	quotesCount := 0
//
//	for i := 0; i < dataLen; i++ {
//		char := data[i]
//		if is.Nil(property) {
//			switch {
//			case char == jsonDoubleQuotes && start == invalidIndex:
//				start = i + 1
//			case char == jsonDoubleQuotes && start > invalidIndex:
//				if posibleProperty, ok := importer.Get(string(data[start:i])); ok {
//					property = posibleProperty
//				}
//				start = invalidIndex
//			}
//		} else {
//			switch {
//			case char == jsonDoubleQuotes && i < last:
//				if data[i-1] != jsonBackslash {
//					quotesCount++
//				}
//			case start == invalidIndex:
//				start = i + 1
//			case (char == jsonComma && (quotesCount == 0 || quotesCount == 2)) || i == last:
//				var value string
//				if i == last {
//					value = string(data[start:])
//				} else {
//					value = string(data[start:i])
//				}
//				j.decodeValue(importer, property, value)
//				quotesCount = 0
//				property = nil
//				start = invalidIndex
//			}
//		}
//	}
//}

//func (j *JsonImpl) decodeValue(importer types.Importer, property types.ImportableProperty, jsonValue string) {
//	switch property.GetKind() {
//	case types.ScalarImportablePropertyKind:
//		j.decodeScalarValue(importer, property, jsonValue)
//	}
//}
//
//func (j *JsonImpl) decodeScalarValue(importer types.Importer, property types.ImportableProperty, jsonValue string) {
//	jsonValue = strings.Trim(jsonValue, ` `)
//	jsonValue = strings.Trim(jsonValue, `"`)
//	switch property.GetProtoKind() {
//	case types.ProtoIntKind:
//		j.decodeInt(importer, property, jsonValue, 0)
//	case types.ProtoInt8Kind:
//		j.decodeInt(importer, property, jsonValue, 8)
//	case types.ProtoInt16Kind:
//		j.decodeInt(importer, property, jsonValue, 16)
//	case types.ProtoInt32Kind:
//		j.decodeInt(importer, property, jsonValue, 32)
//	case types.ProtoInt64Kind:
//		j.decodeInt(importer, property, jsonValue, 64)
//	case types.ProtoUintKind:
//		j.decodeUint(importer, property, jsonValue, 0)
//	case types.ProtoUint8Kind:
//		j.decodeUint(importer, property, jsonValue, 8)
//	case types.ProtoUint16Kind:
//		j.decodeUint(importer, property, jsonValue, 16)
//	case types.ProtoUint32Kind:
//		j.decodeUint(importer, property, jsonValue, 32)
//	case types.ProtoUint64Kind:
//		j.decodeUint(importer, property, jsonValue, 64)
//	case types.ProtoFloat32Kind:
//		j.decodeFloat(importer, property, jsonValue, 32)
//	case types.ProtoFloat64Kind:
//		j.decodeFloat(importer, property, jsonValue, 64)
//	case types.ProtoStringKind:
//		property.SetValue(importer.GetImportable(), jsonValue)
//	case types.ProtoEnumKind:
//		j.decodeEnum(importer, property, jsonValue)
//	case types.ProtoBoolKind:
//		boolValue, _ := strconv.ParseBool(jsonValue)
//		property.SetValue(importer.GetImportable(), boolValue)
//	case types.ProtoTimeKind:
//	}
//}
//

func (j *JsonImpl) decodeInt(value []byte, bitSize int) int64 {
	intValue, err := strconv.ParseInt(string(value), 10, bitSize)
	if is.Nil(err) {
		return intValue
	} else {
		scope.App().GetLogger().Error(err)
		return 0
	}
}
//
//func (j *JsonImpl) decodeUint(importer types.Importer, property types.ImportableProperty, jsonValue string, bitSize int) {
//	uintValue, _ := strconv.ParseUint(jsonValue, 10, bitSize)
//	property.SetValue(importer.GetImportable(), uintValue)
//}
//
//func (j *JsonImpl) decodeFloat(importer types.Importer, property types.ImportableProperty, jsonValue string, bitSize int) {
//	floatValue, _ := strconv.ParseFloat(jsonValue, bitSize)
//	property.SetValue(importer.GetImportable(), floatValue)
//}
//
//func (j *JsonImpl) decodeEnum(importer types.Importer, property types.ImportableProperty, jsonValue string) {
//	property.SetValue(importer.GetImportable(), jsonValue)
//}



