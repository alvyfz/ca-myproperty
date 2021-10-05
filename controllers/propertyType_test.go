package controller

// import (
// 	"ca-myproperty/config"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// )

// func TestGetAllPropertyTypes(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/property-types", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := GetAllPropertyTypes(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }

// func TestCreatPropertyType(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPost, "/property-types", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := CreatePropertyType(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }

// func TestUpdatePropertyType(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPut, "/property-types:id", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := UpdatePropertyType(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }
