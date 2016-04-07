package plugin

import (
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/scope"
)

type RouterImpl struct{}

func (r *RouterImpl) Run() {
	router := mux.NewRouter()
	for _, route := range scope.App().GetRoutes() {
		router.Add(route)
	}
	scope.App().SetRouter(router)
}
