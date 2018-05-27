package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

// Parameters users
const (
	PseudoSizeMin = 4
	PseudoSizeMax = 56
	AgeMin        = 10
	AgeMax        = 142
)

// User data
// Private fields are escaped by json "-"
type User struct {
	PublicID  string `json:"public_id" valid:"-"`
	PrivateID string `json:"-"`
	Pseudo    string `json:"pseudo" valid:"alphanum"`
	Age       uint8  `json:"age" valid:"required"`
	Sex       string `json:"sex" valid:"alpha, in(male|female)"`
	Email     string `json:"email" valid:"email"`
}

// Validate provide a great checker format without work with datas provided by the client
func (u User) Validate() (err error) {
	if ok, err := govalidator.ValidateStruct(u); !ok {
		return err
	}
	if ok := govalidator.IsByteLength(u.Pseudo, PseudoSizeMin, PseudoSizeMax); !ok {
		return errors.New("Inappropriate pseudo size")
	}
	if ok := govalidator.InRange(int(u.Age), int(AgeMin), int(AgeMax)); !ok {
		return errors.New("age field must be in a range of 10 to 142")
	}
	return nil
}
