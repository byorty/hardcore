package is

import "github.com/byorty/hardcore/types"

func EmptySlice(slice types.Slice) bool {
	return slice.Len() == 0
}

func NotEmptySlice(slice types.Slice) bool {
	return !EmptySlice(slice)
}
