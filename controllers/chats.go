package controller

import (
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func DeleteChatByID(c echo.Context) error {
	id := c.Param("id")
	service.DeleteChatByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteChatByID",
	})
}

func CreateChat(c echo.Context) error {
	var newChat model.Chat
	u := c.Get("user")
	claims := u.(jwt.MapClaims)
	userID := claims["userId"].(float64)
	if err := c.Bind(&newChat); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateChatController",
			"error":   err.Error(),
		})
	}
	newChat.UserID = uint(userID)
	newChat = service.CreateChat(newChat)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateChatController",
		"data":    newChat,
	})
}
