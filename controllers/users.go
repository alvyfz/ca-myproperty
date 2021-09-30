package controller

import (
	database "ca-myproperty/lib/databases"
	model "ca-myproperty/models"
	service "ca-myproperty/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
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

func LoginUser(c echo.Context) error {
	var user model.User
	c.Bind(&user)

	users, err := database.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})

}

func GetDetailUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := database.GetDetailUser((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

// func UserLogin(c echo.Context) error {
// 	user := model.User{}
// 	if err := c.Bind(&user); err != nil {
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}

// 	isValid := database.IsValid(user.Email, user.Password)

// 	if !isValid {
// 		return c.String(http.StatusBadRequest, "email atau password salah")
// 	}

// 	claims := jwt.MapClaims{}
// 	claims["userId"] = user.Email
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte("asdasdasd"))
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.String(http.StatusOK, tokenString)
// }
// func BasicAuth(c echo.Context) error {
// 	email := c.Get("email")
// 	return c.String(http.StatusOK, fmt.Sprintf("selamat datang %s", email))
// }
