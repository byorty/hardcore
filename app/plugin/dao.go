package plugin

import (
	"github.com/byorty/hardcore/pool"
	"github.com/byorty/hardcore/scope"
)

type DAOImpl struct{}

func (d *DAOImpl) Run() {
	for _, dao := range scope.App().GetDAOs() {
		db := pool.DB().ByDAO(dao)
		dao.AutoInit(db)
		dao.Init(db)
	}
}
