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
)

type Config struct {
    logger log.Logger
    containers []types.Container
}

func (c *Config) Do(env types.Environment) {
    c.logger = env.GetLogger()
    config := env.GetConfiguration().(*conf.Configuration)
    c.containers = make([]types.Container, 0)
    if config.ControllerContainers != nil {
        for _, container := range config.ControllerContainers {
            c.containers = append(c.containers, c.initControllerEntities(container))
        }
    }
    if config.ModelContainers != nil {
        for _, container := range config.ModelContainers {
            c.containers = append(c.containers, c.initModelEntities(container))
        }
    }

    for _, include := range config.Includes {
        c.logger.Debug("find include %s", include.File)
        incFilename := filepath.Join(env.GetMetaPath(), include.File)
        if utils.FileExists(incFilename) {
            c.logger.Debug("include file %s is exists", include.File)
            data, err := ioutil.ReadFile(incFilename)
            if err == nil {
                c.logger.Debug("success read include file %s", include.File)
                var incConfig conf.Configuration
                err = xml.Unmarshal(data, &incConfig)
                if err == nil {
                    c.logger.Debug("success unmarshal include file %s", include.File)
                    c.mergeModelContainer(incConfig.ModelContainers)
                    c.mergeControllerContainer(incConfig.ControllerContainers)
                } else {
                    c.logger.Error("can't unmarshal include file %s", include.File)
                }
            } else {
                c.logger.Error("can't read include file %s", include.File)
            }
        } else {
            c.logger.Error("include file %s not found", include.File)
        }
    }

    config.ControllerContainers = nil
    config.ModelContainers = nil
    config.Includes = nil
    config.SetContainers(c.containers)
}

func (c *Config) initModelEntities(container *model.Container) types.Container {
    entities := make([]types.Entity, 0)
    for _, entity := range container.Models {
        entities = append(entities, entity)
    }
    for _, entity := range container.Enums {
        entities = append(entities, entity)
    }
    container.Models = nil
    container.Enums = nil
    container.SetEntities(entities)
    return container
}

func (c *Config) initControllerEntities(container *controller.Container) types.Container {
    entities := make([]types.Entity, 0)
    for _, entity := range container.Controllers {
        entities = append(entities, entity)
    }
    container.Controllers = nil
    container.SetEntities(entities)
    return container
}

func (c *Config) mergeModelContainer(containers []*model.Container)  {
    for _, container := range containers {
        c.mergeContainer(c.initModelEntities(container))
    }
}

func (c *Config) mergeContainer(container types.Container) {
    c.logger.Debug("check container %s", container.GetPackage())
    containersLen := len(c.containers)
    x := sort.Search(containersLen, func(j int) bool {
        return c.containers[j].GetPackage() == container.GetPackage()
    })
    if x < containersLen && c.containers[x].GetPackage() == container.GetPackage() {
        c.logger.Error("container %s is duplicate", container.GetPackage())
    } else {
        c.logger.Debug("append container %s in config", container.GetPackage())
        c.containers = append(c.containers, container)
    }
}

func (c *Config) mergeControllerContainer(containers []*controller.Container)  {
    for _, container := range containers {
        c.mergeContainer(c.initControllerEntities(container))
    }
}
