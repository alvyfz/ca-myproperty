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
	e.GET("/users/:id", controller.GetUserByID)
	e.POST("/login", controller.UserLogin)

	// developer
	e.POST("/developers", controller.CreateDeveloper)
	e.PUT("/developers/:id", controller.UpdateDeveloper)
	e.GET("/developers/:id", controller.GetDeveloperByID)

	//wishlist
	e.GET("/wishlists", controller.CreateWishlist)
	e.DELETE("wishlists/:id", controller.DeleteWishlistByID)

	//propertyType
	e.GET("/property-Types", controller.GetAllPropertyTypes)
	e.POST("/property-Types", controller.CreatePropertyType)
	e.PUT("/property-Types", controller.CreatePropertyType)

	//property
	e.GET("/properties", controller.GetAllProperties)
	e.GET("/properties/:id", controller.GetPropertyByID)
	e.POST("/properties", controller.CreateProperty)
	e.DELETE("properties", controller.DeletePropertyByID)
	e.PUT("properties/:id", controller.UpdateProperty)

	//transaction
	e.GET("/transactions", controller.GetAllTransactions)
	e.GET("/transactions/:id", controller.GetTransactionByID)
	e.POST("/trannsactions", controller.CreateTransaction)
	e.DELETE("transactions", controller.DeleteTransactionByID)
	e.PUT("transactions/:id", controller.UpdateTransaction)

	//chat
	e.GET("/chats", controller.CreateChat)
	e.DELETE("chats/:id", controller.DeleteChatByID)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	return e
}
