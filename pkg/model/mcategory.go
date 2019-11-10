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
	Identifier  string `json:"identifier" validate:"omitempty,uuid4"`
	Type        string `json:"type" validate:"required"`
	Title       string `json:"title" validate:"max=255,required"`
	Description string `json:"description"`
	BookID      string `json:"bookId" validate:"omitempty,uuid4"`
}

// Validate implentation of jsonvalidator.Model
func (c Category) Validate() error {
	switch c.Type {
	case tPerson, tCustom, tPlace:
		break
	default:
		return xerrors.New("unknow type category: " + c.Type)
	}
	return nil
}
