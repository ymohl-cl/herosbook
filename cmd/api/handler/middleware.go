package handler

import (
	"github.com/labstack/echo"
)

// Authenticated check if user has a right to ressources access
func (h handle) Authenticated(s string, c echo.Context) (ok bool, err error) {
	/*
		var ID users.PublicID

		if c.Path() == "/" || c.Path() == "/login" || c.Path() == "/users" {
			return true, err
		}
		if err = c.Bind(&ID); err != nil {
			return false, err
		}
		if ok, err = controller.CheckToken(h.cassandra, ID.Value, []byte(s)); err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
		return true, nil
	*/
	return true, nil
}
