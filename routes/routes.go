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
	e.GET("/users", controller.GetAllUsers)
	e.POST("/register", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser, m.JWTAuth)
	e.GET("/users/:id", controller.GetUserByID, m.JWTAuth)
	e.POST("/login", controller.UserLogin)

	// developer
	e.POST("/developers", controller.CreateDeveloper, m.JWTAuth)
	e.PUT("/developers/:id", controller.UpdateDeveloper, m.JWTAuth)
	e.GET("/developers/:id", controller.GetDeveloperByID, m.JWTAuth)

	//wishlist
	e.POST("/wishlists", controller.CreateWishlist, m.JWTAuth)
	e.DELETE("wishlists/:id", controller.DeleteWishlistByID, m.JWTAuth)

	//propertyType
	e.GET("/property-types", controller.GetAllPropertyTypes)
	e.POST("/property-types", controller.CreatePropertyType)
	e.PUT("/property-types", controller.CreatePropertyType)

	//property
	e.GET("/properties", controller.GetAllProperties)
	e.GET("/properties/:id", controller.GetPropertyByID)
	e.POST("/properties", controller.CreateProperty, m.JWTAuth)
	e.DELETE("properties", controller.DeletePropertyByID, m.JWTAuth)
	e.PUT("properties/:id", controller.UpdateProperty, m.JWTAuth)

	//transaction
	e.GET("/transactions", controller.GetAllTransactions, m.JWTAuth)
	e.GET("/transactions/:id", controller.GetTransactionByID, m.JWTAuth)
	e.POST("/transactions", controller.CreateTransaction, m.JWTAuth)
	e.DELETE("transactions", controller.DeleteTransactionByID, m.JWTAuth)
	e.PUT("transactions/:id", controller.UpdateTransaction, m.JWTAuth)

	//chat
	e.GET("/chats", controller.CreateChat, m.JWTAuth)
	e.DELETE("chats/:id", controller.DeleteChatByID, m.JWTAuth)

	// r := e.Group("/jwt")
	// r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	return e
}
