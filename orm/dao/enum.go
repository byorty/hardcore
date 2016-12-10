package dao

import (
	"github.com/byorty/hardcore/types"
)

type EnumScannerImpl struct{}

func (e EnumScannerImpl) Find(scanner types.EnumScanner, slice types.Slice) {
	for i := 0; i < slice.Len(); i++ {
		if scanner.Eq(i) {
			scanner.Scan(i)
			break
		}
	}
}
