package routes

import (
	"ca-myproperty/config/constants"
	controller "ca-myproperty/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// e.GET("/", controller.BasicAuth, middlewares.JWTAuthMiddleware)
	//routes User
	e.GET("/users", controller.GetAllUsers)
	e.POST("/register", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.POST("/login", controller.LoginUser)

	// developer
	e.POST("/developer", controller.CreateDeveloper)
	e.PUT("/developers/:id", controller.UpdateDeveloper)

	// eAuth := e.Group("")
	// eAuth.Use(middleware.BasicAuth(middlewares.BasicAuthDB))
	// eAuth.PUT("users/:id", controller.UpdateUser)
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("users/:id", controller.GetDetailUser)

	return e
}
