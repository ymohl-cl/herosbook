package controller

import (
	"database/sql"
)

// Content is a object to controll the brain of application from one request
type Content interface {
	// IsValid check the content of request
	IsValid() (int, error)
	// Record save on bdd
	Record(psql *sql.DB) (int, error)
	// Provide the response to the client
	Response() (interface{}, int, error)
}
