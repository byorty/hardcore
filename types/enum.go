package types

type Named interface {
	GetRawId() interface{}
	GetName() string
}

type IntEnum interface {
	Named
    IntIdentifiable
	DAO() IntEnumDAO
}

type Int8Enum interface {
	Named
	Int8Identifiable
	DAO() Int8EnumDAO
}

type Int16Enum interface {
	Named
	Int16Identifiable
	DAO() Int16EnumDAO
}

type Int32Enum interface {
	Named
	Int32Identifiable
	DAO() Int32EnumDAO
}

type Int64Enum interface {
	Named
	Int64Identifiable
	DAO() Int64EnumDAO
}

type UintEnum interface {
	Named
	UintIdentifiable
	DAO() UintEnumDAO
}

type Uint8Enum interface {
	Named
	Uint8Identifiable
	DAO() Uint8EnumDAO
}

type Uint16Enum interface {
	Named
	Uint16Identifiable
	DAO() Uint16EnumDAO
}

type Uint32Enum interface {
	Named
	Uint32Identifiable
	DAO() Uint32EnumDAO
}

type Uint64Enum interface {
	Named
	Uint64Identifiable
	DAO() Uint64EnumDAO
}

type Float32Enum interface {
	Named
	Float32Identifiable
	DAO() Float32EnumDAO
}

type Float64Enum interface {
	Named
	Float64Identifiable
	DAO() Float64EnumDAO
}

type StringEnum interface {
	Named
	StringIdentifiable
	DAO() StringEnumDAO
}

type EnumDAO interface {
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

type StringEnumDAO interface {
    EnumDAO
    ById(string) EnumScanner
}

