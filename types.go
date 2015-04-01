package hardcore

const (
	EOL = '\n'
)

type Named interface {
	GetName() string
}

type Enum interface {
	Named
	GetId() int
}

type Enum8 interface {
	Named
	GetId() int8
}

type Enum16 interface {
	Named
	GetId() int16
}

type Enum32 interface {
	Named
	GetId() int32
}

type Enum64 interface {
	Named
	GetId() int64
}

type UEnum interface {
	Named
	GetId() uint
}

type UEnum8 interface {
	Named
	GetId() uint8
}

type UEnum16 interface {
	Named
	GetId() uint16
}

type UEnum32 interface {
	Named
	GetId() uint32
}

type UEnum64 interface {
	Named
	GetId() uint64
}

type Float32Enum interface {
	Named
	GetId() float32
}

type Float64Enum interface {
	Named
	GetId() float64
}

type Complex64Enum interface {
	Named
	GetId() complex64
}

type Complex128Enum interface {
	Named
	GetId() complex128
}

type StringEnum interface {
	Named
	GetId() string
}
