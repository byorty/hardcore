package types

type ContainerKind int

const (
    ControllerContainerKind ContainerKind = iota
    EntityContainerKind
    ExporterContainerKind
)

type Container interface {
    GetPackage() string
    SetShortPackage(string)
    GetShortPackage() string
    SetImport(string)
    GetImport() string
    SetPath(string)
    GetPath()string
    GetContainerKind() ContainerKind
    SetEntities([]Entity)
    GetEntities() []Entity
}
