package slice

import "github.com/byorty/hardcore/types"

type _Uint []uint

func NewUint() types.UintSlice {
	return NewUintBy(make([]uint, 0))
}

func NewUintBy(arr []uint) types.UintSlice {
	s := _Uint(arr)
	return &s
}

func (u _Uint) Len() int {
	return len(u)
}

func (u _Uint) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u _Uint) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u _Uint) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u _Uint) Get(x int) uint {
	return u[x]
}

func (u *_Uint) Add(v uint) {
	(*u) = append((*u), v)
}

func (u _Uint) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint(u[i])
}

type _Uint8 []uint8

func NewUint8() types.Uint8Slice {
	return NewUint8By(make([]uint8, 0))
}

func NewUint8By(arr []uint8) types.Uint8Slice {
	s := _Uint8(arr)
	return &s
}

func (u _Uint8) Len() int {
	return len(u)
}

func (u _Uint8) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u _Uint8) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u _Uint8) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u _Uint8) Get(x int) uint8 {
	return u[x]
}

func (u *_Uint8) Add(v uint8) {
	(*u) = append((*u), v)
}

func (u _Uint8) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint8(u[i])
}

type _Uint16 []uint16

func NewUint16() types.Uint16Slice {
	return NewUint16By(make([]uint16, 0))
}

func NewUint16By(arr []uint16) types.Uint16Slice {
	s:= _Uint16(arr)
	return &s
}

func (u _Uint16) Len() int {
	return len(u)
}

func (u _Uint16) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u _Uint16) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u _Uint16) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u _Uint16) Get(x int) uint16 {
	return u[x]
}

func (u *_Uint16) Add(v uint16) {
	(*u) = append((*u), v)
}

func (u _Uint16) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint16(u[i])
}

type _Uint32 []uint32

func NewUint32() types.Uint32Slice {
	return NewUint32By(make([]uint32, 0))
}

func NewUint32By(arr []uint32) types.Uint32Slice {
	s := _Uint32(arr)
	return &s
}

func (u _Uint32) Len() int {
	return len(u)
}

func (u _Uint32) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u _Uint32) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u _Uint32) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u _Uint32) Get(x int) uint32 {
	return u[x]
}

func (u *_Uint32) Add(v uint32) {
	(*u) = append((*u), v)
}

func (u _Uint32) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint32(u[i])
}

type _Uint64 []uint64

func NewUint64() types.Uint64Slice {
	return NewUint64By(make([]uint64, 0))
}

func NewUint64By(arr []uint64) types.Uint64Slice {
	s := _Uint64(arr)
	return &s
}

func (u _Uint64) Len() int {
	return len(u)
}

func (u _Uint64) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u _Uint64) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u _Uint64) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u _Uint64) Get(x int) uint64 {
	return u[x]
}

func (u *_Uint64) Add(v uint64) {
	(*u) = append((*u), v)
}

func (u _Uint64) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint64(u[i])
}
