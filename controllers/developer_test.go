package controller

import (
	"ca-myproperty/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreateDeveloper(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/developers", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := CreateDeveloper(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestUpdateDeveloper(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/developers/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := UpdateDeveloper(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestGetDeveloperByID(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/users/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := GetDeveloperByID(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
