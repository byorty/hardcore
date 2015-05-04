package types

type IntIdentifier interface {
	ById(int) IntIdentifiable
}

type IntIdentifiers interface {
	ByIds([]int)
}

type Int8Identifier interface {
	ById(int8) Int8Identifiable
}

type Int8Identifiers interface {
	ByIds([]int8)
}

type Int16Identifier interface {
	ById(int16) Int16Identifiable
}

type Int16Identifiers interface {
	ByIds([]int16)
}

type Int32Identifier interface {
	ById(int32) Int32Identifiable
}

type Int32Identifiers interface {
	ByIds([]int32)
}

type Int64Identifier interface {
	ById(int64) Int64Identifiable
}

type Int64Identifiers interface {
	ByIds([]int64)
}

type UintIdentifier interface {
	ById(uint) UintIdentifiable
}

type UintIdentifiers interface {
	ByIds([]uint)
}

type Uint8Identifier interface {
	ById(uint8) Uint8Identifiable
}

type Uint8Identifiers interface {
	ByIds([]uint8)
}

type Uint16Identifier interface {
	ById(uint16) Uint16Identifiable
}

type Uint16Identifiers interface {
	ByIds([]uint16)
}

type Uint32Identifier interface {
	ById(uint32) Uint32Identifiable
}

type Uint32Identifiers interface {
	ByIds([]uint32)
}

type Uint64Identifier interface {
	ById(uint64) Uint64Identifiable
}

type Uint64Identifiers interface {
	ByIds([]uint64)
}

type Float32Identifier interface {
	ById(float32) Float32Identifiable
}

type Float32Identifiers interface {
	ByIds([]float32)
}

type Float64Identifier interface {
	ById(float64) Float64Identifiable
}

type Float64Identifiers interface {
	ByIds([]float64)
}

type Complex64Identifier interface {
	ById(complex64) Complex64Identifiable
}

type Complex64Identifiers interface {
	ByIds([]complex64)
}

type Complex128Identifier interface {
	ById(complex128) Complex128Identifiable
}

type Complex128Identifiers interface {
	ByIds([]complex128)
}

type StringIdentifier interface {
	ById(string) StringIdentifiable
}

type StringIdentifiers interface {
	ByIds([]string)
}

type RuneIdentifier interface {
	ById(rune) RuneIdentifiable
}

type RuneIdentifiers interface {
	ByIds([]rune)
}
