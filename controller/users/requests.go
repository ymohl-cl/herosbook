package users

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
)

const (
	pseudoSizeMin = 4
	pseudoSizeMax = 20
	passSize      = 60
)

// User : _
type User struct {
	Pseudo string `json:"pseudo"`
	Pass   string `json:"password"`
	Age    uint8  `json:"age"`
	Sex    bool   `json:"sex"`
	Email  string `json:"email" valid:"email"`
}

// IsValid : _
func (u User) IsValid() (int, error) {
	if ok, err := govalidator.ValidateStruct(u); !ok {
		return http.StatusBadRequest, err
	}
	if ok := govalidator.IsByteLength(u.Pseudo, pseudoSizeMin, pseudoSizeMax); !ok {
		return http.StatusBadRequest, errors.New("Inappropriate pseudo size")
	}

	return 0, nil
}

func (u User) Record(psql *sql.DB) (int, error) {
	return 0, nil
}

func (u User) Response() (interface{}, int, error) {
	return u, 0, nil
}
