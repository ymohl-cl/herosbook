package model

import (
	"database/sql"

	"github.com/gocql/gocql"
)

// liste of tag
const (
	Create = iota
	Update
	Delete
)

// Account manage create edit and delete user
type Account struct {
	user users.User
	Rec  func(*sql.DB, *gocql.Session) (int, error)
	Resp func() ([]byte, int, error)
}

// Record is implementation interface controller
func (a *Account) Record(psql *sql.DB, cql *gocql.Session) (int, error) {
	return a.Rec(psql, cql)
}

// Response is implementation interface controller
func (a *Account) Response() ([]byte, int, error) {
	return a.Resp()
}
