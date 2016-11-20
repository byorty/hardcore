package types

const (
	TimeFormat = "2006-12-21 15:04:05"
)

type Application interface {
	Run()
	AddPlugin(ApplicationPlugin) Application
}

type ApplicationPlugin interface {
	Run()
}
