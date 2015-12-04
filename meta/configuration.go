package meta

import (
	"github.com/byorty/hardcore/meta/controller"
	"github.com/byorty/hardcore/meta/common"
)

type Configuration struct {
	Includes []common.Include `xml:"includes>include"`
	ControllerContainers []controller.Container `xml:"controllers"`
}
