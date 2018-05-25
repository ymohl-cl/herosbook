package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/ymohl-cl/herosbook/cmd/api/handler"
	"github.com/ymohl-cl/herosbook/cmd/api/manager"
	"github.com/ymohl-cl/herosbook/pkg/cassandra"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
)

func main() {
	var err error
	var c *config
	var clientSQL postgres.Driver
	var clientCQL cassandra.Driver

	// Set validation structure
	govalidator.SetFieldsRequiredByDefault(true)

	// get configuration
	if c, err = configure(); err != nil {
		panic(err)
	}

	// sql.DB instance
	if clientSQL, err = c.SQL.New(); err != nil {
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

	// get Manager
	m := manager.New(clientSQL, clientCQL)

	// get handler
	h := handler.New(m)
	h.SetRoutes(e)
	h.SetMiddlewares(e)

	// Start server
	if err = e.StartTLS(c.SSL.Domain, c.SSL.Cert, c.SSL.Key); err != nil {
		panic(err)
	}
	return
}
