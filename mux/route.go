package mux

import (
	"bytes"
	"fmt"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"sort"
	"strings"
)

const (
	methodGet    = "GET"
	methodPost   = "POST"
	methodPut    = "PUT"
	methodDelete = "DELETE"
	methodWebsocket = "websocket"
)

const (
	defaultScheme = "http"
	defaultHost   = "localhost"
)

type Route struct {
	kind   types.RouteKind
	parent *Route

	method     string
	schemeTpl  string
	hostTpl    string
	portTpl    string
	tpl        string
	headerTpls map[string]string

	scopeConstruct    types.RequestScopeConstructor
	beforeMiddlewares []types.MiddlewareFunc
	construct         types.ControllerConstructor
	handler           interface{}
	afterMiddlewares  []types.MiddlewareFunc
	routes            []*Route
}

func newRoute() *Route {
	return &Route{
		headerTpls:        make(map[string]string),
		beforeMiddlewares: make([]types.MiddlewareFunc, 0),
		afterMiddlewares:  make([]types.MiddlewareFunc, 0),
		scopeConstruct:    scope.NewRequest,
	}
}

func newActionRoute(method, tpl string, handler interface{}) *Route {
	route := newRoute()
	route.kind = types.ActionRouteKind
	route.method = method
	route.tpl = tpl
	route.handler = handler
	return route
}

func newRouteByKindAndMethod(kind types.RouteKind, method, tpl string, handler interface{}) *Route {
	route := newRoute()
	route.kind = kind
	route.method = method
	route.tpl = tpl
	route.handler = handler
	return route
}

func Get(tpl string, handler interface{}) types.Route {
	return newRouteByKindAndMethod(types.ActionRouteKind, methodGet, tpl, handler)
}

func Post(tpl string, handler interface{}) types.Route {
	return newRouteByKindAndMethod(types.ActionRouteKind, methodPost, tpl, handler)
}

func Websocket(tpl string, handler interface{}) types.Route {
	return newRouteByKindAndMethod(types.ActionRouteKind, methodWebsocket, tpl, handler)
}

func Put(tpl string, handler interface{}) types.Route {
	return newRouteByKindAndMethod(types.ActionRouteKind, methodPut, tpl, handler)
}

func Delete(tpl string, handler interface{}) types.Route {
	return newRouteByKindAndMethod(types.ActionRouteKind, methodDelete, tpl, handler)
}

func Path(tpl string, subRoutes ...types.Route) types.Route {
	route := newRoute()
	route.kind = types.PathRouteKind
	route.tpl = tpl
	route.routes = make([]*Route, len(subRoutes))
	for i, subRoute := range subRoutes {
		subRoute.SetParent(route)
		route.routes[i] = subRoute.(*Route)
	}
	return route
}

func Controller(tpl string, construct types.ControllerConstructor) types.Route {
	route := newRoute()
	route.kind = types.ControllerRouteKind
	route.tpl = tpl
	route.construct = construct
	return route
}

func (r *Route) Scope(scopeConstruct types.RequestScopeConstructor) types.Route {
	r.scopeConstruct = scopeConstruct
	return r
}

func (r *Route) Scheme(tpl string) types.Route {
	r.schemeTpl = fmt.Sprintf("{scheme:(%s)+}", tpl)
	return r
}

func (r *Route) Host(tpl string) types.Route {
	r.hostTpl = fmt.Sprintf("{host:(%s)+}", tpl)
	return r
}

func (r *Route) Port(tpl string) types.Route {
	r.portTpl = fmt.Sprintf("{port:(:%s)+}", tpl)
	return r
}

func (r *Route) Header(key, value string) types.Route {
	r.headerTpls[key] = value
	return r
}

func (r *Route) Before(middleware types.MiddlewareFunc) types.Route {
	r.beforeMiddlewares = append(r.beforeMiddlewares, middleware)
	return r
}

func (r *Route) Get(tpl string, handler interface{}) types.Route {
	return r.Add(newRouteByKindAndMethod(types.ControllerActionRouteKind, methodGet, tpl, handler))
}

func (r *Route) Post(tpl string, handler interface{}) types.Route {
	return r.Add(newRouteByKindAndMethod(types.ControllerActionRouteKind, methodPost, tpl, handler))
}

func (r *Route) Websocket(tpl string, handler interface{}) types.Route {
	return r.Add(newRouteByKindAndMethod(types.ControllerActionRouteKind, methodWebsocket, tpl, handler))
}

func (r *Route) GetOrPost(tpl string, handler interface{}) types.Route {
	r.Add(newRouteByKindAndMethod(types.ControllerActionRouteKind, methodGet, tpl, handler))
	return r.Add(newRouteByKindAndMethod(types.ControllerActionRouteKind, methodPost, tpl, handler))
}

func (r *Route) Put(tpl string, handler interface{}) types.Route {
	return r.Add(newRouteByKindAndMethod(types.ControllerActionRouteKind, methodPut, tpl, handler))
}

func (r *Route) Delete(tpl string, handler interface{}) types.Route {
	return r.Add(newRouteByKindAndMethod(types.ControllerActionRouteKind, methodDelete, tpl, handler))
}

func (r *Route) After(middleware types.MiddlewareFunc) types.Route {
	r.afterMiddlewares = append(r.afterMiddlewares, middleware)
	return r
}

func (r *Route) Add(route types.Route) types.Route {
	if route != nil {
		route.SetParent(r)
		route.SetKind(types.ControllerActionRouteKind)
		r.routes = append(r.routes, route.(*Route))
	}
	return r
}

func (r *Route) Batch(routes ...types.Route) types.Route {
	for _, route := range routes {
		r.Add(route)
	}
	return r
}

func (r *Route) SetKind(kind types.RouteKind) {
	r.kind = kind
}

func (r *Route) SetParent(parent types.Route) {
	r.parent = parent.(*Route)
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
			middlewares := append([]types.MiddlewareFunc{}, r.parent.beforeMiddlewares...)
			r.beforeMiddlewares = append(middlewares, r.beforeMiddlewares...)
		}
		if len(r.parent.afterMiddlewares) > 0 {
			middlewares := append([]types.MiddlewareFunc{}, r.parent.afterMiddlewares...)
			r.afterMiddlewares = append(middlewares, r.afterMiddlewares...)
		}
		if r.kind == types.ControllerActionRouteKind {
			r.construct = r.parent.construct
		}
	}

	switch r.kind {
	case types.ActionRouteKind, types.ControllerActionRouteKind:
		matcher := &Matcher{
			path:      r.tpl,
			pathLen:   len(r.tpl),
			lastIndex: len(r.tpl) - 1,
		}

		if len(r.schemeTpl) == 0 {
			r.Scheme(defaultScheme)
		}

		if len(r.hostTpl) == 0 {
			r.Host(defaultHost)
		}

		var buf *bytes.Buffer
		tplLen := len(r.tpl)
		for i := 0; i < tplLen; i++ {
			isSlash := '/' == r.tpl[i]
			if buf != nil && !isSlash {
				buf.WriteByte(r.tpl[i])
			}
			if r.tpl[i] == ':' {
				buf = new(bytes.Buffer)
			}
			if isSlash || i == tplLen-1 {
				if buf != nil {
					name := buf.String()
					nameLen := len(name)
					nameLastIndex := nameLen - 1
					var kind ParamMatcherKind
					if existsKind, ok := paramMatcherKindByByte[name[nameLastIndex]]; ok {
						kind = existsKind
						name = name[:nameLastIndex]
						nameLen = len(name)
					} else {
						kind = RequireParamMatcherKind
					}
					matcher.params = append(matcher.params, &ParamMatcher{
						name:      name,
						len:       nameLen,
						lastIndex: nameLen - 1,
						kind:      kind,
					})
					buf = nil
				}
			}
		}

		matcher.paramsLen = len(matcher.params)

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

		if _, ok := router.matchers[r.method]; !ok {
			router.matchers[r.method] = make(Matchers, 0)
		}
		router.matchers[r.method] = append(router.matchers[r.method], matcher)
		sort.Sort(router.matchers[r.method])
		break
	case types.PathRouteKind, types.ControllerRouteKind:
		for _, route := range r.routes {
			route.toMatcher(router)
		}
		break
	}
}
