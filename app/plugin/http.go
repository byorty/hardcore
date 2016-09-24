package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/scope"
	"net/http"
	"github.com/byorty/hardcore/types"
)

type HttpImpl struct{}

func NewHttp() types.ApplicationPlugin {
	return new(HttpImpl)
}

func (h *HttpImpl) Run() {
	app := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", scope.App().GetHostname(), scope.App().GetPort()),
		Handler:      scope.App().GetRouter(),
		ReadTimeout:  scope.App().GetReadTimeout(),
		WriteTimeout: scope.App().GetWriteTimeout(),
	}
	scope.App().GetLogger().Info("server - run on %s:%d", scope.App().GetHostname(), scope.App().GetPort())
	scope.App().GetLogger().Error("server - %v", app.ListenAndServe())
}
