package routes

import (
	controller "ca-myproperty/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controller.GetAllUsers)
	e.POST("/register", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)

	return e
}
