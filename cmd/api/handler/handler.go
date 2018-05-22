package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/ymohl-cl/herosbook/cmd/api/manager"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

const (
	errorInvalidJSON     = "invalid format json to the request"
	errorResponseBuilder = "invalid object to build respone"
)

// Handler receipt all queries client, build and run the relative manager
type Handler interface {
	SetRoutes(e *echo.Echo) (err error)
	SetMiddlewares(e *echo.Echo) (err error)
}

type handle struct {
	m manager.Manager
}

// New provide the Handler interface to receipt the query
func New(m manager.Manager) (h Handler) {
	h = &handle{
		m: m,
	}
	return h
}

// SetRoutes define the content api
func (h handle) SetRoutes(e *echo.Echo) (err error) {
	if e == nil {
		err = errors.New("server http not initialized")
		return err
	}

	// Home page
	e.GET("/", h.Home)

	// account managment
	e.POST("/login", h.Connect)
	e.DELETE("/login", h.Disconnect)
	e.POST("/users", h.CreateUser)
	e.PUT("/users/:id", h.EditUser)
	e.DELETE("/users/:id", h.DeleteUser)

	/*
		// book managment
		e.GET("/books", h.GetBooks)
		e.GET("/books/:id", h.GetBook)
		e.POST("/books", h.CreateBook)
		e.PUT("/books/:id", h.EditBook)
		e.DELETE("/books/:id", h.DeleteBook)

		// content book managment
		e.GET("/books/:id/contents", h.GetContentsBook)
		e.GET("/books/:id/contents/:idc", h.GetContentBook)
		e.POST("/books/:id/contents", h.CreateContentBook)
		e.PUT("/books/:id/contents/:idc", h.EditContentBook)
		e.DELETE("/books/:id/contents/:idc", h.DeleteContentBook)
	*/
	return nil
}

// SetMiddlewares configures the set of middlewares on the http server
func (h handle) SetMiddlewares(e *echo.Echo) (err error) {
	if e == nil {
		err = errors.New("server http not initialized")
		return err
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:   nil,
		KeyLookup: "header:Authorization",
		Validator: h.Authenticated,
	}))
	return nil
}

func (h handle) jsonParser(c echo.Context, m model.Model) (interface{}, error) {
	var err error

	if err = c.Bind(dest); err != nil {
		return nil, errors.New(errorInvalidJSON + " expected error: " + err.Error())
	}

	if err = dest.Validate(); err != nil {
		return nil, errors.New(errorInvalidJSON + " expected error: " + err.Error())
	}
	return dest, nil
}

func (h handle) jsonBuilder(object interface{}) ([]byte, error) {
	var response []byte
	var err error

	if response, err = json.Marshal(object); err != nil {
		return nil, errors.New(errorResponseBuilder + " expected error " + err.Error())
	}
	return response, nil
}

func (h handle) getHTTPStatus(err error) int {
	var httpStatus int

	if strings.Contains(err.Error(), errorInvalidJSON) {
		return http.StatusBadRequest
	}
	if strings.Contains(err.Error(), errorResponseBuilder) {
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}

/*
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
*/
