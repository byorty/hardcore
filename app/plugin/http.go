package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/scope"
	"net/http"
	"github.com/byorty/hardcore/types"
)

func NewHttp() types.ApplicationPlugin {
	return NewHttpByName(scope.DefaultName)
}

func NewHttpByName(name string) types.ApplicationPlugin {
	return newByName(new(HttpImpl), name)
}

type HttpImpl struct{
	BaseImpl
}

func (h *HttpImpl) Run() {
	app := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", scope.AppByName(h.name).GetHostname(), scope.AppByName(h.name).GetPort()),
		Handler:      scope.AppByName(h.name).GetRouter(),
		ReadTimeout:  scope.AppByName(h.name).GetReadTimeout(),
		WriteTimeout: scope.AppByName(h.name).GetWriteTimeout(),
	}
	scope.AppByName(h.name).GetLogger().Info("http server - run on %s:%d", scope.AppByName(h.name).GetHostname(), scope.AppByName(h.name).GetPort())
	go func() {
		scope.AppByName(h.name).GetLogger().Error("server - %v", app.ListenAndServe())
	}()
}
