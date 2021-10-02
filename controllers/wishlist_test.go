package controller

// import (
// 	"ca-myproperty/config"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo"
// )

// func TestCreatWishlist(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPost, "/wishlists", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := CreateWishlist(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }

// func TestDeleteWishlistByID(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodDelete, "/wishlists/:id", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := DeleteWishlistByID(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }
