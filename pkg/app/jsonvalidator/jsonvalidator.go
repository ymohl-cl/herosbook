package jsonvalidator

import (
	"encoding/json"
	"net/http"

	"golang.org/x/xerrors"
	"gopkg.in/go-playground/validator.v9"
)

// JSONValidator is an implementation of validation data stucture
// will be attached on http driver
type JSONValidator interface {
	Bind(r *http.Request, input interface{}) error
}

type jsonValidator struct {
	driver *validator.Validate
}

// New return an implement of JSONValidator
func New() JSONValidator {
	return &jsonValidator{
		driver: validator.New(),
	}
}

// Bind parse the input json parameter et validate it
func (j jsonValidator) Bind(r *http.Request, input interface{}) error {
	var err error

	if r.Body == nil {
		return xerrors.New("Can't bind json data with empty body")
	}
	if err = json.NewDecoder(r.Body).Decode(input); err != nil {
		return err
	}
	if err = j.driver.Struct(input); err != nil {
		return err
	}
	return nil
}
