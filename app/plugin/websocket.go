package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
	"golang.org/x/net/websocket"
)

func NewWebsocket() types.ApplicationPlugin {
	return NewWebsocketByName(scope.DefaultName)
}

func NewWebsocketByName(name string) types.ApplicationPlugin {
	return newByName(new(WebsocketImpl), name)
}

type WebsocketImpl struct{
	BaseImpl
}

func (w *WebsocketImpl) Run() {
	app := &websocket.Server{
		Handler: scope.AppByName(w.name).GetRouter(),
	}
	scope.AppByName(w.name).GetLogger().Info("websocket server - run on %s:%d", scope.AppByName(w.name).GetHostname(), scope.AppByName(w.name).GetPort())
	go func() {
		scope.AppByName(w.name).
			GetLogger().
			Error(
				"server - %v",
				http.ListenAndServe(fmt.Sprintf("%s:%d", scope.AppByName(w.name).GetHostname(), scope.AppByName(w.name).GetPort()), app),
			)
	}()
}
