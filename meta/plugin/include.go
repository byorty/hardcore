package plugin

import (
	"github.com/byorty/hardcore/meta"
	"path/filepath"
	"github.com/byorty/hardcore/utils"
	"io/ioutil"
	"encoding/xml"
	"github.com/byorty/hardcore/meta/controller"
	"sort"
    "github.com/byorty/hardcore/meta/model"
)

type Include struct {}

func (i *Include) Do(env *meta.Environment) {
	config := env.Configuration
	for _, include := range config.Includes {
		env.Logger.Debug("find include %s", include.File)
		incFilename := filepath.Join(env.MetaPath, include.File)
		if utils.FileExists(incFilename) {
			env.Logger.Debug("include file %s is exists", include.File)
			data, err := ioutil.ReadFile(incFilename)
			if err == nil {
				env.Logger.Debug("success read include file %s", include.File)
				var incConfig meta.Configuration
				err = xml.Unmarshal(data, &incConfig)
				if err == nil {
					env.Logger.Debug("success unmarshal include file %s", include.File)
					i.merge(&incConfig, env)
				} else {
					env.Logger.Error("can't unmarshal include file %s", include.File)
				}
			} else {
				env.Logger.Error("can't read include file %s", include.File)
			}
		} else {
			env.Logger.Error("include file %s not found", include.File)
		}
	}
}

func (i *Include) merge(src *meta.Configuration, env *meta.Environment) {
    i.mergeModelContainer(src.ModelContainers, env)
	i.mergeControllerContainer(src.ControllerContainers, env)
}

func (i *Include) mergeModelContainer(containers []*model.Container, env *meta.Environment) {
    if containers != nil && len(containers) > 0 {
        dest := env.Configuration
        if dest.ModelContainers == nil {
            dest.ModelContainers = make([]*model.Container, 0)
        }
        for _, newContainer := range containers {
            env.Logger.Debug("check model container %s", newContainer.Package)
            i := sort.Search(len(dest.ModelContainers), func(i int) bool {
                return dest.ModelContainers[i].Eq(newContainer)
            })
            if i < len(dest.ModelContainers) && dest.ModelContainers[i].Eq(newContainer) {
                env.Logger.Error("model container %s is duplicate", newContainer.Package)
            } else {
                env.Logger.Debug("insert model container %s in config", newContainer.Package)
                dest.ModelContainers = append(dest.ModelContainers, newContainer)
            }
        }
    }
}

func (i *Include) mergeControllerContainer(containers []*controller.Container, env *meta.Environment) {
	if containers != nil && len(containers) > 0 {
		dest := env.Configuration
		if dest.ControllerContainers == nil {
			dest.ControllerContainers = make([]*controller.Container, 0)
		}
		for _, newContainer := range containers {
			env.Logger.Debug("check controller container %s", newContainer.Package)
			i := sort.Search(len(dest.ControllerContainers), func(i int) bool {
				return dest.ControllerContainers[i].Eq(newContainer)
			})
			if i < len(dest.ControllerContainers) && dest.ControllerContainers[i].Eq(newContainer) {
				env.Logger.Error("controller container %s is duplicate", newContainer.Package)
			} else {
				env.Logger.Debug("insert controller container %s in config", newContainer.Package)
				dest.ControllerContainers = append(dest.ControllerContainers, newContainer)
			}
		}
	}
}
