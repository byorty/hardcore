package types

import "net/http"

type MiddlewareFunc func(RequestScope)

type ActionController interface {
	CallAction(interface{}, RequestScope)
}

type ExtendedViewer interface {
	GetHeaderTmpls() []string
	GetFooterTmpls() []string
}

type CallableAction interface {
	Call(interface{}, RequestScope)
}

type ControllerConstructor func() ActionController

type RouteKind int

const (
	ActionRouteKind RouteKind = iota
	PathRouteKind
	ControllerRouteKind
	ControllerActionRouteKind
)

type Route interface {
	Scope(RequestScopeConstructor) Route
	Scheme(string) Route
	Host(string) Route
	Port(string) Route
	Header(string, string) Route
	Before(MiddlewareFunc) Route
	Get(string, interface{}) Route
	Post(string, interface{}) Route
	GetOrPost(string, interface{}) Route
	Put(string, interface{}) Route
	Delete(string, interface{}) Route
	After(MiddlewareFunc) Route
	Add(Route) Route
	Batch(...Route) Route
	SetKind(RouteKind)
	SetParent(Route)
}

type Router interface {
	http.Handler
	Add(Route) Router
	Batch(...Route) Router
	NotFound(func(RequestScope)) Router
}
