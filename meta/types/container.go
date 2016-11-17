package types

type ContainerKind int

const (
	ControllerContainerKind ContainerKind = iota
	EntityContainerKind
	ExporterContainerKind
	ImporterContainerKind
	MiddlewaresContainerKind
)

type ContainerSlice interface {
	Get(int) Container
	Len() int
}

type Container interface {
	EntitySlice
	GetPackage() string
	SetShortPackage(string)
	GetShortPackage() string
	SetImport(string)
	GetImport() string
	SetPath(string)
	GetPath() string
	GetContainerKind() ContainerKind
	SetEntities([]Entity)
	GetEntities() []Entity
	Init(Environment)
	Merge(EntitySlice)
}
