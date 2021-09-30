package middlewares

import (
	"ca-myproperty/config/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authotized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //token expired 1 jam
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(([]byte(constants.SECRET_JWT)))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}
