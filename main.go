package main

import (
	"ca-myproperty/config"
	"ca-myproperty/middlewares"
	"ca-myproperty/routes"

	"github.com/labstack/echo/v4/middleware"
)

// func init() {
// 	viper.SetConfigFile(`config.json`)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// 	if viper.GetBool(`debug`) {
// 		log.Println("Service RUN on DEBUG mode")
// 	}
// }
func main() {

	// configDB := config.ConfigDB{
	// 	DB_Username: viper.GetString(`database.user`),
	// 	DB_Password: viper.GetString(`database.pass`),
	// 	DB_Host:     viper.GetString(`database.host`),
	// 	DB_Port:     viper.GetString(`database.port`),
	// 	DB_Name:     viper.GetString(`database.name`),
	// }
	config.InitDB()
	config.InitLog()
	e := routes.New()
	e.Use(middlewares.LogMiddlewares)
	// middlewares.LogMiddlewares(e)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))
}
