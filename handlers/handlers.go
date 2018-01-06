package handlers

import (
	"database/sql"

	"github.com/gocql/gocql"
	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/controller"
	"github.com/ymohl-cl/herosbook/controller/account"
)

// Handler contains drivers to execute query
type Handler struct {
	psql      *sql.DB
	cassandra *gocql.Session
}

// New : _
func New(p *sql.DB, c *gocql.Session) *Handler {
	return &Handler{psql: p, cassandra: c}
}

func (h Handler) setters(c echo.Context, control controller.Content) (int, []byte, error) {
	var err error
	var status int
	var r []byte

	// parse json
	if status, err = control.JSONParser(c); err != nil {
		return status, nil, err
	}
	// check validity
	if status, err = control.IsValid(account.Create); err != nil {
		return status, nil, err
	}
	// record
	if status, err = control.Record(h.psql, h.cassandra); err != nil {
		return status, nil, err
	}
	// get response
	if r, status, err = control.Response(); err != nil {
		return status, nil, err
	}

	return 0, r, nil
}

func (h Handler) getters(c echo.Context, control controller.Content) (int, interface{}, error) {
	var err error
	var status int
	var r interface{}

	// parse json
	if status, err = control.JSONParser(c); err != nil {
		return status, nil, err
	}
	// check validity
	if status, err = control.IsValid(account.Create); err != nil {
		return status, nil, err
	}
	// get response
	if r, status, err = control.Response(); err != nil {
		return status, nil, err
	}

	return 0, r, nil
}
