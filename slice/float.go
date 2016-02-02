package slice

import "github.com/byorty/hardcore/types"

type Float32sImpl []float32

func NewFloat32s() types.Float32Slice {
	return make(Float32sImpl, 0)
}

func NewFloat32sBy(arr []float32) types.Float32Slice {
	return Float32sImpl(arr)
}

func (f Float32sImpl) Len() int {
	return len(f)
}

func (f Float32sImpl) Less(x, y int) bool {
	return f[x] < f[y]
}

func (f Float32sImpl) Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f Float32sImpl) GetRaw(x int) interface{} {
	return f.Get(x)
}

func (f Float32sImpl) Get(x int) float32 {
	return f[x]
}

type Float64sImpl []float64

func NewFloat64s() types.Float64Slice {
	return make(Float64sImpl, 0)
}

func NewFloat64sBy(arr []float64) types.Float64Slice {
	return Float64sImpl(arr)
}

func (f Float64sImpl) Len() int {
	return len(f)
}

func (f Float64sImpl) Less(x, y int) bool {
	return f[x] < f[y]
}

func (f Float64sImpl) Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f Float64sImpl) GetRaw(x int) interface{} {
	return f.Get(x)
}

func (f Float64sImpl) Get(x int) float64 {
	return f[x]
}
