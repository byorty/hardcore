package plugin

import (
	"encoding/xml"
	"github.com/byorty/hardcore/log"
	"github.com/byorty/hardcore/meta/conf"
	"github.com/byorty/hardcore/meta/types"
	"github.com/byorty/hardcore/utils"
	"io/ioutil"
	"path/filepath"
)

type Init struct {
	logger     log.Logger
	containers []types.Container
}

func (i *Init) Do(env types.Environment) {
	i.logger = env.GetLogger()
	config := env.GetConfiguration().(*conf.Configuration)
	config.Init()

	containers := config.GetContainers()
	i.init(env, config.GetContainers())

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
					incConfig.Init()
					i.merge(env, incConfig.GetContainers(), &containers)
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
	config.SetContainers(containers)
}

func (i *Init) init(env types.Environment, containers []types.Container) {
	for _, container := range containers {
		container.Init(env)
		for y := 0; y < container.Len(); y++ {
			entity := container.Get(y)
			entity.Init(container)
		}
	}
}

func (i *Init) merge(env types.Environment, srcContainers []types.Container, ptrDestContainers *[]types.Container) {
	destContainers := *ptrDestContainers
	i.init(env, srcContainers)
	for _, srcContainer := range srcContainers {
		isExists := false
		for _, destContainer := range destContainers {
			if destContainer.GetPackage() == srcContainer.GetPackage() {
				isExists = true
				break
			}
		}
		if !isExists {
			(*ptrDestContainers) = append((*ptrDestContainers), srcContainer)
		}
	}
}
