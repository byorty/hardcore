package exporter

import (
"github.com/byorty/hardcore/meta/common"
"github.com/byorty/hardcore/meta/types"
)

type Container struct {
	common.Container
	Exporters []*Exporter `xml:"exporter"`
}

func (c Container) GetContainerKind() types.ContainerKind {
	return types.ExporterContainerKind
}
