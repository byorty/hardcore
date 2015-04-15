package mux

import "net/http"

type Router struct {
	gets         Matchers
	posts        Matchers
	puts         Matchers
	deletes      Matchers
	notFoundFunc MiddlewareFunc
}

func NewRouter(routes ...*Route) *Router {
	router := &Router{
		gets   : make(Matchers, 0),
		posts  : make(Matchers, 0),
		puts   : make(Matchers, 0),
		deletes: make(Matchers, 0),
	}
	for _, route := range routes {
		route.toMatcher(router)
	}
	return router
}

func (r *Router) Add(route *Route) *Router {
	route.toMatcher(r)
	return r
}

func (r *Router) NotFound(handler http.Handler) *Router {
	r.notFoundHandler = handler
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
	scope := &RequestScope{
		UrlStr: req.URL.String(),
		Url: req.URL,
		Writer: rw,
		Request: req,
	}
	var existsMatcher *Matcher
	matchers := r.getMatchersByMethod(req.Method)
	for _, matcher := range *matchers {
		if matcher.Match(scope) {
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
			existsMatcher.handler.(func(*RequestScope))(scope)
			r.doMiddlewares(existsMatcher.afterMiddlewares, scope)
		} else {
			r.notFoundFunc(scope)
		}
	} else {
		r.notFoundFunc(scope)
	}
}

func (r *Router) doMiddlewares(middlewares []MiddlewareFunc, scope *RequestScope) {
	for _, middleware := range middlewares {
		middleware(scope)
	}
}

