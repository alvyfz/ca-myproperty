package main

import (
	"ca-myproperty/config"
	"ca-myproperty/middlewares"
	"ca-myproperty/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.InitDB()
	config.InitLog()
	e := routes.New()
	e.Use(middlewares.LogMiddlewares)
	// middlewares.LogMiddlewares(e)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))
}
