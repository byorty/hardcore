package encoder

import (
	"github.com/byorty/hardcore/types"
	"bytes"
	"time"
	"sync"
)

var (
	startBrace = []byte("{")
	endBrace = []byte("}")
	quotes = []byte("\"")
	colon = []byte(":")
	comma = []byte(",")
	startSquareBracket = []byte("[")
	endSquareBracket = []byte("]")
	null = []byte("null")
	jsonBuf = new(bytes.Buffer)
	jsonMutex = new(sync.Mutex)
)

type JsonImpl struct {
	BaseImpl
}

func NewJson() types.Encoder {
	return &JsonImpl{
		BaseImpl{
			buf: jsonBuf,
		},
	}
}

func (j *JsonImpl) Encode(exporter types.Exporter) []byte {
	jsonMutex.Lock()
	defer jsonMutex.Unlock()

	j.exporter = exporter
	j.buf.Reset()
	j.encodeExporter(j.exporter)
	return j.buf.Bytes()
}

func (j JsonImpl) encodeExporter(exporter types.Exporter) {
	switch exportable := exporter.GetExportable().(type) {
	case types.Slice: j.encodeSlice(exportable)
	case types.Model: j.encodeStruct(exportable)
	default: j.buf.Write(null)
	}
}

func (j *JsonImpl) encodeSlice(slice types.Slice) {
	sliceLen := slice.Len()
	lastIndex := sliceLen - 1
	j.buf.Write(startSquareBracket)
	for i := 0; i < slice.Len(); i++ {
		j.encode(slice.GetRaw(i))
		if i < lastIndex {
			j.buf.Write(comma)
		}
	}
	j.buf.Write(endSquareBracket)
}

func (j *JsonImpl) encodeStruct(model interface{}) {
	propsLen := j.exporter.Len()
	lastIndex := propsLen - 1
	j.buf.Write(startBrace)
	for i := 0; i < propsLen; i++ {
		prop := j.exporter.Get(i)
		j.buf.Write(quotes)
		j.buf.WriteString(prop.GetName())
		j.buf.Write(quotes)
		j.buf.Write(colon)
		j.encode(prop.GetValue(model))
		if i < lastIndex {
			j.buf.Write(comma)
		}
	}
	j.buf.Write(endBrace)
}

func (j *JsonImpl) encode(exportable interface{}) {
	switch value := exportable.(type) {
	case int: j.encodeInt64(int64(value))
	case int8: j.encodeInt64(int64(value))
	case int16: j.encodeInt64(int64(value))
	case int32: j.encodeInt64(int64(value))
	case int64: j.encodeInt64(value)
	case uint: j.encodeUint64(uint64(value))
	case uint8: j.encodeUint64(uint64(value))
	case uint16: j.encodeUint64(uint64(value))
	case uint32: j.encodeUint64(uint64(value))
	case uint64: j.encodeUint64(value)
	case float32: j.encodeFloat32(value)
	case float64: j.encodeFloat64(value)
	case bool: j.encodeBool(value)
	case string:
		j.buf.Write(quotes)
		j.encodeString(value)
		j.buf.Write(quotes)
	case time.Time:
		j.buf.Write(quotes)
		j.encodeTime(value)
		j.buf.Write(quotes)
	case types.Slice: j.encodeSlice(value)
	case types.Model, types.FormError: j.encodeStruct(value)
	case types.Named: j.encode(value.GetRawId())
	case nil: j.buf.Write(null)
	}
}


