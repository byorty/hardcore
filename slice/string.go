package slice

type Strings []string

func NewStrings() Strings {
	return make(Strings, 0)
}

func NewStringsBy(arr []string) Strings {
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
