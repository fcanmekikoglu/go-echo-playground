package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(group *echo.Group) {
	group.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World from API v1!")
	})
}
