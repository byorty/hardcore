package types

type EntityKind int

func (e EntityKind) IsModel() bool {
	return e == ModelEntityKind
}

func (e EntityKind) IsEnum() bool {
	return e == EnumEntityKind
}

const (
	ControllerEntityKind EntityKind = iota
	ModelEntityKind
	EnumEntityKind
	ExporterEntityKind
)

type EntitySlice interface {
	Len() int
	Get(int) Entity
}

type Entity interface {
	GetEntityKind() EntityKind
	GetName() string
	GetFullName() string
	GetPointerName() string
	GetPointerFullName() string
	SetExtends([]Entity)
	GetExtends() []Entity
	SetImports([]string)
	AddImport(string)
	GetImports() []string
	SetFilename(string)
	GetFilename() string
	SetAutoFilename(string)
	GetAutoFilename() string
	GetRawExtends() []string
	SetContainer(Container)
	GetContainer() Container
	ClearName()
	GetMultipleName() string
	GetFullMultipleName() string
	Init(Container)
}

type PostEntity interface {
	PostInit(Environment)
}
