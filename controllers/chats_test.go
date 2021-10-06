package controller

import (
	"ca-myproperty/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func TestCreateChat(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/chats", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := CreateChat(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestDeleteChatByID(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/chats/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := DeleteChatByID(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
