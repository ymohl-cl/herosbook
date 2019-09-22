package app

import (
	"net/http"

	"github.com/labstack/echo"
)

// Pong model to describe the response http to a ping method
type Pong struct {
	Pong bool `json:"pong"`
}

// Ping method http GET
func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &Pong{Pong: true})
}
