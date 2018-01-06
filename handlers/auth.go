package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
** Connect managment
 */

// Connect : _
func (h Handler) Connect(c echo.Context) error {
	c.Set("un", nil)
	return c.String(http.StatusOK, "Connect : _")
}

// Disconnect : _
func (h Handler) Disconnect(c echo.Context) error {
	return c.String(http.StatusOK, "Disconnect : _")
}
