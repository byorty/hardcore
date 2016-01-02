package slice

type Uints []uint

func NewUints() Uints {
	return make(Uints, 0)
}

func NewUintsBy(arr []uint) Uints {
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

type Uint8s []uint8

func NewUint8s() Uint8s {
	return make(Uint8s, 0)
}

func NewUint8sBy(arr []uint8) Uint8s {
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

type Uint16s []uint16

func NewUint16s() Uint16s {
	return make(Uint16s, 0)
}

func NewUint16sBy(arr []uint16) Uint16s {
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

type Uint32s []uint32

func NewUint32s() Uint32s {
	return make(Uint32s, 0)
}

func NewUint32sBy(arr []uint32) Uint32s {
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

type Uint64s []uint64

func NewUint64s() Uint64s {
	return make(Uint64s, 0)
}

func NewUint64sBy(arr []uint64) Uint64s {
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
