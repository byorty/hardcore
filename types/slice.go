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
	Add(int)
	Get(int) int
}

type Int8Slice interface {
	Slice
	Add(int8)
	Get(int) int8
}

type Int16Slice interface {
	Slice
	Add(int16)
	Get(int) int16
}

type Int32Slice interface {
	Slice
	Add(int32)
	Get(int) int32
}

type Int64Slice interface {
	Slice
	Add(int64)
	Get(int) int64
}

type UintSlice interface {
	Slice
	Add(uint)
	Get(int) uint
}

type Uint8Slice interface {
	Slice
	Add(uint8)
	Get(int) uint8
}

type Uint16Slice interface {
	Slice
	Add(uint16)
	Get(int) uint16
}

type Uint32Slice interface {
	Slice
	Add(uint32)
	Get(int) uint32
}

type Uint64Slice interface {
	Slice
	Add(uint64)
	Get(int) uint64
}

type Float32Slice interface {
	Slice
	Add(float32)
	Get(int) float32
}

type Float64Slice interface {
	Slice
	Add(float64)
	Get(int) float64
}

type StringSlice interface {
	Slice
	Add(string)
	Get(int) string
}

type ByteSlice interface {
	Slice
	Add(byte)
	Get(int) byte
}
