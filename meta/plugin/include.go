package plugin

import (
	"github.com/byorty/hardcore/meta"
	"path/filepath"
	"github.com/byorty/hardcore/utils"
	"io/ioutil"
	"encoding/xml"
	"github.com/byorty/hardcore/meta/controller"
	"sort"
)

type IncludePlugin struct {}

func NewIncludePlugin() Plugin {
	return new(IncludePlugin)
}

func (i *IncludePlugin) Do(env *meta.Environment) {
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

func (i *IncludePlugin) merge(src *meta.Configuration, env *meta.Environment) {
	i.mergeControllerContainer(src.ControllerContainers, env)
}

func (i *IncludePlugin) mergeControllerContainer(containers []controller.Container, env *meta.Environment) {
	if containers != nil && len(containers) > 0 {
		dest := env.Configuration
		if dest.ControllerContainers == nil {
			dest.ControllerContainers = make([]controller.Container, 0)
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
