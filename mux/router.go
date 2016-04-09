package mux

import (
	"github.com/byorty/hardcore/helper"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
)

type Router struct {
	gets         Matchers
	posts        Matchers
	puts         Matchers
	deletes      Matchers
	notFoundFunc types.MiddlewareFunc
	routes       []*Route
}

func NewRouter(routes ...*Route) *Router {
	router := &Router{
		gets:    make(Matchers, 0),
		posts:   make(Matchers, 0),
		puts:    make(Matchers, 0),
		deletes: make(Matchers, 0),
		notFoundFunc: func(scope types.RequestScope) {
			scope.GetWriter().WriteHeader(http.StatusNotFound)
			scope.GetWriter().Write([]byte("not found"))
		},
	}
	for _, route := range routes {
		route.toMatcher(router)
	}
	return router
}

func (r *Router) Batch(routes ...types.Route) *Router {
	for _, route := range routes {
		r.Add(route)
	}
	return r
}

func (r *Router) Add(route types.Route) *Router {
	route.(*Route).toMatcher(r)
	return r
}

func (r *Router) NotFound(handler func(types.RequestScope)) *Router {
	r.notFoundFunc = handler
	return r
}

func (r *Router) addMatcher(method string, matcher *Matcher) {
	matchers := r.getMatchersByMethod(method)
	matchers.Add(matcher)
}

func (r *Router) getMatchersByMethod(method string) *Matchers {
	switch method {
	case methodGet:
		return &r.gets
	case methodPost:
		return &r.posts
	case methodPut:
		return &r.puts
	case methodDelete:
		return &r.deletes
	default:
		return nil
	}
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if len(req.URL.Scheme) == 0 {
		req.URL.Scheme = defaultScheme
	}
	if len(req.URL.Host) == 0 {
		req.URL.Host = req.Host
	}
	var rs types.RequestScope
	var match bool
	var existsMatcher *Matcher
	matchers := r.getMatchersByMethod(req.Method)
	for _, matcher := range *matchers {
		match, rs = matcher.Match(req.URL.Path, req, rw)
		if match {
			existsMatcher = matcher
			break
		}
	}
	if existsMatcher != nil {
		hasConstruct := existsMatcher.construct != nil
		hasHandler := existsMatcher.handler != nil
		if hasConstruct && hasHandler {
			r.fetchSession(rs)
			r.doMiddlewares(existsMatcher.beforeMiddlewares, rs)
			if rs.NotPrevented() {
				controller := existsMatcher.construct()
				controller.CallAction(existsMatcher.handler, rs)
			}
			r.doMiddlewares(existsMatcher.afterMiddlewares, rs)
		} else if !hasConstruct && hasHandler {
			r.fetchSession(rs)
			r.doMiddlewares(existsMatcher.beforeMiddlewares, rs)
			if rs.NotPrevented() {
				existsMatcher.handler.(func(types.RequestScope))(rs)
			}
			r.doMiddlewares(existsMatcher.afterMiddlewares, rs)
		} else {
			r.callNotFoundFunc(rw, req)
		}
	} else {
		r.callNotFoundFunc(rw, req)
	}
}

func (r *Router) fetchSession(rs types.RequestScope) {
	if scope.App().GetEnableSession() {
		rs.SetSession(helper.SessionManager().Get(rs))
	}
}

func (r *Router) doMiddlewares(middlewares []types.MiddlewareFunc, rs types.RequestScope) {
	if rs.NotPrevented() {
		for _, middleware := range middlewares {
			if rs.NotPrevented() {
				middleware(rs)
			}
		}
	}
}

func (r *Router) callNotFoundFunc(rw http.ResponseWriter, req *http.Request) {
	requestScope := scope.NewRequest()
	requestScope.SetRequest(req)
	requestScope.SetWriter(rw)
	r.notFoundFunc(requestScope)
}
