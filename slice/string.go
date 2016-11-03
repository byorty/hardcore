package slice

import "github.com/byorty/hardcore/types"

type Strings []string

func NewStrings() types.StringSlice {
	return make(Strings, 0)
}

func NewStringsBy(arr []string) types.StringSlice {
	return Strings(arr)
}

func (s Strings) Len() int {
	return len(s)
}

func (s Strings) Less(x, y int) bool {
	return s[x] < s[y]
}

func (s Strings) Swap(x, y int) {
	s[x], s[y] = s[y], s[x]
}

func (s Strings) GetRaw(x int) interface{} {
	return s.Get(x)
}

func (s Strings) Get(x int) string {
	return s[x]
}

func (s Strings) ExportItem(i int, encoder types.Encoder) {
	encoder.EncodeString(s[i])
}
