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

func NewStaticFile() types.ApplicationPlugin {
	return NewStaticFileByName(scope.DefaultName)
}

func NewStaticFileByName(name string) types.ApplicationPlugin {
	return newByName(new(StaticFileImpl), name)
}

type StaticFileImpl struct{
	BaseImpl
}

func (f *StaticFileImpl) Run() {
	if scope.AppByName(f.name).GetStaticPath() != "" && scope.AppByName(f.name).GetStaticDir() != "" {
		scope.AppByName(f.name).SetRoutes(
			append(
				scope.AppByName(f.name).GetRoutes(),
				mux.Get(
					fmt.Sprintf("%s/:filename>", scope.AppByName(f.name).GetStaticPath()),
					func(rs types.RequestScope) {
						filename := filepath.Join(strings.Split(rs.GetPathParams().GetString("filename"), "/")...)
						http.ServeFile(
							rs.GetWriter(),
							rs.GetRequest(),
							filepath.Join(scope.AppByName(f.name).GetStaticDir(), filename),
						)
					},
				),
			),
		)
	}
}
