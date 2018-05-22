package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// Connect : _
func (h handle) Connect(c echo.Context) error {
	c.Set("un", nil)
	return c.String(http.StatusOK, "Connect : _")
}

// Disconnect : _
func (h handle) Disconnect(c echo.Context) error {
	return c.String(http.StatusOK, "Disconnect : _")
}

// CreateUser : _
func (h handle) CreateUser(c echo.Context) error {
	var err error
	var user model.User
	var object interface{}
	var response []byte

	if user, err = h.jsonParser(c, user); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	if object, err = h.m.CreateUser(user); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	if response, err = h.jsonBuilder(object); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	return c.JSONBlob(http.StatusOK, response)
	// get the controller
	//	if content, err = account.New(account.Create); err != nil {
	//return c.JSON(http.StatusInternalServerError, err.Error())
	//}

	// send the response
}

// EditUser : _
func (h handle) EditUser(c echo.Context) error {
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
func (h handle) DeleteUser(c echo.Context) error {
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
