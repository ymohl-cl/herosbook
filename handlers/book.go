package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

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
