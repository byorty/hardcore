package types

import (
	"sort"
)

type Slice interface {
	sort.Interface
	GetRaw(int) interface{}
}

type IntSlice interface {
	Slice
	Get(int) int
}

type Int8Slice interface {
	Slice
	Get(int) int8
}

type Int16Slice interface {
	Slice
	Get(int) int16
}

type Int32Slice interface {
	Slice
	Get(int) int32
}

type Int64Slice interface {
	Slice
	Get(int) int64
}

type UintSlice interface {
	Slice
	Get(int) uint
}

type Uint8Slice interface {
	Slice
	Get(int) uint8
}

type Uint16Slice interface {
	Slice
	Get(int) uint16
}

type Uint32Slice interface {
	Slice
	Get(int) uint32
}

type Uint64Slice interface {
	Slice
	Get(int) uint64
}

type Float32Slice interface {
	Slice
	Get(int) float32
}

type Float64Slice interface {
	Slice
	Get(int) float64
}

type StringSlice interface {
	Slice
	Get(int) string
}
