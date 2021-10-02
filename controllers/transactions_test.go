package controller

import (
	"ca-myproperty/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetAlTransactions(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/transactions", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := GetAllTransactions(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestCreatTransaction(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/transactions", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := CreateTransaction(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestUpdateTransaction(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/transactions:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := UpdateTransaction(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestGetTransactionByID(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/transactions/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := GetPropertyByID(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
func TestDeleteTransactionByID(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/transactions/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := DeleteTransactionByID(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
