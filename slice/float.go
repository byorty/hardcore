package slice

import "github.com/byorty/hardcore/types"

type _Float32 []float32

func NewFloat32() types.Float32Slice {
	return NewFloat32By(make([]float32, 0))
}

func NewFloat32By(arr []float32) types.Float32Slice {
	s := _Float32(arr)
	return &s
}

func (f _Float32) Len() int {
	return len(f)
}

func (f _Float32) Less(x, y int) bool {
	return f[x] < f[y]
}

func (f _Float32) Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f _Float32) GetRaw(x int) interface{} {
	return f.Get(x)
}

func (f _Float32) Get(x int) float32 {
	return f[x]
}

func (f *_Float32) Add(v float32) {
	(*f) = append((*f), v)
}

func (f _Float32) Export(i int, encoder types.Encoder) {
	encoder.EncodeFloat32(f[i])
}

type _Float64 []float64

func NewFloat64() types.Float64Slice {
	return NewFloat64By(make([]float64, 0))
}

func NewFloat64By(arr []float64) types.Float64Slice {
	s := _Float64(arr)
	return &s
}

func (f _Float64) Len() int {
	return len(f)
}

func (f _Float64) Less(x, y int) bool {
	return f[x] < f[y]
}

func (f _Float64) Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f _Float64) GetRaw(x int) interface{} {
	return f.Get(x)
}

func (f _Float64) Get(x int) float64 {
	return f[x]
}

func (f *_Float64) Add(v float64) {
	(*f) = append((*f), v)
}

func (f _Float64) Export(i int, encoder types.Encoder) {
	encoder.EncodeFloat64(f[i])
}
