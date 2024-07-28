package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	handlers "go-echo-playground/handlers/api/v1"
)

func RegisterRoutes(group *echo.Group) {
	group.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World from API v1!")
	})

	group.POST("/users", handlers.SaveUser)
	group.GET("/users/:id", handlers.GetUser)
	group.PUT("/users/:id", handlers.UpdateUser)
	group.DELETE("/users/:id", handlers.DeleteUser)
}
