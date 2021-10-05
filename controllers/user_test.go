package controller

// import (
// 	"ca-myproperty/config"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// )

// func TestGetAllUsers(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/users", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := GetAllUsers(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }

// func TestCreateUser(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPost, "/users", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := CreateUser(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }

// func TestUpdateUser(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPut, "/users:id", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := UpdateUser(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }

// func TestGetUserByID(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPut, "/users/:id", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := GetUserByID(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }
