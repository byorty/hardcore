package plugin

import (
	"github.com/byorty/hardcore/pool"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/types"
)

func NewDAO() types.ApplicationPlugin {
	return NewDAOByName(scope.DefaultName)
}

func NewDAOByName(name string) types.ApplicationPlugin {
	return newByName(new(DAOImpl), name)
}

type DAOImpl struct{
	BaseImpl
}

func (d *DAOImpl) Run() {
	for _, dao := range scope.AppByName(d.name).GetDAOs() {
		db := pool.DB().ByDAO(dao)
		dao.AutoInit(db)
		dao.Init(db)
	}
}
