package current

import (
	_ "github.com/lib/pq"
	"github.com/byorty/hardcore/server"
	"github.com/byorty/hardcore/test/configs"
	"github.com/byorty/hardcore/orm/db"
	"github.com/byorty/hardcore/pool"
)

func init()  {
	server.Environment().
		SetProjectName("Test").
		SetRouter(configs.Router())

	pool.DB().
		Add(
			"default",
			db.NewSqlDB("postgres://byorty:MK99rc@localhost:5432/hardcore?sslmode=disable&client_encoding=utf-8"),
		)

}
