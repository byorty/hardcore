package plugin

import "github.com/byorty/hardcore/types"

func newByName(plugin types.ApplicationPlugin, name string) types.ApplicationPlugin {
	plugin.SetName(name)
	return plugin
}

type BaseImpl struct {
	name string
}

func (b *BaseImpl) SetName(name string) {
	b.name = name
}