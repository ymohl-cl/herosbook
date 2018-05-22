package manager

import (
	"database/sql"

	"github.com/ymohl-cl/herosbook/pkg/cassandra"
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"

	"github.com/gocql/gocql"
	"github.com/labstack/echo"
)

// Manager implement type to manage all queries from client
type Manager interface {
	Connect() (err error)
	Disconnect() (err error)
	CreateUser(user model.User) (object interface{}, err error)
	EditUser() (err error)
	DeleteUser() (err error)
}

type manage struct {
	clientSQL postgres.Driver
	clientCQL cassandra.Driver
}

// Operater type define interface to manage the specific query
//type Operater interface {
	// JSONParser parse the body request to get the relative object
	//JSONParser(echo.Context) (int, error)
	// IsValid check the marshaller's object and valid the legitimate query
	//IsValid(int) (int, error)
	// Execute query on drivers
	//Record(*sql.DB, *gocql.Session) (int, error)
	// Provide the response to the client
	//Response() ([]byte, int, error)
}

// New provide the Manager interface to execute the action relative at client's query
func New(clientSQL postgres.Driver, clientCQL cassandra.Driver) (m Manager) {
	m = &manage{
		clientCQL: clientCQL,
		clientSQL: clientSQL,
	}
	return m
}
