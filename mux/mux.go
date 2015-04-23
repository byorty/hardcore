package mux

import (
	"net/url"
	"net/http"
	"strconv"
)

type MiddlewareFunc func(*RequestScope)

type ActionController interface {
	CallAction(interface{}, *RequestScope)
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

func (r RequestScopeParams) GetInt(key string) (int, error) {
	i, err := strconv.ParseInt(r.GetString(key), 10, 0)
	return int(i), err
}

func (r RequestScopeParams) GetInt8(key string) (int8, error) {
	i, err := strconv.ParseInt(r.GetString(key), 10, 8)
	return int8(i), err
}

func (r RequestScopeParams) GetInt16(key string) (int16, error) {
	i, err := strconv.ParseInt(r.GetString(key), 10, 16)
	return int16(i), err
}

func (r RequestScopeParams) GetInt32(key string) (int32, error) {
	i, err := strconv.ParseInt(r.GetString(key), 10, 32)
	return int32(i), err
}

func (r RequestScopeParams) GetInt64(key string) (int64, error) {
	return strconv.ParseInt(r.GetString(key), 10, 64)
}

func (r RequestScopeParams) GetUint(key string) (uint, error) {
	i, err := strconv.ParseUint(r.GetString(key), 10, 0)
	return uint(i), err
}

func (r RequestScopeParams) GetUint8(key string) (uint8, error) {
	i, err := strconv.ParseUint(r.GetString(key), 10, 8)
	return uint8(i), err
}

func (r RequestScopeParams) GetUint16(key string) (uint16, error) {
	i, err := strconv.ParseUint(r.GetString(key), 10, 16)
	return uint16(i), err
}

func (r RequestScopeParams) GetUint32(key string) (uint32, error) {
	i, err := strconv.ParseUint(r.GetString(key), 10, 32)
	return uint32(i), err
}

func (r RequestScopeParams) GetUint64(key string) (uint64, error) {
	return strconv.ParseUint(r.GetString(key), 10, 64)
}

func (r RequestScopeParams) GetBool(key string) (bool, error) {
	return strconv.ParseBool(r.GetString(key))
}

func (r RequestScopeParams) GetFloat32(key string) (float32, error) {
	float, err := strconv.ParseFloat(r.GetString(key), 32)
	return float32(float), err
}

func (r RequestScopeParams) GetFloat64(key string) (float64, error) {
	return strconv.ParseFloat(r.GetString(key), 64)
}

type RequestScope struct {
	UrlStr       string
	Url          *url.URL
	Writer       http.ResponseWriter
	Request      *http.Request
	PathParams   RequestScopeParams
	HeaderParams RequestScopeParams
}
