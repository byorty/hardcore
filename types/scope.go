package types

import (
	"html/template"
	"net/http"
	"time"
)

type ApplicationScope interface {
	GetProjectName() string
	SetProjectName(string) ApplicationScope
	GetRouter() Router
	SetRouter(Router) ApplicationScope
	GetHostname() string
	SetHostname(string) ApplicationScope
	GetPort() int
	SetPort(int) ApplicationScope
	GetReadTimeout() time.Duration
	SetReadTimeout(time.Duration) ApplicationScope
	GetWriteTimeout() time.Duration
	SetWriteTimeout(time.Duration) ApplicationScope
	GetRootPath() string
	SetRootPath(string) ApplicationScope
	GetLogger() Logger
	SetLogger(Logger) ApplicationScope
	GetDAOs() []ModelDAO
	SetDAOs([]ModelDAO) ApplicationScope
	GetTmplPath() string
	SetTmplPath(string) ApplicationScope
	GetTmplCache() map[string]*template.Template
	SetTmplCache(map[string]*template.Template) ApplicationScope
	GetCookieName() string
	SetCookieName(string) ApplicationScope
	GetEnableSession() bool
	SetEnableSession(bool) ApplicationScope
	GetSessionProvider() Cache
	SetSessionProvider(Cache) ApplicationScope
}

type SessionScope interface {
	Get(string) interface{}
	GetId() string
	Has(string) bool
	Remove(string) SessionScope
	Set(string, interface{}) SessionScope
}

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
	FormVerifiable
	GetWriter() http.ResponseWriter
	SetWriter(http.ResponseWriter)
	GetRequest() *http.Request
	SetRequest(*http.Request)
	GetPathParams() RequestScopeParams
	SetPathParams(params RequestScopeParams)
	GetHeaderParams() RequestScopeParams
	SetHeaderParams(params RequestScopeParams)
	SetHeaderParam(key, value string)
	GetSession() SessionScope
	SetSession(SessionScope)
}
