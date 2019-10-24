package model

import (
	"golang.org/x/xerrors"
)

// type list
const (
	tPerson = "person"
	tPlace  = "place"
	tCustom = "custom"
)

// Category describe the section to attach node and help to the organization
type Category struct {
	Identifier  string `json:"id" validate:"omitempty,uuid4"`
	Type        string `json:"type" validate:"required"`
	Title       string `json:"title" validate:"max=255,required"`
	Description string `json:"description"`
}

// Validate implentation of jsonvalidator.Model
func (c Category) Validate() error {
	switch c.Type {
	case tPerson, tCustom, tPlace:
		break
	default:
		return xerrors.New("unknown type provided")
	}
	return nil
}
