package routes

import (
	v1 "go-echo-playground/routes/api/v1"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")
	v1.RegisterRoutes(apiV1)
}
