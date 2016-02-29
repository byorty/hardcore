package pool

import (
	"github.com/byorty/hardcore/types"
)

type DBPoolImpl map[string]types.DB

var (
	dbPool = make(DBPoolImpl)
)

func DB() types.DBPool {
	return dbPool
}

func (p DBPoolImpl) Add(name string, db types.DB) types.DBPool {
	p[name] = db
	return p
}

func (p DBPoolImpl) ByDAO(impl types.ModelDAO) types.DB {
	return p.ByName(impl.GetDB())
}

func (p DBPoolImpl) ByName(name string) types.DB {
	if db, ok := p[name]; ok {
		return db
	} else {
		return nil
	}
}
