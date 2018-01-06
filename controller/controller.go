package controller

import (
	"database/sql"

	"github.com/gocql/gocql"
	"github.com/labstack/echo"
)

// Content is a object to controll the brain of application from one request
type Content interface {
	// JSONParser()
	JSONParser(echo.Context) (int, error)
	// IsValid check the content of request
	IsValid(int) (int, error)
	// Record save on bdd
	Record(*sql.DB, *gocql.Session) (int, error)
	// Provide the response to the client
	Response() ([]byte, int, error)
}
