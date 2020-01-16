package model

import (
	"golang.org/x/xerrors"
)

// type list
const (
	tPerson   = "person"
	tLocation = "location"
	tCustom   = "custom"
)

// Categories model to easy group categories by type
type Categories struct {
	Persons   []Category `json:"persons"`
	Locations []Category `json:"locations"`
	Customs   []Category `json:"customs"`
}

// Add category on the adaptive list
func (c *Categories) Add(cat Category) {
	switch cat.Type {
	case tPerson:
		c.Persons = append(c.Persons, cat)
	case tLocation:
		c.Locations = append(c.Locations, cat)
	case tCustom:
		c.Customs = append(c.Customs, cat)
	}
}

// Category describe the section to attach node and help to the organization
type Category struct {
	Identifier  string `json:"identifier" validate:"omitempty,uuid4"`
	BookID      string `json:"bookId" validate:"omitempty,uuid4"`
	Type        string `json:"type" validate:"required"`
	Img         string `json:"img" validate:"omitempty"`
	Title       string `json:"title" validate:"max=255,required"`
	Description string `json:"description"`
}

// Validate implentation of jsonvalidator.Model
func (c Category) Validate() error {
	switch c.Type {
	case tPerson, tCustom, tLocation:
		break
	default:
		return xerrors.New("unknow type category: " + c.Type)
	}
	return nil
}
