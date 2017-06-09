package slice

import "github.com/byorty/hardcore/types"

type _String []string

func NewString() types.StringSlice {
	return NewStringBy(make([]string, 0))
}

func NewStringBy(arr []string) types.StringSlice {
	s := _String(arr)
	return &s
}

func (s _String) Len() int {
	return len(s)
}

func (s _String) Less(x, y int) bool {
	return s[x] < s[y]
}

func (s _String) Swap(x, y int) {
	s[x], s[y] = s[y], s[x]
}

func (s _String) GetRaw(x int) interface{} {
	return s.Get(x)
}

func (s _String) Get(x int) string {
	return s[x]
}

func (s *_String) Add(v string) {
	(*s) = append((*s), v)
}

func (s _String) Export(i int, encoder types.Encoder) {
	encoder.EncodeString(s[i])
}
