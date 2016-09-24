package mux

import (
	"github.com/byorty/hardcore/helper"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
)

type Router struct {
	matchers     map[string]Matchers
	notFoundFunc types.MiddlewareFunc
	routes       []*Route
}

func NewRouter(routes ...*Route) *Router {
	router := &Router{
		matchers: make(map[string]Matchers),
		notFoundFunc: func(rs types.RequestScope) {
			rs.GetWriter().WriteHeader(http.StatusNotFound)
			rs.GetWriter().Write([]byte("not found"))
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

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if len(req.URL.Scheme) == 0 {
		req.URL.Scheme = defaultScheme
	}
	if len(req.URL.Host) == 0 {
		req.URL.Host = req.Host
	}
	var matchersKey string
	if scope.App().GetEnableWebsocket() {
		if req.Header.Get("Upgrade") == "websocket" && req.Header.Get("Connection") == "Upgrade" {
			matchersKey = websocket
		} else {
			matchersKey = req.Method
		}
	} else {
		matchersKey = req.Method
	}
	var rs types.RequestScope
	var existsMatcher *Matcher
	for _, matcher := range r.matchers[matchersKey] {
		if ok, newScope := matcher.Match(req.URL.Path, req, rw); ok {
			existsMatcher = matcher
			rs = newScope
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
