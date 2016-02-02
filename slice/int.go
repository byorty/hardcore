package slice

import "github.com/byorty/hardcore/types"

type IntsImpl []int

func NewInts() types.IntSlice {
	return make(IntsImpl, 0)
}

func NewIntsBy(arr []int) types.IntSlice {
	return IntsImpl(arr)
}

func (i IntsImpl) Len() int {
	return len(i)
}

func (i IntsImpl) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i IntsImpl) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i IntsImpl) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i IntsImpl) Get(x int) int {
	return i[x]
}

type Int8sImpl []int8

func NewInt8s() types.Int8Slice {
	return make(Int8sImpl, 0)
}

func NewInt8sBy(arr []int8) types.Int8Slice {
	return Int8sImpl(arr)
}

func (i Int8sImpl) Len() int {
	return len(i)
}

func (i Int8sImpl) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i Int8sImpl) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Int8sImpl) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i Int8sImpl) Get(x int) int8 {
	return i[x]
}

type Int16sImpl []int16

func NewInt16s() types.Int16Slice {
	return make(Int16sImpl, 0)
}

func NewInt16sBy(arr []int16) types.Int16Slice {
	return Int16sImpl(arr)
}

func (i Int16sImpl) Len() int {
	return len(i)
}

func (i Int16sImpl) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i Int16sImpl) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Int16sImpl) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i Int16sImpl) Get(x int) int16 {
	return i[x]
}

type Int32sImpl []int32

func NewInt32s() types.Int32Slice {
	return make(Int32sImpl, 0)
}

func NewInt32sBy(arr []int32) types.Int32Slice {
	return Int32sImpl(arr)
}

func (i Int32sImpl) Len() int {
	return len(i)
}

func (i Int32sImpl) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i Int32sImpl) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Int32sImpl) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i Int32sImpl) Get(x int) int32 {
	return i[x]
}

type Int64sImpl []int64

func NewInt64s() types.Int64Slice {
	return make(Int64sImpl, 0)
}

func NewInt64sBy(arr []int64) types.Int64Slice {
	return Int64sImpl(arr)
}

func (i Int64sImpl) Len() int {
	return len(i)
}

func (i Int64sImpl) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i Int64sImpl) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Int64sImpl) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i Int64sImpl) Get(x int) int64 {
	return i[x]
}
