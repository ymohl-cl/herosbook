package main

import (
	"database/sql"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gocql/gocql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/ymohl-cl/herosbook/config"
	"github.com/ymohl-cl/herosbook/handlers"
	"github.com/ymohl-cl/herosbook/routes"
)

func Skip(c echo.Context) bool {
	fmt.Println("Skip ? ", c.Path())
	if c.Path() == "/" {
		fmt.Println("equal")
		return true
	}
	fmt.Println("not equal")
	return false
}

func Test(s string, c echo.Context) (bool, error) {
	if s == "api-key" {
		return true, nil
	}
	return false, nil
}

func main() {
	var err error
	var c *config.Config
	var dbSQL *sql.DB
	var dbCQL *gocql.Session

	// Set validation structure
	govalidator.SetFieldsRequiredByDefault(true)

	// get configuration
	if c, err = config.New(); err != nil {
		panic(err)
	}

	// sql.DB instance
	if dbSQL, err = c.Psql.Init(); err != nil {
		panic(err)
	}
	defer dbSQL.Close()

	// gocql.DB instance
	if dbCQL, err = c.Cass.Init(); err != nil {
		panic(err)
	}
	defer dbCQL.Close()

	// Handler instance and Echo instace
	h := handlers.New(dbSQL, dbCQL)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//	e.Use(middleware.KeyAuth(Test))

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:   Skip,
		KeyLookup: "header:Authorization",
		Validator: Test,
	}))

	// defines routes
	routes.Descriptions(e, h)

	// Start server
	if err = e.StartTLS(c.API.Domain, c.API.Cert, c.API.Key); err != nil {
		panic(err)
	}
	return
}
