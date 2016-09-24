package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
)

type WebsocketImpl struct {}

func NewWebsocket() types.ApplicationPlugin {
	return new(WebsocketImpl)
}

func (w *WebsocketImpl) Run() {
	scope.App().SetEnableWebsocket(true)
	app := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", scope.App().GetHostname(), scope.App().GetPort()),
		Handler:      scope.App().GetRouter(),
		ReadTimeout:  scope.App().GetReadTimeout(),
		WriteTimeout: scope.App().GetWriteTimeout(),
	}
	scope.App().GetLogger().Info("websocket server - run on %s:%d", scope.App().GetHostname(), scope.App().GetPort())
	go func() {
		scope.App().GetLogger().Error("websocket server - %v", app.ListenAndServe())
	}()
}
