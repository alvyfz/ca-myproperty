package middlewares

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func LogMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "[${time_rfc3339}], method=${method}, host=${host}, path=${path}, latecy_human=${latecy_human}, status=${status}\n"}))
}
