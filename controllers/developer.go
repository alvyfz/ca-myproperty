package controller

import (
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"

	"github.com/labstack/echo"
)

func CreateDeveloper(c echo.Context) error {
	var newDeveloper model.Developer
	if err := c.Bind(&newDeveloper); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateDeveloperController",
			"error":   err.Error(),
		})
	}
	newDeveloper = service.CreateDeveloper(newDeveloper)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateDeveloperController",
		"data":    newDeveloper,
	})
}

func UpdateDeveloper(c echo.Context) error {
	id := c.Param("id")

	var Developer model.Developer
	if err := c.Bind(&Developer); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateDeveloperController",
			"error":   err.Error(),
		})
	}
	service.UpdateDeveloper(id, Developer)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update success",
		"data":    Developer,
	})
}
