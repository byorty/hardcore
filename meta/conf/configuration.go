package conf

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/controller"
	"github.com/byorty/hardcore/meta/exporter"
	"github.com/byorty/hardcore/meta/model"
	"github.com/byorty/hardcore/meta/types"
)

type Configuration struct {
	Includes             []common.Include      `xml:"includes>include"`
	ControllerContainers controller.Containers `xml:"controllers"`
	ModelContainers      model.Containers      `xml:"models"`
	ExporterContainers   exporter.Containers   `xml:"exporters"`
	Files                []types.File
	сontainers           []types.Container
}

func (c *Configuration) Init() {
	c.сontainers = make([]types.Container, 0)
	c.addContainers(c.ControllerContainers)
	c.addContainers(c.ModelContainers)
	c.addContainers(c.ExporterContainers)
}

func (c *Configuration) addContainers(slice types.ContainerSlice) {
	for i := 0; i < slice.Len(); i++ {
		c.сontainers = append(c.сontainers, slice.Get(i))
	}
}

func (c Configuration) GetContainers() []types.Container {
	return c.сontainers
}

func (c *Configuration) SetContainers(containers []types.Container) {
	c.сontainers = containers
}

func (c *Configuration) AddFile(name, tpl string, params map[string]interface{}) {
	c.Files = append(c.Files, common.NewFile(name+types.FileExt, tpl, params, types.NormalFileKind))
}

func (c *Configuration) AddAutoFile(name, tpl string, params map[string]interface{}) {
	c.Files = append(c.Files, common.NewFile(name+types.FileExt, tpl, params, types.AutoFileKind))
}

func (c Configuration) GetFiles() []types.File {
	return c.Files
}
