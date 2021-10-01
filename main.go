package main

import (
	"ca-myproperty/config"
	"ca-myproperty/middlewares"
	"ca-myproperty/routes"
)

func main() {
	config.InitDB()
	// config.InitLog()
	e := routes.New()
	e.Use(middlewares.LogMiddlewares)
	// m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
