package main

import (
	"ca-myproperty/config"
	m "ca-myproperty/middlewares"
	"ca-myproperty/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
