package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bearer token-blablala=
		authorizationFromHeader := c.Request().Header.Get("authorization")
		if authorizationFromHeader == "" {
			return c.String(http.StatusForbidden, "token salah")
		}

		// Replace Bearer, Get token-blablala=
		tokenString := strings.ReplaceAll(authorizationFromHeader, "Bearer ", "")

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("asdasdasd"), nil
		})
		if err != nil {
			return c.String(http.StatusForbidden, "token salah")
		}

		c.Set("user", claims)
		return next(c)
	}
}
