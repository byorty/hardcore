package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
)

type WebServerImpl struct{}

func NewWebServer() types.ApplicationPlugin {
	return new(WebServerImpl)
}

func (w *WebServerImpl) Run() {
	app := w.createServer(scope.App().GetHostname(), scope.App().GetPort())
	go func() {
		scope.App().GetLogger().Finest("web server - start on %s:%d", scope.App().GetHostname(), scope.App().GetPort())
		err := app.ListenAndServe()
		if err != nil {
			scope.App().GetLogger().Error("web server - %v", err)
			scope.App().Exit()
		}
	}()
}

func (w *WebServerImpl) createServer(hostname string, port int) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", hostname, port),
		Handler:      scope.App().GetRouter(),
		ReadTimeout:  scope.App().GetReadTimeout(),
		WriteTimeout: scope.App().GetWriteTimeout(),
	}
}

type SecureWebServerImpl struct {
	WebServerImpl
}

func NewSecureWebServer() types.ApplicationPlugin {
	return new(SecureWebServerImpl)
}

func (s *SecureWebServerImpl) Run() {
	app := s.createServer(scope.App().GetHostname(), scope.App().GetSecurePort())
	app.TLSConfig = scope.App().GetTlsConfig()
	go func() {
		scope.App().GetLogger().Finest("secure web server - start on %s:%d", scope.App().GetHostname(), scope.App().GetSecurePort())
		err := app.ListenAndServeTLS(scope.App().GetCertFilename(), scope.App().GetPrivateKeyFilename())
		if err != nil {
			scope.App().GetLogger().Error("secure web server - %v", err)
			scope.App().Exit()
		}
	}()
}
