package middlewares

import (
	"ca-myproperty/config"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// func LogMiddlewares(e *echo.Echo) {
// 	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "[${time_rfc3339}], method=${method}, host=${host}, path=${path}, latecy_human=${latecy_human}, status=${status}\n"}))

// }
// // if wanna using logger with mongodb
func LogMiddlewares(next echo.HandlerFunc) echo.HandlerFunc {
	coll := config.DBLog.Database("myproperty").Collection("logger")

	return func(c echo.Context) error {
		request := c.Request()
		response := c.Response()
		log := bson.M{
			"time":   time.Now(),
			"method": request.Method,
			"path":   request.URL.Path,
			"status": response.Status,
		}

		err := next(c)
		coll.InsertOne(request.Context(), log)
		return err

	}
}
