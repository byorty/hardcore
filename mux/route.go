package mux

import (
	"github.com/byorty/hardcore"
	"regexp"
	"net/http"
	"strings"
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
	scheme            *regexp.Regexp
	host              *regexp.Regexp
	path              *regexp.Regexp
	headers           map[string]*regexp.Regexp

	schemeTpl         string
	hostTpl           string
	tpl               string
	headerTpls        map[string]string

	beforeMiddlewares []hardcore.MiddlewareFunc
	construct         hardcore.ControllerFunc
	handler           interface{}
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

func (r *Route) prepare() {
	if r.parent != nil {
		if len(r.parent.schemeTpl) > 0 && len(r.schemeTpl) == 0 {
			r.schemeTpl = r.parent.schemeTpl
		}
		if len(r.parent.hostTpl) > 0 && len(r.hostTpl) == 0 {
			r.hostTpl = r.parent.hostTpl
		}
		r.tpl = r.parent.tpl + strings.TrimLeft(r.tpl, "/")
		if len(r.parent.headerTpls) > 0 {
			for key, value := range r.parent.headerTpls {
				if _, ok := r.headerTpls[key]; !ok {
					r.headerTpls[key] = value
				}
			}
		}
		if len(r.parent.beforeMiddlewares) > 0 {
			middlewares := append([]hardcore.MiddlewareFunc{}, r.parent.beforeMiddlewares...)
			r.beforeMiddlewares = append(middlewares, r.beforeMiddlewares...)
		}
		if len(r.parent.afterMiddlewares) > 0 {
			middlewares := append([]hardcore.MiddlewareFunc{}, r.parent.afterMiddlewares...)
			r.afterMiddlewares = append(middlewares, r.afterMiddlewares...)
		}
	}

	switch r.kind {
	case kindAction:

		break
	}
}

func (r *Route) build() {

}

func (r *Route) Match(request *http.Request) bool {

	return true
}
