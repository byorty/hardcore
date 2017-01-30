package plugin

import (
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
)

type HttpRouterImpl struct{}

func NewHttpRouter() types.ApplicationPlugin {
	return new(HttpRouterImpl)
}

func (h *HttpRouterImpl) run(router types.Router) {
	for _, route := range scope.App().GetRoutes() {
		router.Add(route)
	}
	scope.App().SetRouter(router)
}

func (h *HttpRouterImpl) Run() {
	h.run(mux.NewHttpRouter())
}

type WebsocketRouterImpl struct {
	HttpRouterImpl
}

func NewWebsocketRouter() types.ApplicationPlugin {
	return new(WebServerImpl)
}

func (w *WebsocketRouterImpl) Run() {
	w.run(mux.NewWebsocketRouter())
}
