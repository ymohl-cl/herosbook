package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// ConnectAccount provide a fresh token if authentication is available
func (h handle) ConnectAccount(c echo.Context) (err error) {
	var account model.Account
	var response []byte

	if err = h.jsonParser(c, &account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	if err = h.m.ConnectAccount(&account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	account.Passwords.Reset()
	if response, err = h.jsonBuilder(account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	return c.JSONBlob(http.StatusOK, response)
}

/*
// Disconnect : _
func (h handle) Disconnect(c echo.Context) error {
	return c.String(http.StatusOK, "Disconnect : _")
}
*/

// CreateAccount handler
func (h handle) CreateAccount(c echo.Context) (err error) {
	var account model.Account
	var response []byte

	if err = h.jsonParser(c, &account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	if err = h.m.CreateAccount(&account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	account.Passwords.Reset()
	if response, err = h.jsonBuilder(account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	return c.JSONBlob(http.StatusOK, response)
}

// UpdateUser handler
func (h handle) UpdateUser(c echo.Context) (err error) {
	var account model.Account
	var response []byte

	if err = h.jsonParser(c, &account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	if err = h.m.UpdateUser(&account.User); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	account.Passwords.Reset()
	if response, err = h.jsonBuilder(account.User); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	return c.JSONBlob(http.StatusOK, response)
}

// DeleteAccount handler
func (h handle) DeleteAccount(c echo.Context) (err error) {
	var account model.Account
	var response []byte

	if err = h.jsonParser(c, &account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	if err = h.m.DeleteAccount(&account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	if response, err = h.jsonBuilder(account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	return c.JSONBlob(http.StatusOK, response)
}

// UpdatePassword handler
func (h handle) UpdatePassword(c echo.Context) (err error) {
	var account model.Account
	var response []byte

	if err = h.jsonParser(c, &account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	if err = h.m.UpdatePassword(&account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	if response, err = h.jsonBuilder(account); err != nil {
		return c.JSON(h.getHTTPStatus(err), err.Error())
	}

	account.Passwords.Reset()
	return c.JSONBlob(http.StatusOK, response)
}
