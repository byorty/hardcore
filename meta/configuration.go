package meta

import (
	"github.com/byorty/hardcore/meta/common"
	"github.com/byorty/hardcore/meta/controller"
)

type Configuration struct {
	Includes []common.Include `xml:"includes>include"`
	ControllerContainers []*controller.Container `xml:"controllers"`
}
