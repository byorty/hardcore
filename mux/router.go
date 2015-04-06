package mux

import "net/http"

type Router struct {
	routes          []*Route
	notFoundHandler http.Handler
}

func NewRouter(routes ...*Route) *Router {
	return &Router{
		routes  : routes,
	}
}

func (r *Router) prepare() {
//	for _, route := range r.routes {

//	}
}

func (r *Router) Add(route *Route) *Router {
	r.routes = append(r.routes, route)
	return r
}

func (r *Router) NotFound(handler http.Handler) *Router {
	r.notFoundHandler = handler
	return r
}
