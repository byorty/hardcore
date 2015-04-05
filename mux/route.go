package mux

import (
	"github.com/byorty/hardcore"
	"net/http"
)

const (
	methodGet    = "GET"
	methodPost   = "POST"
	methodPut    = "PUT"
	methodDelete = "DELETE"
)

type kindRoute int

const (
	kindAction           kindRoute = iota
	kindPath
	kindController
	kindControllerAction
)

type Route struct {
	kind              kindRoute
	parent            *Route
	method            string
	schemeTpl         string
	hostTpl           string
	tpl               string
	headerTpls        map[string]string
	beforeMiddlewares []hardcore.MiddlewareFunc
	construct         hardcore.ControllerFunc
	handler           interface{}
	notFoundHandler   http.Handler
	afterMiddlewares  []hardcore.MiddlewareFunc
	routes            []*Route
}

func newRoute() *Route {
	return &Route{
		headerTpls       : make(map[string]string),
		beforeMiddlewares: make([]hardcore.MiddlewareFunc, 0),
		afterMiddlewares : make([]hardcore.MiddlewareFunc, 0),
	}
}

func newActionRoute(method, tpl string, handler interface{}) *Route {
	route := newRoute()
	route.kind = kindAction
	route.method = method
	route.tpl = tpl
	route.handler = handler
	return route
}

func Get(tpl string, handler interface{}) *Route {
	return newActionRoute(methodGet, tpl, handler)
}

func Post(tpl string, handler interface{}) *Route {
	return newActionRoute(methodPost, tpl, handler)
}

func Put(tpl string, handler interface{}) *Route {
	return newActionRoute(methodPut, tpl, handler)
}

func Delete(tpl string, handler interface{}) *Route {
	return newActionRoute(methodDelete, tpl, handler)
}

func Path(tpl string, subRoutes ...*Route) *Route {
	route := newRoute()
	route.kind = kindPath
	route.tpl = tpl
	route.routes = make([]*Route, len(subRoutes))
	for i, subRoute := range subRoutes {
		subRoute.parent = route
		route.routes[i] = subRoute
	}
	return route
}

func Controller(tpl string, construct hardcore.ControllerFunc) *Route {
	route := newRoute()
	route.kind = kindController
	route.tpl = tpl
	route.construct = construct
	return route
}

func (r *Route) Scheme(tpl string) *Route {
	r.schemeTpl = tpl
	return r
}

func (r *Route) Host(tpl string) *Route {
	r.hostTpl = tpl
	return r
}

func (r *Route) Header(key, value string) *Route {
	r.headerTpls[key] = value
	return r
}

func (r *Route) Before(middleware hardcore.MiddlewareFunc) *Route {
	r.beforeMiddlewares = append(r.beforeMiddlewares, middleware)
	return r
}

func (r *Route) Get(tpl string, handler interface{}) *Route {
	r.addAction(methodGet, tpl, handler)
	return r
}

func (r *Route) Post(tpl string, handler interface{}) *Route {
	r.addAction(methodPost, tpl, handler)
	return r
}

func (r *Route) Put(tpl string, handler interface{})*Route {
	r.addAction(methodPut, tpl, handler)
	return r
}

func (r *Route) Delete(tpl string, handler interface{}) *Route {
	r.addAction(methodDelete, tpl, handler)
	return r
}

func (r *Route) After(middleware hardcore.MiddlewareFunc) *Route {
	r.afterMiddlewares = append(r.afterMiddlewares, middleware)
	return r
}

func (r *Route) addAction(method, tpl string, handler interface{}) *Route {
	route := newRoute()
	route.parent = r
	route.kind = kindControllerAction
	route.method = method
	route.tpl = tpl
	route.handler = handler
	r.routes = append(r.routes, route)
	return r
}

func (r *Route) NotFound(handler http.Handler) *Route {
	r.notFoundHandler = handler
	return r
}
