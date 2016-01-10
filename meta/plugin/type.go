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
				t.logger.Debug("check %s parent entities", entity.GetFullName())
				parentEntities := make([]types.Entity, 0)
				for _, extend := range entity.GetRawExtends() {
					parentEntity := t.getEntity(extend)
					if parentEntity != nil && parentEntity.GetEntityKind() == entity.GetEntityKind() {
						if parentEntity.GetContainer().GetImport() != container.GetImport() {
							entity.AddImport(parentEntity.GetContainer().GetImport())
						}
						parentEntities = append(parentEntities, parentEntity)
						t.logger.Debug("found %s parent entity", parentEntity.GetFullName())
					} else {
						t.logger.Error("parent entity %s not found", extend)
					}
				}
				entity.SetExtends(parentEntities)
			}
			switch entity.GetEntityKind() {
			case types.ControllerEntityKind:
			case types.ModelEntityKind: t.fillProperties(entity.(types.ModelEntity))
			}
		}
	}
	for _, container := range t.containers {
		if container.GetContainerKind() == types.EntityContainerKind {
			for _, entity := range container.GetEntities() {
				if entity.GetEntityKind() == types.ModelEntityKind {
					modelEntity := entity.(types.ModelEntity)
					if modelEntity.GetPattern() == types.StraightMappingPattern {
						for _, property := range modelEntity.GetProperties() {
							relation := property.GetRelation()
							hasMany := relation.IsOneToMany() || relation.IsManyToMany()
							isModel := property.GetEntity() != nil && property.GetEntity().GetEntityKind() == types.ModelEntityKind
							if hasMany && isModel {
								relModel := property.GetEntity().(types.ModelEntity)
								for _, relProp := range relModel.GetProperties() {
									relPropRelation := relProp.GetRelation()
									hasRelMany := relPropRelation.IsOneToOne() || relPropRelation.IsManyToMany()
									equalKind := entity.GetFullName() == relProp.GetKind()
									if hasRelMany && equalKind {
										property.SetRelationProperty(relProp)
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}
//	for _, container := range t.containers {
//		if container.GetContainerKind() == types.ControllerContainerKind {
//			for _, controller := range container.GetEntities() {
//
//			}
//		}
//	}
}

func (t *Type) getEntity(name string) types.Entity {
	var needle types.Entity = nil
	for _, container := range t.containers {
		for _, entity := range container.GetEntities() {
			if name == entity.GetFullName() {
				needle = entity
				break
			}
		}
	}
	return needle
}

func (t *Type) fillProperties(entity types.ModelEntity) {
	modelEntity := entity.(*model.Model)

	t.initModelProperties(modelEntity)
	properties := entity.GetProperties()

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
			} else {
				property.SetSelfPackage(entity.GetContainer().GetImport() == relEntity.GetContainer().GetImport())
				if !property.IsSelfPackage() {
					entity.AddImport(relEntity.GetContainer().GetImport())
				}
				property.SetEntity(relEntity)
				properties = append(properties, property)

				if modelEntity.GetPattern() == types.StraightMappingPattern {
					relation := property.GetRelation()
					if relation.IsOneToOne() {
						var kind string
						if relEntity.GetEntityKind() == types.EnumEntityKind {
							enumEntity := relEntity.(*model.Enum)
							kind = string(enumEntity.GetKind())
						} else if relEntity.GetEntityKind() == types.ModelEntityKind {
							relModelEntity := relEntity.(types.ModelEntity)
							t.initModelProperties(relModelEntity)
							identifier := relModelEntity.GetProperties()[0]
							kind = identifier.GetKind()
						}

						prop := &model.Property{
							Name:     fmt.Sprintf("%sId", property.GetName()),
							Kind:     kind,
							Size:     0,
							Required: property.IsRequired(),
							Relation: types.NoneRelation,
						}
						prop.SetUpperName(utils.UpperFirst(prop.GetName()))
						properties = append(properties, prop)
					}
				}
			}
		} else {
			if property.GetKind() == model.TimeKind {
				entity.AddImport(model.TimeImport)
			}
			properties = append(properties, property)
		}
	}

	modelEntity.SetProperties(properties)
}

func (t *Type) initModelProperties(entity types.ModelEntity) {
	modelEntity := entity.(*model.Model)
	isStraightMapping := entity.GetPattern() == types.StraightMappingPattern
	hasntProperties := entity.GetProperties() == nil

	if hasntProperties {
		entity.SetProperties(make([]types.Property, 0))
	}

	if modelEntity.Identifier == nil && isStraightMapping {
		t.logger.Error("model %s should be have identifier", entity.GetName())
	}

	if hasntProperties && isStraightMapping {
		identifier := modelEntity.Identifier
		if len(identifier.Kind) == 0 {
			identifier.Kind = model.DefaultIdentifierKind
		}
		prop := &model.Property{
			Name:     model.DefaultIdentifierName,
			Kind:     identifier.Kind,
			Size:     0,
			Required: true,
			Relation: types.NoneRelation,
		}
		prop.SetUpperName(utils.UpperFirst(prop.GetName()))

		entity.SetProperties(append(entity.GetProperties(), prop))
	}
}
