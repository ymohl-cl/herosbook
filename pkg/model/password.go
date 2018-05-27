package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

// Passwords informations
// When one and two are equal, we can perform edit or create a new password to the current user
type Password struct {
	One string `json:"password_1" valid:"printableascii"`
	Two string `json:"password_2" valid:"printableascii"`
	Old string `json:"oldPassword" valid:"-"`
}

// Reset Passwords
func (p *Password) Reset() {
	p.One = ""
	p.Two = ""
	p.Old = ""
}

// Validate provide a great checker format without work with datas provided by the client
func (p Password) Validate() (err error) {
	if ok := govalidator.StringMatches(p.One, p.Two); !ok {
		return errors.New("passwords differs")
	}
	return nil
}
