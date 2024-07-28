package main

import (
	"go-echo-playground/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "client/dist")

	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
