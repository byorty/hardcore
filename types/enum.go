package types

type Named interface {
	GetRawId() interface{}
	GetName() string
}

type Enum interface {
	Named
	Identifiable
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

type EnumDAO interface {
    EnumScanner
    GetList() []Named
    Eq(Named, interface{}) bool
    Scan(src, dest Named)
}

type IntEnumDAO interface {
    EnumDAO
    ById(int) EnumScanner
}

type Int8EnumDAO interface {
    EnumDAO
    ById(int8) EnumScanner
}

type Int16EnumDAO interface {
    EnumDAO
    ById(int16) EnumScanner
}

type Int32EnumDAO interface {
    EnumDAO
    ById(int32) EnumScanner
}

type Int64EnumDAO interface {
    EnumDAO
    ById(int64) EnumScanner
}

type UintEnumDAO interface {
    EnumDAO
    ById(uint) EnumScanner
}

type Uint8EnumDAO interface {
    EnumDAO
    ById(uint8) EnumScanner
}

type Uint16EnumDAO interface {
    EnumDAO
    ById(uint16) EnumScanner
}

type Uint32EnumDAO interface {
    EnumDAO
    ById(uint32) EnumScanner
}

type Uint64EnumDAO interface {
    EnumDAO
    ById(uint64) EnumScanner
}

type Float32EnumDAO interface {
    EnumDAO
    ById(float32) EnumScanner
}

type Float64EnumDAO interface {
    EnumDAO
    ById(float64) EnumScanner
}

type Complex64EnumDAO interface {
    EnumDAO
    ById(complex64) EnumScanner
}

type Complex128EnumDAO interface {
    EnumDAO
    ById(complex128) EnumScanner
}

type StringEnumDAO interface {
    EnumDAO
    ById(string) EnumScanner
}

type RuneEnumDAO interface {
    EnumDAO
    ById(rune) EnumScanner
}
