package dao

import "github.com/byorty/hardcore/types"

type EnumScannerImpl struct {
	id  interface{}
	dao types.EnumDAO
}

func NewEnumScanner(id interface{}, dao types.EnumDAO) types.EnumScanner {
	return &EnumScannerImpl{id, dao}
}

func (e EnumScannerImpl) One(enum types.Named) {
	for _, named := range e.dao.GetList() {
		if e.dao.Eq(named, e.id) {
			e.dao.Scan(named, enum)
		}
	}
}
