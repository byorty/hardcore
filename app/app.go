package app

import (
	"github.com/byorty/hardcore/app/plugin"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/scope"
)

type Application struct {
	name string
	plugins []types.ApplicationPlugin
}

func New() types.Application {
	return NewByName(scope.DefaultName)
}

func NewByName(name string) types.Application {
	return &Application{
		name: name,
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
	<- scope.AppByName(a.name).IsExit()
}

func NewHttpServer() types.Application {
	return New().
		AddPlugin(plugin.NewRouter()).
		AddPlugin(plugin.NewDAO()).
		AddPlugin(plugin.NewTmpl()).
		AddPlugin(plugin.NewSession()).
		AddPlugin(plugin.NewHttp())
}
