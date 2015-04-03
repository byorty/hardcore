package hardcore

const (
	EOL = '\n'
)

type Named interface {
	GetName() string
}

type IntEnum interface {
	Named
	IntIdentifiable
}

type Int8Enum interface {
	Named
	Int8Identifiable
}

type Int16Enum interface {
	Named
	Int16Identifiable
}

type Int32Enum interface {
	Named
	Int32Identifiable
}

type Int64Enum interface {
	Named
	Int64Identifiable
}

type UintEnum interface {
	Named
	UintIdentifiable
}

type Uint8Enum interface {
	Named
	Uint8Identifiable
}

type Uint16Enum interface {
	Named
	Uint16Identifiable
}

type Uint32Enum interface {
	Named
	Uint32Identifiable
}

type Uint64Enum interface {
	Named
	Uint64Identifiable
}

type Float32Enum interface {
	Named
	Float32Identifiable
}

type Float64Enum interface {
	Named
	Float64Identifiable
}

type Complex64Enum interface {
	Named
	Complex64Identifiable
}

type Complex128Enum interface {
	Named
	Complex128Identifiable
}

type StringEnum interface {
	Named
	StringIdentifiable
}

type RuneEnum interface {
	Named
	RuneIdentifiable
}
