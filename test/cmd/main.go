package main

import (
	"github.com/byorty/hardcore/server"
	_ "github.com/byorty/hardcore/test/configs/current"
	"runtime"
	"github.com/byorty/hardcore/env"
)

func main() {
	runtime.GOMAXPROCS(100000)
	app := server.New()
	app.SetEnv(env.Me())
	app.Serve()
}
