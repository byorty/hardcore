package plugin

import (
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
)

type RouterImpl struct{}

func NewRouter() types.ApplicationPlugin {
	return new(RouterImpl)
}

func (r *RouterImpl) Run() {
	router := mux.NewRouter()
	for _, route := range scope.App().GetRoutes() {
		router.Add(route)
	}
	scope.App().SetRouter(router)
}
