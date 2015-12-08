package plugin

import (
	"github.com/byorty/hardcore/meta"
	"path/filepath"
	"github.com/byorty/hardcore/utils"
	"io/ioutil"
	"encoding/xml"
	"github.com/byorty/hardcore/meta/controller"
)

type IncludePlugin struct {}

func NewIncludePlugin() Plugin {
	return new(IncludePlugin)
}

func (i *IncludePlugin) Do(config *meta.Configuration) {
	for _, include := range config.Includes {
		incFilename := filepath.Join(config.MetaPath, include.File)
		if utils.FileExists(incFilename) {
			data, err := ioutil.ReadFile(incFilename)
			if err == nil {
				var incConfig meta.Configuration
				err = xml.Unmarshal(data, &incConfig)
				if err == nil {
					i.merge(&incConfig, config)
				} else {

				}
			} else {

			}
		} else {

		}
	}
}

func (i *IncludePlugin) merge(src, dest *meta.Configuration) {
	i.mergeControllers(src.ControllerContainers, dest)
}

func (i *IncludePlugin) mergeControllers(containers []controller.Container, dest *meta.Configuration) {
	if containers != nil && len(containers) > 0 {
		if dest.ControllerContainers == nil {
//			dest.ControllerContainers = make([]controller.Controller, 0)
		}
//		for _, newContainer := range containers {
//
//			for _, existsContainer := range dest.ControllerContainers {
//
//			}
//		}
	}
}
