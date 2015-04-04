package hardcore

import "net/http"

type Router struct {
	// Routes to be matched, in order.
	routes []*Route
	// Parent route, if this is subrouter
	parent *Router
	path string
	// Configurable Handler to be used when no route matches
	notFoundHandler http.Handler
}

func (r *Router) Scheme(scheme string) *Router {

	return r
}

func (r *Router) Host(host string) *Router {

	return r
}

func (r *Router) Header(key, value string) *Router {

	return r
}

func (r *Router) Before() *Router {

	return r
}

func (r *Router) Get(path string, handler interface{}) *Router {

	return r
}

func (r *Router) Post(path string, handler interface{}) *Router {

	return r
}

func (r *Router) Put(path string, handler interface{})*Router {

	return r
}

func (r *Router) Delete(path string, handler interface{}) *Router {

	return r
}

func (r *Router) After() *Router {

	return r
}

func (r *Router) NotFound(handler http.Handler) *Router {
	r.notFoundHandler = handler
	return r
}

func (r *Router) Sub(path string, router *Router) *Router {

	return r
}
