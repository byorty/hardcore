package conf

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/controller"
	"github.com/byorty/hardcore/meta/model"
    "github.com/byorty/hardcore/meta/types"
)

type Configuration struct {
	Includes []common.Include `xml:"includes>include"`
	ControllerContainers []*controller.Container `xml:"controllers"`
	ModelContainers []*model.Container `xml:"models"`
	Files []common.File
	сontainers []types.Container
}

func (c Configuration) GetContainers() []types.Container {
    return c.сontainers
}

func (c *Configuration) SetContainers(containers []types.Container) {
    c.сontainers = containers
}

func (c *Configuration) AddFile(name, tpl string, params map[string]interface{})  {
    c.Files = append(c.Files, common.File{name + common.FileExt, tpl, params, common.NormalFileKind})
}

func (c *Configuration) AddAutoFile(name, tpl string, params map[string]interface{})  {
    c.Files = append(c.Files, common.File{name + common.FileExt, tpl, params, common.AutoFileKind})
}
