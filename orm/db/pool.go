package db

import (
	"github.com/byorty/hardcore/types"
)

type PoolImpl map[string]types.DB

var (
	pool = make(PoolImpl)
)

func Pool() types.Pool {
	return pool
}

func (p PoolImpl) Add(name string, db types.DB) types.Pool {
	if db == nil {
//		logger.Err(`pool - can't add disconnected db`)
	} else {
		p[name] = db
	}
	return p
}

func (p PoolImpl) ByDAO(impl types.DAO) types.DB {
	return p.ByName(impl.GetDB())
}

func (p PoolImpl) ByName(name string) types.DB {
	if db, ok := p[name]; ok {
		return db
	} else {
		return nil
	}
}
