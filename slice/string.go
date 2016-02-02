package slice

import "github.com/byorty/hardcore/types"

type StringsImpl []string

func NewStrings() types.StringSlice {
	return make(StringsImpl, 0)
}

func NewStringsBy(arr []string) types.StringSlice {
	return StringsImpl(arr)
}

func (s StringsImpl) Len() int {
	return len(s)
}

func (s StringsImpl) Less(x, y int) bool {
	return s[x] < s[y]
}

func (s StringsImpl) Swap(x, y int) {
	s[x], s[y] = s[y], s[x]
}

func (s StringsImpl) GetRaw(x int) interface{} {
	return s.Get(x)
}

func (s StringsImpl) Get(x int) string {
	return s[x]
}
