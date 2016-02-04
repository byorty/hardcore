package main

import (
	"github.com/byorty/hardcore/server"
	_ "github.com/byorty/hardcore/test/configs/current"
)

func main() {
	app := server.New()
	app.SetEnvironment(server.Environment())
	app.Serve()
}
