package slice

import (
	"github.com/byorty/hardcore/types"
)

type Ints []int

func NewInts() types.IntSlice {
	return make(Ints, 0)
}

func NewIntsBy(arr []int) types.IntSlice {
	return Ints(arr)
}

func (i Ints) Len() int {
	return len(i)
}

func (i Ints) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i Ints) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Ints) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i Ints) Get(x int) int {
	return i[x]
}

func (i Ints) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt(i[x])
}

type Int8s []int8

func NewInt8s() types.Int8Slice {
	return make(Int8s, 0)
}

func NewInt8sBy(arr []int8) types.Int8Slice {
	return Int8s(arr)
}

func (i Int8s) Len() int {
	return len(i)
}

func (i Int8s) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i Int8s) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Int8s) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i Int8s) Get(x int) int8 {
	return i[x]
}

func (i Int8s) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt8(i[x])
}

type Int16s []int16

func NewInt16s() types.Int16Slice {
	return make(Int16s, 0)
}

func NewInt16sBy(arr []int16) types.Int16Slice {
	return Int16s(arr)
}

func (i Int16s) Len() int {
	return len(i)
}

func (i Int16s) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i Int16s) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Int16s) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i Int16s) Get(x int) int16 {
	return i[x]
}

func (i Int16s) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt16(i[x])
}

type Int32s []int32

func NewInt32s() types.Int32Slice {
	return make(Int32s, 0)
}

func NewInt32sBy(arr []int32) types.Int32Slice {
	return Int32s(arr)
}

func (i Int32s) Len() int {
	return len(i)
}

func (i Int32s) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i Int32s) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Int32s) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i Int32s) Get(x int) int32 {
	return i[x]
}

func (i Int32s) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt32(i[x])
}

type Int64s []int64

func NewInt64s() types.Int64Slice {
	return make(Int64s, 0)
}

func NewInt64sBy(arr []int64) types.Int64Slice {
	return Int64s(arr)
}

func (i Int64s) Len() int {
	return len(i)
}

func (i Int64s) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i Int64s) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i Int64s) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i Int64s) Get(x int) int64 {
	return i[x]
}

func (i Int64s) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt64(i[x])
}
