package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/controller"
	"github.com/ymohl-cl/herosbook/controller/account"
)

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
	var err error
	var status int
	var r []byte

	// get the controller
	if content, err = account.New(account.Create); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// call the walkthrought to the setting method
	if status, r, err = h.setters(c, content); err != nil {
		return c.JSON(status, err.Error())
	}

	// send the response
	return c.JSONBlob(http.StatusOK, r)
}

// EditUser : _
func (h Handler) EditUser(c echo.Context) error {
	var content controller.Content
	var err error
	var status int
	var r []byte

	// get the controller
	if content, err = account.New(account.Update); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// call the walkthrought to the setting method
	if status, r, err = h.setters(c, content); err != nil {
		return c.JSON(status, err.Error())
	}

	// send the response
	return c.JSONBlob(http.StatusOK, r)
}

// DeleteUser : _
func (h Handler) DeleteUser(c echo.Context) error {
	var content controller.Content
	var err error
	var status int
	var r []byte

	// get the controller
	if content, err = account.New(account.Delete); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// call the walkthrought to the setting method
	if status, r, err = h.setters(c, content); err != nil {
		return c.JSON(status, err.Error())
	}

	// send the response
	return c.JSONBlob(http.StatusOK, r)
}
