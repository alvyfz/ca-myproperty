package controller

import (
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllProperties(c echo.Context) error {
	properties := service.GetAllProperties()
	return c.JSON(http.StatusOK, properties)
}

func CreateProperty(c echo.Context) error {
	var newProperty model.Property
	if err := c.Bind(&newProperty); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatePropertyController",
			"error":   err.Error(),
		})
	}
	newProperty = service.CreateProperty(newProperty)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreatePropertyController",
		"data":    newProperty,
	})
}

func UpdateProperty(c echo.Context) error {
	id := c.Param("id")

	var Property model.Property
	if err := c.Bind(&Property); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatePropertyController",
			"error":   err.Error(),
		})
	}
	service.UpdateProperty(id, Property)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update success",
		"data":    Property,
	})
}
func GetPropertyByID(c echo.Context) error {
	id := c.Param("id")
	Property := service.GetPropertyByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetPropertyByIDController",
		"data":    Property,
	})
}

func DeletePropertyByID(c echo.Context) error {
	id := c.Param("id")
	service.DeletePropertyByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeletePropertyByID",
	})
}
