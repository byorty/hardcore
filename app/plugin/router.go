package plugin

import (
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
)

func NewRouter() types.ApplicationPlugin {
	return NewRouterByName(scope.DefaultName)
}

func NewRouterByName(name string) types.ApplicationPlugin {
	return newByName(new(RouterImpl), name)
}

type RouterImpl struct{
	BaseImpl
}

func (r *RouterImpl) Run() {
	router := mux.NewRouter()
	for _, route := range scope.AppByName(r.name).GetRoutes() {
		router.Add(route)
	}
	scope.AppByName(r.name).SetRouter(router)
}
