package mux

import (
	"net/http"
	"github.com/byorty/hardcore/types"
	"strings"
)

type Router struct {
	gets         Matchers
	posts        Matchers
	puts         Matchers
	deletes      Matchers
	notFoundFunc types.MiddlewareFunc
}

func NewRouter(routes ...*Route) *Router {
	router := &Router{
		gets        : make(Matchers, 0),
		posts       : make(Matchers, 0),
		puts        : make(Matchers, 0),
		deletes     : make(Matchers, 0),
		notFoundFunc: func (scope types.RequestScope) {
			scope.GetWriter().WriteHeader(http.StatusNotFound)
			scope.GetWriter().Write([]byte("not found"))
		},
	}
	for _, route := range routes {
		route.toMatcher(router)
	}
	return router
}

func (r *Router) Add(route *Route) *Router {
	if route != nil {
		route.toMatcher(r)
	}
	return r
}

func (r *Router) NotFound(handler func (types.RequestScope)) *Router {
	r.notFoundFunc = handler
	return r
}

func (r *Router) addMatcher(method string, matcher *Matcher) {
	matchers := r.getMatchersByMethod(method)
	matchers.Add(matcher)
}

func (r *Router) getMatchersByMethod(method string) *Matchers {
	switch method {
	case methodGet   : return &r.gets
	case methodPost  : return &r.posts
	case methodPut   : return &r.puts
	case methodDelete: return &r.deletes
	default: return nil
	}
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if len(req.URL.Scheme) == 0 {
		req.URL.Scheme = defaultScheme
	}
	if len(req.URL.Host) == 0 {
		req.URL.Host = req.Host
	}
	urlParts := strings.Split(req.URL.String(), "?")
	urlStr := urlParts[0]
	var scope types.RequestScope
	var match bool
	var existsMatcher *Matcher
	matchers := r.getMatchersByMethod(req.Method)
	for _, matcher := range *matchers {
		match, scope = matcher.Match(urlStr, req, rw)
		if match {
			existsMatcher = matcher
			break
		}
	}
	if existsMatcher != nil {
		hasConstruct := existsMatcher.construct != nil
		hasHandler := existsMatcher.handler != nil
		if hasConstruct && hasHandler {
			r.doMiddlewares(existsMatcher.beforeMiddlewares, scope)
			controller := existsMatcher.construct()
			controller.CallAction(existsMatcher.handler, scope)
			r.doMiddlewares(existsMatcher.afterMiddlewares, scope)
		} else if !hasConstruct && hasHandler {
			r.doMiddlewares(existsMatcher.beforeMiddlewares, scope)
			existsMatcher.handler.(func(types.RequestScope))(scope)
			r.doMiddlewares(existsMatcher.afterMiddlewares, scope)
		} else {
			r.callNotFoundFunc(rw, req)
		}
	} else {
		r.callNotFoundFunc(rw, req)
	}
}

func (r *Router) doMiddlewares(middlewares []types.MiddlewareFunc, scope types.RequestScope) {
	for _, middleware := range middlewares {
		middleware(scope)
	}
}

func (r *Router) callNotFoundFunc(rw http.ResponseWriter, req *http.Request) {
	scope := NewRequestScope()
	scope.SetRequest(req)
	scope.SetWriter(rw)
	r.notFoundFunc(scope)
}

