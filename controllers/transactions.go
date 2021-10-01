package controller

import (
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllTransactions(c echo.Context) error {
	transactions := service.GetAllProperties()
	return c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c echo.Context) error {
	var newTransaction model.Transaction
	if err := c.Bind(&newTransaction); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateTransactionController",
			"error":   err.Error(),
		})
	}
	newTransaction = service.CreateTransaction(newTransaction)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateTransactionController",
		"data":    newTransaction,
	})
}

func UpdateTransaction(c echo.Context) error {
	id := c.Param("id")

	var transaction model.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateTransactionController",
			"error":   err.Error(),
		})
	}
	service.UpdateTransaction(id, transaction)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update success",
		"data":    transaction,
	})
}
func GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	transaction := service.GetTransactionByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetTransactionByIDController",
		"data":    transaction,
	})
}

func DeleteTransactionByID(c echo.Context) error {
	id := c.Param("id")
	service.DeleteTransactionByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteTransactionByID",
	})
}
