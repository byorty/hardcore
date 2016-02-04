package main

import (
	"github.com/byorty/hardcore/server"
	_ "github.com/byorty/hardcore/test/configs/current"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1000)
	app := server.New()
	app.SetEnvironment(server.Environment())
	app.Serve()
}
