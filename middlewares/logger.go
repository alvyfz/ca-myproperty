package middlewares

import (
	"ca-myproperty/config"
	model "ca-myproperty/models"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

// func LogMiddlewares(e *echo.Echo) {
// 	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "[${time_rfc3339}], method=${method}, host=${host}, path=${path}, latecy_human=${latecy_human}, status=${status}\n"}))
// }

// if wanna using logger with mongodb
func LogMiddlewares(next echo.HandlerFunc) echo.HandlerFunc {
	coll := config.DBLog.Database("myproperty").Collection("logger")

	return func(c echo.Context) error {
		data := new(model.User)
		if c.Request().Method != http.MethodGet {
			if err := c.Bind(&data); err != nil {
				fmt.Println(err)
			}
		}

		log := bson.M{
			"time":    time.Now(),
			"method":  c.Request().Method,
			"path":    c.Path(),
			"request": data,
		}
		response := next(c)
		log["response"] = c.Response().Status
		coll.InsertOne(c.Request().Context(), log)
		fmt.Println(log)
		return response
	}
}
