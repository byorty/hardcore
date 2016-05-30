package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
	"path/filepath"
	"strings"
)

type StaticFileImpl struct{}

func (f *StaticFileImpl) Run() {
	if scope.App().GetStaticPath() != "" && scope.App().GetStaticDir() != "" {
		scope.App().SetRoutes(
			append(
				scope.App().GetRoutes(),
				mux.Get(
					fmt.Sprintf("%s/:filename>", scope.App().GetStaticPath()),
					func(rs types.RequestScope) {
						filename := filepath.Join(strings.Split(rs.GetPathParams().GetString("filename"), "/")...)
						http.ServeFile(
							rs.GetWriter(),
							rs.GetRequest(),
							filepath.Join(scope.App().GetStaticDir(), filename),
						)
					},
				),
			),
		)
	}
}
