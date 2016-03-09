package main

import (
	"github.com/byorty/hardcore/app"
	_ "github.com/byorty/hardcore/test/configs/current"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(100000)
	server := app.NewHttpServer()
	server.Run()
}
