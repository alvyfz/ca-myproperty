package controller

import (
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	users := service.GetAllUsers()
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}
	newUser = service.CreateUser(newUser)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateUserController",
		"data":    newUser,
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}
	service.UpdateUser(id, user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update success",
		"data":    user,
	})
}