package current

import (
	"github.com/byorty/hardcore/log"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/pool"
	"github.com/byorty/hardcore/scope"
	"github.com/byorty/hardcore/test/configs"
	_ "github.com/lib/pq"
)

func init() {
	scope.App().
		SetProjectName("Test").
		SetRouter(configs.Routers).
		SetLogger(log.NewDefaultLogger(log.ERROR))

	pool.DB().
		Add(
			"default",
			db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/hardcore?sslmode=disable&client_encoding=utf-8"),
		)

}
