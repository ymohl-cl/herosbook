package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
** Home page
 */

// Home page, provide the books being read and books by themes and new books
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! \n It's a home page")
}

/*
** Users managment
 */

// GetUsers provide the users list
func GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "GetUsers provide the users list")
}

// GetUser provide the user defined by id
func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "GetUser provide the user defined by id")
}

// CreateUser : _
func CreateUser(c echo.Context) error {
	return c.String(http.StatusOK, "CreateUser : _")
}

// EditUser : _
func EditUser(c echo.Context) error {
	return c.String(http.StatusOK, "EditUser : _")
}

// DeleteUser : _
func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteUser : _")
}

/*
** Connect managment
 */

// Connect : _
func Connect(c echo.Context) error {
	return c.String(http.StatusOK, "Connect : _")
}

// Disconnect : _
func Disconnect(c echo.Context) error {
	return c.String(http.StatusOK, "Disconnect : _")
}

/*
** Header Book managment
 */

// GetBooks provide all books headers
func GetBooks(c echo.Context) error {
	return c.String(http.StatusOK, "GetBooks provide all books headers")
}

// GetBook provide the book header
func GetBook(c echo.Context) error {
	return c.String(http.StatusOK, "GetBook provide the book header")
}

// CreateBook : _
func CreateBook(c echo.Context) error {
	return c.String(http.StatusOK, "CreateBook : _")
}

// EditBook : _
func EditBook(c echo.Context) error {
	return c.String(http.StatusOK, "EditBook : _")
}

// DeleteBook : _
func DeleteBook(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteBook : _")
}

/*
** Content book managment
 */

// GetContentsBook provide the book contents
func GetContentsBook(c echo.Context) error {
	return c.String(http.StatusOK, "GetContentsBook provide the book contents")
}

// GetContentBook provide the book content define by idc
func GetContentBook(c echo.Context) error {
	return c.String(http.StatusOK, "GetContentBook provide the book content define by idc")
}

// CreateContentBook : _
func CreateContentBook(c echo.Context) error {
	return c.String(http.StatusOK, "CreateContentBook : _")
}

// EditContentBook : _
func EditContentBook(c echo.Context) error {
	return c.String(http.StatusOK, "EditContentBook : _")
}

// DeleteContentBook : _
func DeleteContentBook(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteContentBook : _")
}
