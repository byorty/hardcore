package db

import (
	"github.com/byorty/hardcore/logger"
	"github.com/byorty/hardcore/types"
)

type dbs map[string]types.DB

var (
	pool = make(dbs)
)

func Pool() types.Pool {
	return pool
}

func (d dbs) Add(name string, db types.DB) types.Pool {
	if db == nil {
		logger.Err(`pool - can't add disconnected db`)
	} else {
		d[name] = db
	}
	return d
}

func (d dbs) ByDAO(impl types.DAO) types.DB {
	return d.ByName(impl.GetDB())
}

func (d dbs) ByName(name string) types.DB {
	if db, ok := d[name]; ok {
		return db
	} else {
		return nil
	}
}
