package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/controller"
	"github.com/ymohl-cl/herosbook/controller/users"
)

// Handler contains drivers to execute query
type Handler struct {
	psql *sql.DB
	//mongo
}

// New : _
func New(p *sql.DB) *Handler {
	return &Handler{psql: p}
}

/*
** Home page
 */

// Home page, provide the books being read and books by themes and new books
func (h Handler) Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! \n It's a home page")
}

/*
** Users managment
 */

// GetUsers provide the users list
func (h Handler) GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "GetUsers provide the users list")
}

// GetUser provide the user defined by id
func (h Handler) GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "GetUser provide the user defined by id")
}

// CreateUser : _
func (h Handler) CreateUser(c echo.Context) error {
	var content controller.Content

	content = new(users.User)
	//r := new(users.ResponseInfosUser)

	if err := c.Bind(content); err != nil {
		return c.JSON(http.StatusPreconditionFailed, err)
	}

	if status, err := content.IsValid(); err != nil {
		return c.JSON(status, err)
	}

	if status, err := content.Record(h.psql); err != nil {
		return c.JSON(status, err)
	}

	r, status, err := content.Response()
	if err != nil {
		return c.JSON(status, err)
	}
	return c.JSON(http.StatusOK, r)
}

// EditUser : _
func (h Handler) EditUser(c echo.Context) error {
	return c.String(http.StatusOK, "EditUser : _")
}

// DeleteUser : _
func (h Handler) DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteUser : _")
}

/*
** Connect managment
 */

// Connect : _
func (h Handler) Connect(c echo.Context) error {
	return c.String(http.StatusOK, "Connect : _")
}

// Disconnect : _
func (h Handler) Disconnect(c echo.Context) error {
	return c.String(http.StatusOK, "Disconnect : _")
}

/*
** Header Book managment
 */

// GetBooks provide all books headers
func (h Handler) GetBooks(c echo.Context) error {
	return c.String(http.StatusOK, "GetBooks provide all books headers")
}

// GetBook provide the book header
func (h Handler) GetBook(c echo.Context) error {
	return c.String(http.StatusOK, "GetBook provide the book header")
}

// CreateBook : _
func (h Handler) CreateBook(c echo.Context) error {
	return c.String(http.StatusOK, "CreateBook : _")
}

// EditBook : _
func (h Handler) EditBook(c echo.Context) error {
	return c.String(http.StatusOK, "EditBook : _")
}

// DeleteBook : _
func (h Handler) DeleteBook(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteBook : _")
}

/*
** Content book managment
 */

// GetContentsBook provide the book contents
func (h Handler) GetContentsBook(c echo.Context) error {
	return c.String(http.StatusOK, "GetContentsBook provide the book contents")
}

// GetContentBook provide the book content define by idc
func (h Handler) GetContentBook(c echo.Context) error {
	return c.String(http.StatusOK, "GetContentBook provide the book content define by idc")
}

// CreateContentBook : _
func (h Handler) CreateContentBook(c echo.Context) error {
	return c.String(http.StatusOK, "CreateContentBook : _")
}

// EditContentBook : _
func (h Handler) EditContentBook(c echo.Context) error {
	return c.String(http.StatusOK, "EditContentBook : _")
}

// DeleteContentBook : _
func (h Handler) DeleteContentBook(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteContentBook : _")
}
