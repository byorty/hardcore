package plugin

import (
	"github.com/byorty/hardcore/meta/types"
)

type Extend struct{}

func (e *Extend) Do(env types.Environment) {
	containers := env.GetConfiguration().GetContainers()
	logger := env.GetLogger()
	for _, container := range containers {
		for _, entity := range container.GetEntities() {
			if entity.GetRawExtends() != nil {
				logger.Debug("check %s parent entities", entity.GetFullName())
				parentEntities := make([]types.Entity, 0)
				for _, extend := range entity.GetRawExtends() {
					parentEntity := env.GetConfiguration().GetEntity(extend)
					if parentEntity != nil && parentEntity.GetEntityKind() == entity.GetEntityKind() {
						if parentEntity.GetContainer().GetImport() != container.GetImport() {
							entity.AddImport(parentEntity.GetContainer().GetImport())
						}
						parentEntities = append(parentEntities, parentEntity)
						logger.Debug("found %s parent entity", parentEntity.GetFullName())
					} else {
						logger.Error("parent entity %s not found", extend)
					}
				}
				entity.SetExtends(parentEntities)
			}
		}
	}
}
