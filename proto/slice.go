package proto

import "github.com/byorty/hardcore/types"

type protoSlice []types.ProtoProperty

func (s protoSlice) HasNum(i int) bool {
	return i < len(s)
}

func (s protoSlice) Get(i int) types.ProtoProperty {
	if s.HasNum(i) {
		return s[i]
	} else {
		return nil
	}
}
