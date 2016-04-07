package app

import (
	"github.com/byorty/hardcore/app/plugin"
	"github.com/byorty/hardcore/types"
)

type Application struct {
	plugins []types.ApplicationPlugin
}

func New() types.Application {
	return &Application{
		plugins: make([]types.ApplicationPlugin, 0),
	}
}

func (a *Application) AddPlugin(plugin types.ApplicationPlugin) types.Application {
	a.plugins = append(a.plugins, plugin)
	return a
}

func (a Application) Run() {
	for _, plugin := range a.plugins {
		plugin.Run()
	}
}

func NewHttpServer() types.Application {
	return New().
		AddPlugin(new(plugin.RouterImpl)).
		AddPlugin(new(plugin.DAOImpl)).
		AddPlugin(new(plugin.TmplImpl)).
		AddPlugin(new(plugin.SessionImpl)).
		AddPlugin(new(plugin.HttpImpl))
}
