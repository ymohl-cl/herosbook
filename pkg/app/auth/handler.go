package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// Login parameters (header control)
const (
	usernameKey = "username"
	passwordKey = "password"
)

// Login handler
func (a Auth) Login(c echo.Context) error {
	var err error
	var u User

	login := c.Request().Header.Get(usernameKey)
	pass := c.Request().Header.Get(passwordKey)
	decoded := []byte{}
	if decoded, err = base64.StdEncoding.DecodeString(pass); err != nil {
		fmt.Printf("Login - base64.StdEncoding.DecodeString - %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal error",
		})
	}
	pass = string(decoded)
	if !strings.HasPrefix(pass, login+":") {
		fmt.Printf("Login - strings.HasPrefix - search: %s | got: %s\n", login+":", pass)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	pass = strings.TrimPrefix(pass, login+":")

	if u, err = a.User(login, pass); err != nil {
		fmt.Printf("Login - a.User - error: %s \n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	var token string
	if token, err = NewToken(u); err != nil {
		fmt.Printf("Login - NewToken - error: %s \n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  u,
		"token": token,
	})
}

// Register endpoint api
// - decode the password from password header (encodedBase64 with pseudo:password)
// - parse the user body
// - trime the decoded password to remove the pseudo: pattern
// - record the user
func (a Auth) Register(c echo.Context) error {
	var user User
	var err error

	pass := c.Request().Header.Get(passwordKey)
	decoded := []byte{}
	if decoded, err = base64.StdEncoding.DecodeString(pass); err != nil {
		fmt.Printf("Register - base64.StdEncoding.DecodeString - %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal error",
		})
	}
	pass = string(decoded)
	if err = a.jsonValidator.Bind(c.Request(), &user); err != nil {
		fmt.Printf("Register - jsonvalidator.Bind - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	if !strings.HasPrefix(pass, user.Pseudo+":") {
		fmt.Printf("Register - strings.HasPrefix - search: %s | got: %s\n", user.Pseudo+":", pass)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	pass = strings.TrimPrefix(pass, user.Pseudo+":")
	if user, err = a.Create(user, pass); err != nil {
		fmt.Printf("Register - auth.Create - %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal error",
		})
	}
	return c.JSON(http.StatusOK, &user)
}
