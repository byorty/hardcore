package plugin

import (
	"github.com/byorty/hardcore/meta/types"
)

var (
	middlewareTpl = `package {{.Package}}

import ({{range .Imports}}
	"{{.}}"{{end}}
)

func {{.Name}}(rs types.RequestScope) {

}
`
)

type Middleware struct{}

func (m *Middleware) Do(env types.Environment) {
	config := env.GetConfiguration()
	for _, container := range config.GetContainers() {
		if container.GetContainerKind() == types.MiddlewaresContainerKind {
			for _, entity := range container.GetEntities() {
				if entity.GetEntityKind() == types.MiddlewareEntityKind {

					tplParams := map[string]interface{}{
						"Name":    entity.GetName(),
						"Package": container.GetShortPackage(),
						"Imports": []string{
							types.DefaultImport,
						},
					}

					env.GetConfiguration().AddFile(entity.GetFilename(), middlewareTpl, tplParams)
				}
			}
		}
	}
}
