package types

type Identifiable interface {
	GetId() int
}

type Int8Identifiable interface {
	GetId() int8
}

type Int16Identifiable interface {
	GetId() int16
}

type Int32Identifiable interface {
	GetId() int32
}

type Int64Identifiable interface {
	GetId() int64
}

type UintIdentifiable interface {
	GetId() uint
}

type Uint8Identifiable interface {
	GetId() uint8
}

type Uint16Identifiable interface {
	GetId() uint16
}

type Uint32Identifiable interface {
	GetId() uint32
}

type Uint64Identifiable interface {
	GetId() uint64
}

type Float32Identifiable interface {
	GetId() float32
}

type Float64Identifiable interface {
	GetId() float64
}

type Complex64Identifiable interface {
	GetId() complex64
}

type Complex128Identifiable interface {
	GetId() complex128
}

type StringIdentifiable interface {
	GetId() string
}

type RuneIdentifiable interface {
	GetId() rune
}
