package encoder

import (
	"bytes"
	"github.com/byorty/hardcore/types"
	"sync"
	"time"
)

var (
	startBrace         = []byte("{")
	endBrace           = []byte("}")
	quotes             = []byte("\"")
	colon              = []byte(":")
	comma              = []byte(",")
	startSquareBracket = []byte("[")
	endSquareBracket   = []byte("]")
	null               = []byte("null")
	jsonBuf            = new(bytes.Buffer)
	jsonMutex          = new(sync.Mutex)
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
	j.encodeExporter(exporter)
	buf := j.buf.Bytes()
	j.buf.Reset()
	jsonMutex.Unlock()
	return buf
}

func (j JsonImpl) encodeExporter(exporter types.Exporter) {
	switch exportable := exporter.GetExportable().(type) {
	case types.Slice:
		j.encodeSlice(exporter, exportable)
	case types.ProtoEntity, types.FormError:
		j.encodeStruct(exporter, exportable)
	default:
		j.buf.Write(null)
	}
}

func (j *JsonImpl) encodeSlice(exporter types.Exporter, slice types.Slice) {
	sliceLen := slice.Len()
	lastIndex := sliceLen - 1
	j.buf.Write(startSquareBracket)
	for i := 0; i < slice.Len(); i++ {
		j.encode(exporter, slice.GetRaw(i))
		if i < lastIndex {
			j.buf.Write(comma)
		}
	}
	j.buf.Write(endSquareBracket)
}

func (j *JsonImpl) encodeStruct(exporter types.Exporter, model interface{}) {
	propsLen := exporter.Len()
	lastIndex := propsLen - 1
	j.buf.Write(startBrace)
	for i := 0; i < propsLen; i++ {
		prop := exporter.Get(i)
		j.buf.Write(quotes)
		j.buf.WriteString(prop.GetName())
		j.buf.Write(quotes)
		j.buf.Write(colon)
		j.encode(exporter, prop.GetValue(model))
		if i < lastIndex {
			j.buf.Write(comma)
		}
	}
	j.buf.Write(endBrace)
}

func (j *JsonImpl) encode(exporter types.Exporter, exportable interface{}) {
	switch value := exportable.(type) {
	case int:
		j.encodeInt64(int64(value))
	case int8:
		j.encodeInt64(int64(value))
	case int16:
		j.encodeInt64(int64(value))
	case int32:
		j.encodeInt64(int64(value))
	case int64:
		j.encodeInt64(value)
	case uint:
		j.encodeUint64(uint64(value))
	case uint8:
		j.encodeUint64(uint64(value))
	case uint16:
		j.encodeUint64(uint64(value))
	case uint32:
		j.encodeUint64(uint64(value))
	case uint64:
		j.encodeUint64(value)
	case float32:
		j.encodeFloat32(value)
	case float64:
		j.encodeFloat64(value)
	case bool:
		j.encodeBool(value)
	case string:
		j.buf.Write(quotes)
		j.encodeString(value)
		j.buf.Write(quotes)
	case time.Time:
		j.buf.Write(quotes)
		j.encodeTime(value)
		j.buf.Write(quotes)
	case types.Exporter:
		j.encodeExporter(value)
	case types.Slice:
		j.encodeSlice(exporter, value)
	case types.Model, types.FormError, types.Named:
		j.encodeStruct(exporter, value)
	case nil:
		j.buf.Write(null)
	}
}
