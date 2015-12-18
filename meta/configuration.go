package meta

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/controller"
	"github.com/byorty/hardcore/meta/model"
)

type Configuration struct {
	Includes []common.Include `xml:"includes>include"`
	ControllerContainers []*controller.Container `xml:"controllers"`
	ModelContainers []*model.Container `xml:"models"`
	Files []common.File
}

func (c *Configuration) AddFile(name, tpl string, params map[string]interface{})  {
    c.Files = append(c.Files, common.File{name + common.FileExt, tpl, params, common.NormalFileKind})
}

func (c *Configuration) AddAutoFile(name, tpl string, params map[string]interface{})  {
    c.Files = append(c.Files, common.File{name + common.FileExt, tpl, params, common.AutoFileKind})
}
