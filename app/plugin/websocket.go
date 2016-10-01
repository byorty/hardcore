package plugin

import (
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
)

type WebsocketServerImpl struct{
	WebServerImpl
}

func NewWebsocketServer() types.ApplicationPlugin {
	return new(WebsocketServerImpl)
}

func (w *WebsocketServerImpl) Run() {
	scope.App().SetEnableWebsocket(true)
	w.WebServerImpl.Run()
	scope.App().GetLogger().Finest("web server - websocket enable")
}

type SecureWebsocketServerImpl struct{
	SecureWebServerImpl
}

func NewSecureWebsocketServer() types.ApplicationPlugin {
	return new(SecureWebsocketServerImpl)
}

func (s *SecureWebsocketServerImpl) Run() {
	scope.App().SetEnableWebsocket(true)
	s.SecureWebServerImpl.Run()
	scope.App().GetLogger().Finest("secure web server - websocket enable")
}
