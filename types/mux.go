package types

import "net/http"

type MiddlewareFunc func(RequestScope)

type ActionController interface {
	CallAction(interface{}, RequestScope)
}

type CallableAction interface {
	Call(interface{}, RequestScope)
}

type ControllerConstructor func() ActionController

type RequestScopeParams interface {
	Set(string, string)
	GetString(string) string
	GetInt(string) int
	GetInt8(string) int8
	GetInt16(string) int16
	GetInt32(string) int32
	GetInt64(string) int64
	GetUint(string) uint
	GetUint8(string) uint8
	GetUint16(string) uint16
	GetUint32(string) uint32
	GetUint64(string) uint64
	GetBool(string) bool
	GetFloat32(string) float32
	GetFloat64(string) float64
}

type RequestScopeConstructor func() RequestScope

type RequestScope interface {
	GetWriter() http.ResponseWriter
	SetWriter(http.ResponseWriter)
	GetRequest() *http.Request
	SetRequest(*http.Request)
	GetPathParams() RequestScopeParams
	SetPathParams(params RequestScopeParams)
	GetHeaderParams() RequestScopeParams
	SetHeaderParams(params RequestScopeParams)
	SetHeaderParam(key, value string)
}
