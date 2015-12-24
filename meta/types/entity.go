package types

type EntityKind int

const (
    ControllerEntityKind EntityKind = iota
    ModelEntityKind
    EnumEntityKind
)

type Entity interface {
    GetKind() EntityKind
    GetName() string
    SetExtends([]Entity)
    GetExtends() []Entity
    SetImports([]string)
    GetImports() []string
    SetFilename(string)
    GetFilename()string
    SetAutoFilename(string)
    GetAutoFilename()string
    GetRawExtends() []string
}