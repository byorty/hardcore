package meta

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/controller"
)

type Configuration struct {
	AbsPath string
	MetaPath string
	Includes []common.Include `xml:"includes>include"`
	ControllerContainers []controller.Container `xml:"controllers"`
}

//func (c Configuration) GetControllerContainers() []types.ControllerContainer {
//	return c.ControllerContainers
//}
