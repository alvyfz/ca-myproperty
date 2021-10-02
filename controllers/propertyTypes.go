package controller

import (
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllPropertyTypes(c echo.Context) error {
	propertyTypes := service.GetAllPropertyTypes()
	return c.JSON(http.StatusOK, propertyTypes)
}

func CreatePropertyType(c echo.Context) error {
	var newPropertyType model.PropertyType
	if err := c.Bind(&newPropertyType); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatePropertyTypeController",
			"error":   err.Error(),
		})
	}
	newPropertyType = service.CreatePropertyType(newPropertyType)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreatePropertyTypeController",
		"data":    newPropertyType,
	})
}

func UpdatePropertyType(c echo.Context) error {
	id := c.Param("id")

	var propertyType model.PropertyType
	if err := c.Bind(&propertyType); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatePropertyTypeController",
			"error":   err.Error(),
		})
	}
	service.UpdatePropertyType(id, propertyType)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update success",
		"data":    propertyType,
	})
}
