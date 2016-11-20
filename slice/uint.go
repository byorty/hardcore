package slice

import "github.com/byorty/hardcore/types"

type Uints []uint

func NewUints() types.UintSlice {
	return make(Uints, 0)
}

func NewUintsBy(arr []uint) types.UintSlice {
	return Uints(arr)
}

func (u Uints) Len() int {
	return len(u)
}

func (u Uints) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u Uints) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Uints) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Uints) Get(x int) uint {
	return u[x]
}

func (u Uints) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint(u[i])
}

type Uint8s []uint8

func NewUint8s() types.Uint8Slice {
	return make(Uint8s, 0)
}

func NewUint8sBy(arr []uint8) types.Uint8Slice {
	return Uint8s(arr)
}

func (u Uint8s) Len() int {
	return len(u)
}

func (u Uint8s) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u Uint8s) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Uint8s) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Uint8s) Get(x int) uint8 {
	return u[x]
}

func (u Uint8s) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint8(u[i])
}

type Uint16s []uint16

func NewUint16s() types.Uint16Slice {
	return make(Uint16s, 0)
}

func NewUint16sBy(arr []uint16) types.Uint16Slice {
	return Uint16s(arr)
}

func (u Uint16s) Len() int {
	return len(u)
}

func (u Uint16s) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u Uint16s) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Uint16s) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Uint16s) Get(x int) uint16 {
	return u[x]
}

func (u Uint16s) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint16(u[i])
}

type Uint32s []uint32

func NewUint32s() types.Uint32Slice {
	return make(Uint32s, 0)
}

func NewUint32sBy(arr []uint32) types.Uint32Slice {
	return Uint32s(arr)
}

func (u Uint32s) Len() int {
	return len(u)
}

func (u Uint32s) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u Uint32s) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Uint32s) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Uint32s) Get(x int) uint32 {
	return u[x]
}

func (u Uint32s) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint32(u[i])
}

type Uint64s []uint64

func NewUint64s() types.Uint64Slice {
	return make(Uint64s, 0)
}

func NewUint64sBy(arr []uint64) types.Uint64Slice {
	return Uint64s(arr)
}

func (u Uint64s) Len() int {
	return len(u)
}

func (u Uint64s) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u Uint64s) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Uint64s) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Uint64s) Get(x int) uint64 {
	return u[x]
}

func (u Uint64s) Export(i int, encoder types.Encoder) {
	encoder.EncodeUint64(u[i])
}
