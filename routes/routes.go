package routes

import (
	controller "ca-myproperty/controllers"
	m "ca-myproperty/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// e.GET("/", controller.BasicAuth, middlewares.JWTAuthMiddleware)
	//routes User
	// e.GET("/user", controller.GetAllUsers)
	e.POST("/register", controller.CreateUser)
	e.PUT("/user/:id", controller.UpdateUser, m.JWTAuth)
	e.GET("/user/:id", controller.GetUserByID, m.JWTAuth)
	e.POST("/login", controller.UserLogin)

	// developer
	e.POST("/developer", controller.CreateDeveloper, m.JWTAuth)
	e.PUT("/developer/:id", controller.UpdateDeveloper, m.JWTAuth)
	e.GET("/developer/:id", controller.GetDeveloperByID, m.JWTAuth)

	//wishlist
	e.POST("/wishlist", controller.CreateWishlist, m.JWTAuth)
	e.DELETE("wishlist/:id", controller.DeleteWishlistByID, m.JWTAuth)

	//propertyType
	e.GET("/property-type", controller.GetAllPropertyTypes)
	e.POST("/property-type", controller.CreatePropertyType)
	e.PUT("/property-type", controller.CreatePropertyType)

	//property
	e.GET("/property", controller.GetAllProperties)
	e.GET("/property/:id", controller.GetPropertyByID)
	e.POST("/property", controller.CreateProperty, m.JWTAuth)
	e.DELETE("property", controller.DeletePropertyByID, m.JWTAuth)
	e.PUT("property/:id", controller.UpdateProperty, m.JWTAuth)

	//transaction
	e.GET("/transaction", controller.GetAllTransactions, m.JWTAuth)
	e.GET("/transaction/:id", controller.GetTransactionByID, m.JWTAuth)
	e.POST("/transaction", controller.CreateTransaction, m.JWTAuth)
	e.DELETE("transaction", controller.DeleteTransactionByID, m.JWTAuth)
	e.PUT("transaction/:id", controller.UpdateTransaction, m.JWTAuth)

	//chat
	e.GET("/chat", controller.CreateChat, m.JWTAuth)
	e.DELETE("chat/:id", controller.DeleteChatByID, m.JWTAuth)

	//location
	e.GET("/location", controller.CheckLocationByIP)

	// r := e.Group("/jwt")
	// r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	return e
}
