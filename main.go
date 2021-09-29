package main

import (
	"ca-myproperty/config"
	"ca-myproperty/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}