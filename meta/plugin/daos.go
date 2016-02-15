package plugin

import (
	"github.com/byorty/hardcore/meta/types"
	"path/filepath"
)

var (
	daosTpl = `package configs

import ({{range .Imports}}
	"{{.}}"{{end}}
)

var (
	DAOs = []types.ModelDAO{ {{range .Entities}}
		{{.GetContainer.GetShortPackage}}.{{.GetName}}DaoInst(),{{end}}
	}
)
`
)

type DAOs struct {}

func (d *DAOs) Do(env types.Environment) {
	imports := make([]string, 0)
	entities := make([]types.Entity, 0)
	for _, container := range env.GetConfiguration().GetContainers() {
		if container.GetContainerKind() == types.EntityContainerKind {
			hasImport := false
			for _, existsImport := range imports {
				if existsImport == container.GetImport() {
					hasImport = true
					break
				}
			}
			if !hasImport {
				imports = append(imports, container.GetImport())
			}
			for _, entity := range container.GetEntities() {
				if entity.GetEntityKind() == types.ModelEntityKind {
					entities = append(entities, entity)
				}
			}
		}
	}
	tplParams := map[string]interface{}{
		"Imports": append([]string{
			types.DefaultImport,
		}, imports...),
		"Entities": entities,
	}

	filename := filepath.Join(env.GetAbsPath(), "configs", "daos_auto")
	env.GetConfiguration().AddAutoFile(filename, daosTpl, tplParams)
}
