package decoder

import (
	"github.com/byorty/hardcore/types"
	"strings"
	"strconv"
)

const (
	invalidIndex = -1
)

func NewJson() types.Decoder {
	return new(JsonImpl)
}

type JsonImpl struct {
	deep int
}

func (j *JsonImpl) Decode(data []byte, importer types.Importer) {
	j.deep = 0
	j.decodeObjectOrArray(data, importer)
}

func (j *JsonImpl) decodeObjectOrArray(data []byte, importer types.Importer) {
	dataLen := len(data)
	for i := 0;i < dataLen;i++ {
		char := data[i]
		switch char {
		case '{':
			j.deep++
			objLen := dataLen - 2
			j.decodeObject(data[i+1:dataLen-1], objLen, importer)
			i += objLen
		case '}':
			j.deep--
		}
	}
}

func (j *JsonImpl) decodeObject(data []byte, dataLen int, importer types.Importer) {
	var property types.ImportableProperty
	start := invalidIndex
	last := dataLen - 1

	for i := 0;i < dataLen;i++ {
		char := data[i]
		if property == nil {
			switch {
			case char == '"' && start == invalidIndex:
				start = i + 1
			case char == '"' && start > invalidIndex:
				if posibleProperty, ok := importer.Get(string(data[start:i])); ok {
					property = posibleProperty
				}
				start = invalidIndex
			}
		} else {
			switch {
			case start == invalidIndex:
				start = i + 1
			case (char == ',' && i < last && data[i + 1] == '"') || i == last:
				var value string
				if i == last {
					value = string(data[start:])
				} else {
					value = string(data[start:i])
				}
				j.decodeValue(importer, property, value)
				property = nil
				start = invalidIndex
			}
		}
	}
}

func (j *JsonImpl) decodeValue(importer types.Importer, property types.ImportableProperty, jsonValue string) {
	switch property.GetKind() {
	case types.ScalarImportablePropertyKind: j.decodeScalarValue(importer, property, jsonValue)
	}
}

func (j *JsonImpl) decodeScalarValue(importer types.Importer, property types.ImportableProperty, jsonValue string) {
	jsonValue = strings.Trim(jsonValue, `"`)
	switch property.GetProtoKind() {
	case types.ProtoIntKind: j.decodeInt(importer, property, jsonValue, 0)
	case types.ProtoInt8Kind: j.decodeInt(importer, property, jsonValue, 8)
	case types.ProtoInt16Kind: j.decodeInt(importer, property, jsonValue, 16)
	case types.ProtoInt32Kind: j.decodeInt(importer, property, jsonValue, 32)
	case types.ProtoInt64Kind: j.decodeInt(importer, property, jsonValue, 64)
	case types.ProtoUintKind: j.decodeUint(importer, property, jsonValue, 0)
	case types.ProtoUint8Kind: j.decodeUint(importer, property, jsonValue, 8)
	case types.ProtoUint16Kind: j.decodeUint(importer, property, jsonValue, 16)
	case types.ProtoUint32Kind: j.decodeUint(importer, property, jsonValue, 32)
	case types.ProtoUint64Kind: j.decodeUint(importer, property, jsonValue, 64)
	case types.ProtoFloat32Kind: j.decodeFloat(importer, property, jsonValue, 32)
	case types.ProtoFloat64Kind: j.decodeFloat(importer, property, jsonValue, 64)
	case types.ProtoStringKind: property.SetValue(importer.GetImportable(), jsonValue)
	case types.ProtoBoolKind:
		boolValue, _ := strconv.ParseBool(jsonValue)
		property.SetValue(importer.GetImportable(), boolValue)
	case types.ProtoTimeKind:
	}
}

func (j *JsonImpl) decodeInt(importer types.Importer, property types.ImportableProperty, jsonValue string, bitSize int) {
	intValue, _ := strconv.ParseInt(jsonValue, 10, bitSize)
	property.SetValue(importer.GetImportable(), intValue)
}

func (j *JsonImpl) decodeUint(importer types.Importer, property types.ImportableProperty, jsonValue string, bitSize int) {
	uintValue, _ := strconv.ParseUint(jsonValue, 10, bitSize)
	property.SetValue(importer.GetImportable(), uintValue)
}

func (j *JsonImpl) decodeFloat(importer types.Importer, property types.ImportableProperty, jsonValue string, bitSize int) {
	floatValue, _ := strconv.ParseFloat(jsonValue, bitSize)
	property.SetValue(importer.GetImportable(), floatValue)
}