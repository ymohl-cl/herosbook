package auth

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtCustomClaims struct {
	jwt.StandardClaims
	User User `json:"user"`
}

// NewToken create a token with embeded user
func NewToken(user User) (string, error) {
	var t string
	var err error

	claims := &jwtCustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	if t, err = token.SignedString([]byte("--key-jwt--")); err != nil {
		return "", err
	}
	return t, nil
}

// ParseToken get the user configured inside
func ParseToken(c echo.Context) User {
	netUser := c.Get("user").(*jwt.Token)
	claims := netUser.Claims.(jwt.MapClaims)
	userBuff := claims["user"]

	// convert map to json
	jsonString, _ := json.Marshal(userBuff)

	var user User
	json.Unmarshal(jsonString, &user)
	return user
}
