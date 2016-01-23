package plugin

import (
    "github.com/byorty/hardcore/meta/types"
    "github.com/byorty/hardcore/meta/conf"
    "encoding/xml"
    "io/ioutil"
    "github.com/byorty/hardcore/utils"
    "path/filepath"
    "github.com/byorty/hardcore/meta/controller"
    "github.com/byorty/hardcore/meta/model"
    "sort"
    "github.com/byorty/hardcore/log"
	"github.com/byorty/hardcore/meta/exporter"
)

type Include struct {
    logger log.Logger
    containers []types.Container
}

func (i *Include) Do(env types.Environment) {
    i.logger = env.GetLogger()
    config := env.GetConfiguration().(*conf.Configuration)
    i.containers = make([]types.Container, 0)
    if config.ControllerContainers != nil {
        for _, container := range config.ControllerContainers {
            i.containers = append(i.containers, i.initControllerEntities(container))
        }
    }
    if config.ModelContainers != nil {
        for _, container := range config.ModelContainers {
            i.containers = append(i.containers, i.initModelEntities(container))
        }
    }
    if config.ExporterContainers != nil {
        for _, container := range config.ExporterContainers {
            i.containers = append(i.containers, i.initExporterEntities(container))
        }
    }

    for _, include := range config.Includes {
        i.logger.Debug("find include %s", include.File)
        incFilename := filepath.Join(env.GetMetaPath(), include.File)
        if utils.FileExists(incFilename) {
            i.logger.Debug("include file %s is exists", include.File)
            data, err := ioutil.ReadFile(incFilename)
            if err == nil {
                i.logger.Debug("success read include file %s", include.File)
                var incConfig conf.Configuration
                err = xml.Unmarshal(data, &incConfig)
                if err == nil {
                    i.logger.Debug("success unmarshal include file %s", include.File)
                    i.mergeModelContainer(incConfig.ModelContainers)
                    i.mergeControllerContainer(incConfig.ControllerContainers)
                    i.mergeExporterContainer(incConfig.ExporterContainers)
                } else {
                    i.logger.Error("can't unmarshal include file %s", include.File)
                }
            } else {
                i.logger.Error("can't read include file %s", include.File)
            }
        } else {
            i.logger.Error("include file %s not found", include.File)
        }
    }

    config.ControllerContainers = nil
    config.ModelContainers = nil
    config.ExporterContainers = nil
    config.Includes = nil
    config.SetContainers(i.containers)
}

func (i *Include) initModelEntities(container *model.Container) types.Container {
    entities := make([]types.Entity, 0)
    for _, entity := range container.Models {
        entities = append(entities, i.initEntity(container, entity))
    }
    for _, entity := range container.Enums {
        entities = append(entities, i.initEntity(container, entity))
    }
    container.Models = nil
    container.Enums = nil
    container.SetEntities(entities)
    return container
}

func (i *Include) initEntity(container types.Container, entity types.Entity) types.Entity {
    entity.ClearName()
    entity.SetContainer(container)

	if entity.GetEntityKind() == types.ControllerEntityKind {
		ctrl := entity.(*controller.Controller)
		actions := make([]types.Action, 0)

		for _, action := range ctrl.Actions {
			params := make([]types.ActionParam, 0)
			for _, param := range action.Params {
				params = append(params, param)
			}
            if len(params) == 0 {
                param := &controller.Param{
                    Name: "scope",
                    Required: true,
                    Source: "",
                    Kind: controller.RequestScopeKind,
                }
                params = append(params, param)
            }
			action.SetParams(params)
			actions = append(actions, action)
			action.Params = nil
		}
		ctrl.SetActions(actions)
		ctrl.Actions = nil
	} else if entity.GetEntityKind() == types.ExporterEntityKind {
		exp := entity.(*exporter.Exporter)
		props := make([]types.ExportableProperty, 0)
		for _, prop := range exp.Properties {
			props = append(props, prop)
		}
		exp.SetProperties(props)
		exp.Properties = nil
	}

    return entity
}

func (i *Include) initControllerEntities(container *controller.Container) types.Container {
    entities := make([]types.Entity, 0)
    for _, entity := range container.Controllers {
        entities = append(entities, i.initEntity(container, entity))
    }
    container.Controllers = nil
    container.SetEntities(entities)
    return container
}

func (i *Include) mergeModelContainer(containers []*model.Container)  {
    for _, container := range containers {
        i.mergeContainer(i.initModelEntities(container))
    }
}

func (i *Include) mergeContainer(container types.Container) {
    i.logger.Debug("check container %s", container.GetPackage())
    containersLen := len(i.containers)
    x := sort.Search(containersLen, func(j int) bool {
        return i.containers[j].GetPackage() == container.GetPackage()
    })
    if x < containersLen && i.containers[x].GetPackage() == container.GetPackage() {
        i.logger.Error("container %s is duplicate", container.GetPackage())
    } else {
        i.logger.Debug("append container %s in config", container.GetPackage())
        i.containers = append(i.containers, container)
    }
}

func (i *Include) mergeControllerContainer(containers []*controller.Container)  {
    for _, container := range containers {
        i.mergeContainer(i.initControllerEntities(container))
    }
}

func (i *Include) initExporterEntities(container *exporter.Container) types.Container {
	entities := make([]types.Entity, 0)
	for _, entity := range container.Exporters {
		entities = append(entities, i.initEntity(container, entity))
	}
	container.Exporters = nil
	container.SetEntities(entities)
	return container
}

func (i *Include) mergeExporterContainer(containers []*exporter.Container)  {
	for _, container := range containers {
		i.mergeContainer(i.initExporterEntities(container))
	}
}
