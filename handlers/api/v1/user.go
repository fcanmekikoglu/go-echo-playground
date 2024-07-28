package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SaveUser(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{"message": "User saved"})
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{"message": "Get user " + id})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusCreated, map[string]string{"message": "Update user " + id})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusNoContent, map[string]string{"message": "Delete user " + id})
}
