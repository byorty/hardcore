package slice

type Float32s []float32

func NewFloat32s() Float32s {
	return make(Float32s, 0)
}

func NewFloat32sBy(arr []float32) Float32s {
	return Float32s(arr)
}

func (f Float32s) Len() int {
	return len(f)
}

func (f Float32s) Less(x, y int) bool {
	return f[x] < f[y]
}

func (f Float32s) Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f Float32s) GetRaw(x int) interface{} {
	return f.Get(x)
}

func (f Float32s) Get(x int) float32 {
	return f[x]
}

type Float64s []float64

func NewFloat64s() Float64s {
	return make(Float64s, 0)
}

func NewFloat64sBy(arr []float64) Float64s {
	return Float64s(arr)
}

func (f Float64s) Len() int {
	return len(f)
}

func (f Float64s) Less(x, y int) bool {
	return f[x] < f[y]
}

func (f Float64s) Swap(x, y int) {
	f[x], f[y] = f[y], f[x]
}

func (f Float64s) GetRaw(x int) interface{} {
	return f.Get(x)
}

func (f Float64s) Get(x int) float64 {
	return f[x]
}
