package controller

import (
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func DeleteWishlistByID(c echo.Context) error {
	id := c.Param("id")
	service.DeleteWishlistByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeletewishlistByID",
	})
}

func CreateWishlist(c echo.Context) error {
	var newWishlist model.Wishlist
	u := c.Get("user")
	claims := u.(jwt.MapClaims)
	userID := claims["userId"].(float64)
	if err := c.Bind(&newWishlist); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateWishlistController",
			"error":   err.Error(),
		})
	}
	newWishlist.UserId = uint(userID)
	newWishlist = service.CreateWishlist(newWishlist)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateWishlistController",
		"data":    newWishlist,
	})
}
