package routes

import (
	controller "ca-myproperty/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	//routes User
	e.GET("/users", controller.GetAllUsers)
	e.POST("/register", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)

	return e
}
