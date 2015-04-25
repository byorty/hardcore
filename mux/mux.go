package mux

import (
	"net/http"
	"strconv"
)

type MiddlewareFunc func(RequestScope)

type ActionController interface {
	CallAction(interface{}, RequestScope)
}

type ControllerFunc func() ActionController

type RequestScopeParams map[string]string

func (r RequestScopeParams) GetString(key string) string {
	if value, ok := r[key]; ok {
		return value
	} else {
		return ""
	}
}

func (r RequestScopeParams) GetInt(key string) int {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 0)
	return int(i)
}

func (r RequestScopeParams) GetInt8(key string) int8 {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 8)
	return int8(i)
}

func (r RequestScopeParams) GetInt16(key string) int16 {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 16)
	return int16(i)
}

func (r RequestScopeParams) GetInt32(key string) int32 {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 32)
	return int32(i)
}

func (r RequestScopeParams) GetInt64(key string) int64 {
	i, _ := strconv.ParseInt(r.GetString(key), 10, 64)
	return i
}

func (r RequestScopeParams) GetUint(key string) uint {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 0)
	return uint(i)
}

func (r RequestScopeParams) GetUint8(key string) uint8 {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 8)
	return uint8(i)
}

func (r RequestScopeParams) GetUint16(key string) uint16 {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 16)
	return uint16(i)
}

func (r RequestScopeParams) GetUint32(key string) uint32 {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 32)
	return uint32(i)
}

func (r RequestScopeParams) GetUint64(key string) uint64 {
	i, _ := strconv.ParseUint(r.GetString(key), 10, 64)
	return i
}

func (r RequestScopeParams) GetBool(key string) bool {
	b, _ := strconv.ParseBool(r.GetString(key))
	return b
}

func (r RequestScopeParams) GetFloat32(key string) float32 {
	f, _ := strconv.ParseFloat(r.GetString(key), 32)
	return float32(f)
}

func (r RequestScopeParams) GetFloat64(key string) float64 {
	f, _ := strconv.ParseFloat(r.GetString(key), 64)
	return f
}

type RequestScopeFunc func() RequestScope

type RequestScope interface {
	GetWriter() http.ResponseWriter
	setWriter(http.ResponseWriter)
	GetRequest() *http.Request
	setRequest(*http.Request)
	GetPathParams() RequestScopeParams
	setPathParams(params RequestScopeParams)
	GetHeaderParams() RequestScopeParams
	setHeaderParams(params RequestScopeParams)
	setHeaderParam(key, value string)
}

type BaseRequestScope struct {
	writer       http.ResponseWriter
	request      *http.Request
	pathParams   RequestScopeParams
	headerParams RequestScopeParams
}

func NewBaseRequestScope() RequestScope {
	return new(BaseRequestScope)
}

func (b BaseRequestScope) GetWriter() http.ResponseWriter {
	return b.writer
}

func (b *BaseRequestScope) setWriter(writer http.ResponseWriter) {
	b.writer = writer
}

func (b BaseRequestScope) GetRequest() *http.Request {
	return b.request
}

func (b *BaseRequestScope) setRequest(request *http.Request) {
	b.request = request
}

func (b BaseRequestScope) GetPathParams() RequestScopeParams {
	return b.pathParams
}

func (b *BaseRequestScope) setPathParams(params RequestScopeParams) {
	b.pathParams = params
}

func (b BaseRequestScope) GetHeaderParams() RequestScopeParams {
	return b.headerParams
}

func (b *BaseRequestScope) setHeaderParams(params RequestScopeParams) {
	b.headerParams = params
}

func (b *BaseRequestScope) setHeaderParam(key, value string) {
	b.headerParams[key] = value
}
