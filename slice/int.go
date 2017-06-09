package slice

import (
	"github.com/byorty/hardcore/types"
)

type _Int []int

func NewInt() types.IntSlice {
	return NewIntBy(make([]int, 0))
}

func NewIntBy(arr []int) types.IntSlice {
	s := _Int(arr)
	return &s
}

func (i _Int) Len() int {
	return len(i)
}

func (i _Int) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i _Int) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i _Int) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i _Int) Get(x int) int {
	return i[x]
}

func (i *_Int) Add(v int) {
	(*i) = append((*i), v)
}

func (i _Int) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt(i[x])
}

type _Int8 []int8

func NewInt8() types.Int8Slice {
	return NewInt8By(make([]int8, 0))
}

func NewInt8By(arr []int8) types.Int8Slice {
	s := _Int8(arr)
	return &s
}

func (i _Int8) Len() int {
	return len(i)
}

func (i _Int8) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i _Int8) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i _Int8) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i _Int8) Get(x int) int8 {
	return i[x]
}

func (i *_Int8) Add(v int8) {
	(*i) = append((*i), v)
}

func (i _Int8) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt8(i[x])
}

type _Int16 []int16

func NewInt16() types.Int16Slice {
	return NewInt16By(make([]int16, 0))
}

func NewInt16By(arr []int16) types.Int16Slice {
	s := _Int16(arr)
	return &s
}

func (i _Int16) Len() int {
	return len(i)
}

func (i _Int16) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i _Int16) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i _Int16) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i _Int16) Get(x int) int16 {
	return i[x]
}

func (i *_Int16) Add(v int16) {
	(*i) = append((*i), v)
}

func (i _Int16) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt16(i[x])
}

type _Int32 []int32

func NewInt32() types.Int32Slice {
	return NewInt32By(make([]int32, 0))
}

func NewInt32By(arr []int32) types.Int32Slice {
	s := _Int32(arr)
	return &s
}

func (i _Int32) Len() int {
	return len(i)
}

func (i _Int32) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i _Int32) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i _Int32) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i _Int32) Get(x int) int32 {
	return i[x]
}

func (i *_Int32) Add(v int32) {
	(*i) = append((*i), v)
}

func (i _Int32) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt32(i[x])
}

type _Int64 []int64

func NewInt64() types.Int64Slice {
	return NewInt64By(make([]int64, 0))
}

func NewInt64By(arr []int64) types.Int64Slice {
	s := _Int64(arr)
	return &s
}

func (i _Int64) Len() int {
	return len(i)
}

func (i _Int64) Less(x, y int) bool {
	return i[x] < i[y]
}

func (i _Int64) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i _Int64) GetRaw(x int) interface{} {
	return i.Get(x)
}

func (i _Int64) Get(x int) int64 {
	return i[x]
}

func (i *_Int64) Add(v int64) {
	(*i) = append((*i), v)
}

func (i _Int64) Export(x int, encoder types.Encoder) {
	encoder.EncodeInt64(i[x])
}
