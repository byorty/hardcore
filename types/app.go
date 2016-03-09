package types

type Application interface {
	Run()
	AddPlugin(ApplicationPlugin) Application
}

type ApplicationPlugin interface {
	Run()
}
