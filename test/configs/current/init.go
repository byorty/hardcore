package current

import (
	_ "github.com/lib/pq"
	"github.com/byorty/hardcore/test/configs"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/pool"
	"github.com/byorty/hardcore/env"
	"github.com/byorty/hardcore/log"
)

func init()  {
	env.Me().
		SetProjectName("Test").
		SetRouter(configs.Router()).
		SetLogger(log.NewDefaultLogger(log.ERROR))

	pool.DB().
		Add(
			"default",
			db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/hardcore?sslmode=disable&client_encoding=utf-8"),
		)

}
