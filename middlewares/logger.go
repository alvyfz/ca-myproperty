package middlewares

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func LogMiddlewares(e *echo.Echo) {

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "[${time_rfc3339}], method=${method}, host=${host}, path=${path}, latecy_human=${latecy_human}, status=${status}\n",
	// 	//`[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latecy_human}` + "\n",
	// }))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "[${time_rfc3339}], method=${method}, host=${host}, path=${path}, latecy_human=${latecy_human}, status=${status}\n"}))
}
