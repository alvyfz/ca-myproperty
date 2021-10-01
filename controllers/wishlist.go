package controller

import (
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"

	"github.com/labstack/echo"
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
	if err := c.Bind(&newWishlist); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateWishlistController",
			"error":   err.Error(),
		})
	}
	newWishlist = service.CreateWishlist(newWishlist)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateWishlistController",
		"data":    newWishlist,
	})
}