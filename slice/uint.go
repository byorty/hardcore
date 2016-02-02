package slice
import "github.com/byorty/hardcore/types"

type UintsImpl []uint

func NewUints() types.UintSlice {
	return make(UintsImpl, 0)
}

func NewUintsBy(arr []uint) types.UintSlice {
	return UintsImpl(arr)
}

func (u UintsImpl) Len() int {
	return len(u)
}

func (u UintsImpl) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u UintsImpl) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u UintsImpl) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u UintsImpl) Get(x int) uint {
	return u[x]
}

type Uint8sImpl []uint8

func NewUint8s() types.Uint8Slice {
	return make(Uint8sImpl, 0)
}

func NewUint8sBy(arr []uint8) types.Uint8Slice {
	return Uint8sImpl(arr)
}

func (u Uint8sImpl) Len() int {
	return len(u)
}

func (u Uint8sImpl) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u Uint8sImpl) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Uint8sImpl) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Uint8sImpl) Get(x int) uint8 {
	return u[x]
}

type Uint16sImpl []uint16

func NewUint16s() types.Uint16Slice {
	return make(Uint16sImpl, 0)
}

func NewUint16sBy(arr []uint16) types.Uint16Slice {
	return Uint16sImpl(arr)
}

func (u Uint16sImpl) Len() int {
	return len(u)
}

func (u Uint16sImpl) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u Uint16sImpl) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Uint16sImpl) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Uint16sImpl) Get(x int) uint16 {
	return u[x]
}

type Uint32sImpl []uint32

func NewUint32s() types.Uint32Slice {
	return make(Uint32sImpl, 0)
}

func NewUint32sBy(arr []uint32) types.Uint32Slice {
	return Uint32sImpl(arr)
}

func (u Uint32sImpl) Len() int {
	return len(u)
}

func (u Uint32sImpl) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u Uint32sImpl) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Uint32sImpl) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Uint32sImpl) Get(x int) uint32 {
	return u[x]
}

type Uint64sImpl []uint64

func NewUint64s() types.Uint64Slice {
	return make(Uint64sImpl, 0)
}

func NewUint64sBy(arr []uint64) types.Uint64Slice {
	return Uint64sImpl(arr)
}

func (u Uint64sImpl) Len() int {
	return len(u)
}

func (u Uint64sImpl) Less(x, y int) bool {
	return u[x] < u[y]
}

func (u Uint64sImpl) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

func (u Uint64sImpl) GetRaw(x int) interface{} {
	return u.Get(x)
}

func (u Uint64sImpl) Get(x int) uint64 {
	return u[x]
}
