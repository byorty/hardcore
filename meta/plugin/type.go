package plugin

import (
	"github.com/byorty/hardcore/log"
	"github.com/byorty/hardcore/meta/model"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"fmt"
)

type Type struct {
	logger     log.Logger
	containers []types.Container
}

func (t *Type) Do(env types.Environment) {
	t.containers = env.GetConfiguration().GetContainers()
	t.logger = env.GetLogger()
	for _, container := range t.containers {
		for _, entity := range container.GetEntities() {
			entity.SetImports(make([]string, 0))
			if entity.GetRawExtends() != nil {
				t.logger.Debug("check %s parent entities", entity.GetFullname())
				parentEntities := make([]types.Entity, 0)
				for _, extend := range entity.GetRawExtends() {
					parentEntity := t.getEntity(extend)
					if parentEntity != nil && parentEntity.GetEntityKind() == entity.GetEntityKind() {
						if parentEntity.GetContainer().GetImport() != container.GetImport() {
							entity.AddImport(parentEntity.GetContainer().GetImport())
						}
						parentEntities = append(parentEntities, parentEntity)
						t.logger.Debug("found %s parent entity", parentEntity.GetFullname())
					} else {
						t.logger.Error("parent entity %s not found", extend)
					}
				}
				entity.SetExtends(parentEntities)
			}
			switch entity.GetEntityKind() {
			case types.ControllerEntityKind:
			case types.ModelEntityKind:
				t.fillProperties(entity.(types.ModelEntity))
				//            case types.EnumEntityKind:
			}
		}
	}
}

func (t *Type) getEntity(name string) types.Entity {
	var needle types.Entity = nil
	for _, container := range t.containers {
		for _, entity := range container.GetEntities() {
			if name == entity.GetFullname() {
				needle = entity
				break
			}
		}
	}
	return needle
}

func (t *Type) fillProperties(entity types.ModelEntity) {
	modelEntity := entity.(*model.Model)
	properties := make([]types.Property, 0)

	if modelEntity.Identifier != nil {
		identifier := modelEntity.Identifier
		if len(identifier.Kind) == 0 {
			identifier.Kind = "int"
		}
		prop := &model.Property{
			Name:     "id",
			Kind:     identifier.Kind,
			Size:     0,
			Required: true,
			Relation: types.NoneRelation,
		}
		prop.SetUpperName(utils.UpperFirst(prop.GetName()))

		properties = append(properties, prop)
	}

	for _, property := range modelEntity.Properties {
		property.SetUpperName(utils.UpperFirst(property.GetName()))

		if property.HasRelation() {
			relEntity := t.getEntity(property.GetKind())
			if relEntity == nil {
				t.logger.Error(
					"invalid type %s in property %s entity %s",
					property.GetKind(),
					property.GetName(),
					entity.GetName(),
				)
			}

			relation := property.GetRelation()
			if relation.IsOneToOne() {
				property.SetEntity(relEntity)

				var kind string
				if relEntity.GetEntityKind() == types.EnumEntityKind {
					enumEntity := relEntity.(*model.Enum)
					kind = enumEntity.GetKind()
				} else {

				}
				prop := &model.Property{
					fmt.Sprintf("%sId", property.GetName()),
					kind,
					0,
					property.IsRequired(),
					types.NoneRelation,
				}
				prop.SetUpperName(utils.UpperFirst(prop.GetName()))
			} else if relation.IsOneToMany() {
			} else if relation.IsManyToMany() {
			}
		} else {
			properties = append(properties, property)
		}
	}

	modelEntity.SetProperties(properties)
}
