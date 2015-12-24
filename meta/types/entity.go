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
    GetFullname() string
    GetPointerName() string
    GetPointerFullName() string
    SetExtends([]Entity)
    GetExtends() []Entity
    SetImports([]string)
    AddImport(string)
    GetImports() []string
    SetFilename(string)
    GetFilename()string
    SetAutoFilename(string)
    GetAutoFilename()string
    GetRawExtends() []string
    SetContainer(Container)
    GetContainer() Container
    ClearName()
}