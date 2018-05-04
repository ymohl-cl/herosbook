package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/ymohl-cl/herosbook/pkg/cassandra"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
)

func main() {
	var err error
	var c config
	var clientSQL postgres.Driver
	var clientCQL cassandra.Driver

	// Set validation structure
	govalidator.SetFieldsRequiredByDefault(true)

	// get configuration
	if c, err = configure(); err != nil {
		panic(err)
	}

	// sql.DB instance
	if clientSQL, err = c.SQL.Client(); err != nil {
		panic(err)
	}
	defer clientSQL.Close()

	// gocql.DB instance
	if clientCQL, err = c.CQL.Client(); err != nil {
		panic(err)
	}
	defer clientCQL.Close()

	// server http instance
	e := echo.New()
	setRoutes(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:   Skip,
		KeyLookup: "header:Authorization",
		Validator: Test,
	}))

	// Start server
	if err = e.StartTLS(c.SSL.Domain, c.SSL.Cert, c.SSL.Key); err != nil {
		panic(err)
	}
	return
}

/*
** Shema

cmd/
	api/
		main.go
		config.go
		server/
			server.go
			router.go
		handler.go
			handler.go // perform queries responses
			object1 // handle the query question
			object2 // handle the query question
			object3 // handle the query question
			object4 // handle the query question
		middleware/
			middleware.go
		manager
			manager.go // with interface
			object1 // implement manager
			object2 // implement manager
			object3 // implement manager
			object4 // implement manager
	others1/
	others2/

pkg/
	elastic client/
		object1
		object2
		object3
	cassandra client/
		object1
		object2
		object3
	models/
		object1
		object2
		object3
		object4
	httpstatus/
		httpstatus.go // interpret error
*/
