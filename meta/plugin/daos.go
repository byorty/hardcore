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

type DAOs struct {
	PackageImporter
}

func (d *DAOs) Do(env types.Environment) {
	entities := make([]types.Entity, 0)
	for _, container := range env.GetConfiguration().GetContainers() {
		if container.GetContainerKind() == types.EntityContainerKind {
			d.addImport(container.GetImport())

			for _, entity := range container.GetEntities() {
				if entity.GetEntityKind() == types.ModelEntityKind {
					modelEntity := entity.(types.ModelEntity)
					if modelEntity.GetPattern() == types.StraightMappingPattern {
						entities = append(entities, entity)
					}
				}
			}
		}
	}

	if len(entities) > 0 {
		tplParams := map[string]interface{}{
			"Imports": append([]string{
				types.DefaultImport,
			}, d.imports...),
			"Entities": entities,
		}

		filename := filepath.Join(env.GetAbsPath(), "configs", "daos_auto")
		env.GetConfiguration().AddAutoFile(filename, daosTpl, tplParams)
	}
}
