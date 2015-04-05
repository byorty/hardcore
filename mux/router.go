package mux

type router struct {
	routes          []*Route
}

func Router(routes ...*Route) *router {
	return &router{routes: routes}
}

func (r *router) Add(route *Route) *router {
	r.routes = append(r.routes, route)
	return r
}
