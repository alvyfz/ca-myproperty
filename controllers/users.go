package controller

import (
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	users := service.GetAllUsers()
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}

	newUser = service.CreateUser(newUser)
	newUser.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateUserController",
		"data":    newUser,
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}
	service.UpdateUser(id, user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update success",
		"data":    user,
	})
}
func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user := service.GetUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetUserByIDController",
		"data":    user,
	})
}

func UserLogin(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := service.IsValid(user.Email, user.Password)

	if err != nil {
		return c.String(http.StatusBadRequest, "email atau password salah")
	}

	claims := jwt.MapClaims{}
	claims["userId"] = user.ID
	claims["userEmail"] = user.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("asdasdasd"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "Login success",
		"data":   tokenString,
	})
}
