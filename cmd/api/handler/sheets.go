package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

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
