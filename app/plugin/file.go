package plugin

import (
	"fmt"
	"github.com/byorty/hardcore/mux"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
	"net/http"
)

type StaticFileImpl struct{}

func (f *StaticFileImpl) Run() {
	if scope.App().GetStaticPath() != "" && scope.App().GetStaticDir() != "" {
		scope.App().SetRoutes(
			append(
				scope.App().GetRoutes(),
				mux.Get(
					fmt.Sprintf("%s/:dir/:filename", scope.App().GetStaticPath()),
					func(rs types.RequestScope) {
						http.ServeFile(
							rs.GetWriter(),
							rs.GetRequest(),
							fmt.Sprintf(
								"%s/%s/%s",
								scope.App().GetStaticDir(),
								rs.GetPathParams().GetString("dir"),
								rs.GetPathParams().GetString("filename"),
							),
						)
					},
				),
			),
		)
	}
}
