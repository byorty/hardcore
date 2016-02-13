package utils

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/pool"
)

func InitDAOs(daos ...types.ModelDAO)  {
	for _, dao := range daos {
		db := pool.DB().ByDAO(dao)
		dao.AutoInit(db)
		dao.Init(db)
	}
}
