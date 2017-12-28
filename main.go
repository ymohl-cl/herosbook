package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ymohl-cl/herosbook/routes"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Descriptions(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	return
}
