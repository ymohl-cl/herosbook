package main

import (
	"database/sql"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/ymohl-cl/herosbook/config"
	"github.com/ymohl-cl/herosbook/handlers"
	"github.com/ymohl-cl/herosbook/routes"
)

func main() {
	var err error
	var c *config.Config
	var db *sql.DB

	// Set validation structure
	govalidator.SetFieldsRequiredByDefault(true)

	// get configuration
	if c, err = config.New(); err != nil {
		panic(err)
	}

	// sql.DB instace
	if db, err = c.Psql.Init(); err != nil {
		panic(err)
	}

	// Handler instance and Echo instace
	h := handlers.New(db)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// defines routes
	routes.Descriptions(e, h)

	// Start server
	if err = e.StartTLS(c.API.Domain, c.API.Cert, c.API.Key); err != nil {
		panic(err)
	}
	return
}
