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
	scope.App().GetLogger().Finest("router - initialize %T", router)
	for _, route := range scope.App().GetRoutes() {
		scope.App().GetLogger().Finest("router - add router %v", route)
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
	return new(WebsocketRouterImpl)
}

func (w *WebsocketRouterImpl) Run() {
	w.run(mux.NewWebsocketRouter())
}
