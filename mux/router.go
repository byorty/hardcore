package mux

import (
	"github.com/byorty/hardcore/helper"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"github.com/gorilla/websocket"
	"net/http"
)

type Router struct {
	matchers     map[string]Matchers
	notFoundFunc types.MiddlewareFunc
	routes       []*Route
}

func newRouter() *Router {
	router := &Router{
		matchers: make(map[string]Matchers),
		notFoundFunc: func(rs types.RequestScope) {
			rs.GetWriter().WriteHeader(http.StatusNotFound)
			rs.GetWriter().Write([]byte("not found"))
		},
	}
	return router
}

func (r *Router) Batch(routes ...types.Route) types.Router {
	for _, route := range routes {
		r.Add(route)
	}
	return r
}

func (r *Router) Add(route types.Route) types.Router {
	route.(*Route).toMatcher(r)
	return r
}

func (r *Router) NotFound(handler func(types.RequestScope)) types.Router {
	r.notFoundFunc = handler
	return r
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {}

func (r *Router) find(method, path string) (*Matcher, types.RequestScope) {
	scope.App().GetLogger().Debug("mux: request method is %s", method)
	var rs types.RequestScope
	var existsMatcher *Matcher
	for _, matcher := range r.matchers[method] {
		if ok, newScope := matcher.Match(path); ok {
			existsMatcher = matcher
			rs = newScope
			break
		}
	}
	return existsMatcher, rs
}

func (r *Router) serve(matcher *Matcher, rs types.RequestScope, rw http.ResponseWriter, req *http.Request) {
	if matcher != nil && rs != nil {
		rs.SetRequest(req)
		rs.SetWriter(rw)
		scope.App().GetLogger().Debug("mux: router found matcher")
		hasConstruct := matcher.construct != nil
		hasHandler := matcher.handler != nil
		if hasConstruct && hasHandler {
			r.fetchSession(rs)
			r.doMiddlewares(matcher.beforeMiddlewares, rs)
			if rs.NotPrevented() {
				controller := matcher.construct()
				scope.App().GetLogger().Error("mux: router call method %T of controller %T", matcher.handler, controller)
				controller.CallAction(matcher.handler, rs)
			}
			r.doMiddlewares(matcher.afterMiddlewares, rs)
		} else if !hasConstruct && hasHandler {
			r.fetchSession(rs)
			r.doMiddlewares(matcher.beforeMiddlewares, rs)
			if rs.NotPrevented() {
				scope.App().GetLogger().Error("mux: router call function %T", matcher.handler)
				matcher.handler.(func(types.RequestScope))(rs)
			}
			r.doMiddlewares(matcher.afterMiddlewares, rs)
		} else {
			scope.App().GetLogger().Error("mux: router can`t to call matcher")
			r.callNotFoundFunc(rw, req)
		}
	} else {
		scope.App().GetLogger().Error("mux: router can`t to find matcher")
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

type HttpRouter struct {
	*Router
}

func NewHttpRouter() types.Router {
	return newHttpRouter()
}

func newHttpRouter() *HttpRouter {
	router := new(HttpRouter)
	router.Router = newRouter()
	return router
}

func (h *HttpRouter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	matcher, rs := h.find(req.Method, req.URL.Path)
	h.serve(matcher, rs, rw, req)
}

type WebsocketRouter struct {
	*HttpRouter
	upgrader *websocket.Upgrader
}

func NewWebsocketRouter() types.Router {
	router := new(WebsocketRouter)
	router.HttpRouter = newHttpRouter()
	if scope.App().GetUpgrader() != nil {
		router.upgrader = scope.App().GetUpgrader()
	} else {
		router.upgrader = new(websocket.Upgrader)
	}
	return router
}

func (w *WebsocketRouter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if scope.App().GetEnableWebsocket() {
		upgrade := req.Header.Get("Upgrade")
		connection := req.Header.Get("Connection")
		scope.App().GetLogger().Debug("mux: upgrade header contain %s", upgrade)
		scope.App().GetLogger().Debug("mux: connection header contain - %s", connection)
	}
	scope.App().GetLogger().Debug("websocket.IsWebSocketUpgrade %v", websocket.IsWebSocketUpgrade(req))
	if websocket.IsWebSocketUpgrade(req) {
		conn, err := w.upgrader.Upgrade(rw, req, nil)
		if err == nil {
			matcher, rs := w.find(methodWebsocket, req.URL.Path)
			scope.App().GetLogger().Debug("matcher - %v, rs - %v", matcher, rs)
			ws := scope.NewWebsocket(rs, conn)
			w.serve(matcher, ws, rw, req)
			defer conn.Close()
		} else {
			scope.App().GetLogger().Debug("mux: can`t upgrade connection^ err - %v", err)
		}
	} else {
		w.HttpRouter.ServeHTTP(rw, req)
	}
}
