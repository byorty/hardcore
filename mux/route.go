package mux

import (
	"strings"
	"fmt"
)

const (
	methodGet    = "GET"
	methodPost   = "POST"
	methodPut    = "PUT"
	methodDelete = "DELETE"
)

const (
	defaultScheme = "http"
	defaultHost   = "localhost"
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
	portTpl           string
	tpl               string
	headerTpls        map[string]string

	scopeConstruct    RequestScopeFunc
	beforeMiddlewares []MiddlewareFunc
	construct         ControllerFunc
	handler           interface{}
	afterMiddlewares  []MiddlewareFunc
	routes            []*Route
}

func newRoute() *Route {
	return &Route{
		headerTpls       : make(map[string]string),
		beforeMiddlewares: make([]MiddlewareFunc, 0),
		afterMiddlewares : make([]MiddlewareFunc, 0),
		scopeConstruct   : NewBaseRequestScope,
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

func newRouteByKindAndMethod(kind kindRoute, method, tpl string, handler interface{}) *Route {
	route := newRoute()
	route.kind = kind
	route.method = method
	route.tpl = tpl
	route.handler = handler
	return route
}

func Get(tpl string, handler interface{}) *Route {
	return newRouteByKindAndMethod(kindAction, methodGet, tpl, handler)
}

func Post(tpl string, handler interface{}) *Route {
	return newRouteByKindAndMethod(kindAction, methodPost, tpl, handler)
}

func Put(tpl string, handler interface{}) *Route {
	return newRouteByKindAndMethod(kindAction, methodPut, tpl, handler)
}

func Delete(tpl string, handler interface{}) *Route {
	return newRouteByKindAndMethod(kindAction, methodDelete, tpl, handler)
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

func Controller(tpl string, construct ControllerFunc) *Route {
	route := newRoute()
	route.kind = kindController
	route.tpl = tpl
	route.construct = construct
	return route
}

func (r *Route) Scope(scopeConstruct RequestScopeFunc) *Route {
	r.scopeConstruct = scopeConstruct
	return r
}

func (r *Route) Scheme(tpl string) *Route {
	r.schemeTpl = fmt.Sprintf("{scheme:(%s)+}", tpl)
	return r
}

func (r *Route) Host(tpl string) *Route {
	r.hostTpl = fmt.Sprintf("{host:(%s)+}", tpl)
	return r
}

func (r *Route) Port(tpl string) *Route {
	r.portTpl = fmt.Sprintf("{port:(:%s)+}", tpl)
	return r
}

func (r *Route) Header(key, value string) *Route {
	r.headerTpls[key] = value
	return r
}

func (r *Route) Before(middleware func (RequestScope)) *Route {
	r.beforeMiddlewares = append(r.beforeMiddlewares, middleware)
	return r
}

func (r *Route) Get(tpl string, handler interface{}) *Route {
	return r.Add(newRouteByKindAndMethod(kindControllerAction, methodGet, tpl, handler))
}

func (r *Route) Post(tpl string, handler interface{}) *Route {
	return r.Add(newRouteByKindAndMethod(kindControllerAction, methodPost, tpl, handler))
}

func (r *Route) Put(tpl string, handler interface{})*Route {
	return r.Add(newRouteByKindAndMethod(kindControllerAction, methodPut, tpl, handler))
}

func (r *Route) Delete(tpl string, handler interface{}) *Route {
	return r.Add(newRouteByKindAndMethod(kindControllerAction, methodDelete, tpl, handler))
}

func (r *Route) After(middleware func (RequestScope)) *Route {
	r.afterMiddlewares = append(r.afterMiddlewares, middleware)
	return r
}

func (r *Route) Add(route *Route) *Route {
	route.parent = r
	route.kind = kindControllerAction
	r.routes = append(r.routes, route)
	return r
}

func (r *Route) toMatcher(router *Router) {
	if r.parent != nil {
		if len(r.parent.schemeTpl) > 0 && len(r.schemeTpl) == 0 {
			r.schemeTpl = r.parent.schemeTpl
		}
		if len(r.parent.hostTpl) > 0 && len(r.hostTpl) == 0 {
			r.hostTpl = r.parent.hostTpl
		}
		if len(r.parent.portTpl) > 0 && len(r.portTpl) == 0 {
			r.portTpl = r.parent.portTpl
		}
		r.tpl = r.parent.tpl + strings.TrimRight(r.tpl, "/")
		r.tpl = strings.Replace(r.tpl, "//", "/", -1)
		if len(r.parent.headerTpls) > 0 {
			for key, value := range r.parent.headerTpls {
				if _, ok := r.headerTpls[key]; !ok {
					r.headerTpls[key] = value
				}
			}
		}
		if len(r.parent.beforeMiddlewares) > 0 {
			middlewares := append([]MiddlewareFunc{}, r.parent.beforeMiddlewares...)
			r.beforeMiddlewares = append(middlewares, r.beforeMiddlewares...)
		}
		if len(r.parent.afterMiddlewares) > 0 {
			middlewares := append([]MiddlewareFunc{}, r.parent.afterMiddlewares...)
			r.afterMiddlewares = append(middlewares, r.afterMiddlewares...)
		}
		if r.kind == kindControllerAction {
			r.construct = r.parent.construct
		}
	}

	switch r.kind {
	case kindAction, kindControllerAction:
		matcher := new(Matcher)

		if len(r.schemeTpl) == 0 {
			r.Scheme(defaultScheme)
		}

		if len(r.hostTpl) == 0 {
			r.Host(defaultHost)
		}

		var tpl string
		if len(r.portTpl) > 0 {
			tpl = fmt.Sprintf("^%s://%s%s%s$", r.schemeTpl, r.hostTpl, r.portTpl, r.tpl)
		} else {
			tpl = fmt.Sprintf("^%s://%s%s$", r.schemeTpl, r.hostTpl, r.tpl)
		}

		matcher.urlParams = newParamMatcher(tpl)

		if len(r.headerTpls) > 0 {
			matcher.headers = make([]*HeaderMatcher, 0)
			for name, tpl := range r.headerTpls {
				matcher.headers = append(matcher.headers, newHeaderMatcher(name, tpl))
			}
		}

		matcher.scopeConstruct = r.scopeConstruct
		if len(r.beforeMiddlewares) > 0 {
			matcher.beforeMiddlewares = r.beforeMiddlewares
		}
		if r.construct != nil {
			matcher.construct = r.construct
		}
		matcher.handler = r.handler
		if len(r.afterMiddlewares) > 0 {
			matcher.afterMiddlewares = r.afterMiddlewares
		}

		router.addMatcher(r.method, matcher)
		break
	case kindPath, kindController:
		for _, route := range r.routes {
			route.toMatcher(router)
		}
		break
	}
}
