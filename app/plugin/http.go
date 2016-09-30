package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
)

type HttpImpl struct{}

func NewHttp() types.ApplicationPlugin {
	return new(HttpImpl)
}

func (h *HttpImpl) Run() {
	app := h.createServer(scope.App().GetHostname(), scope.App().GetPort())
	go func () {
		err := app.ListenAndServe()
		if err == nil {
			scope.App().GetLogger().Finest("http server - run on %s:%d", scope.App().GetHostname(), scope.App().GetPort())
		} else {
			scope.App().GetLogger().Error("http server - %v", err)
			scope.App().Exit()
		}
	}()
	if scope.App().GetCertFilename() != "" && scope.App().GetPrivateKeyFilename() != "" {
		app := h.createServer(scope.App().GetHostname(), scope.App().GetSecurityPort())
		go func () {
			err := app.ListenAndServeTLS(scope.App().GetCertFilename(), scope.App().GetPrivateKeyFilename())
			if err == nil {
				scope.App().GetLogger().Finest("https server - run on %s:%d", scope.App().GetHostname(), scope.App().GetPort())
			} else {
				scope.App().GetLogger().Error("https server - %v", err)
				scope.App().Exit()
			}
		}()
	}
}

func (h *HttpImpl) createServer(hostname string, port int) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", hostname, port),
		Handler:      scope.App().GetRouter(),
		ReadTimeout:  scope.App().GetReadTimeout(),
		WriteTimeout: scope.App().GetWriteTimeout(),
	}
}
