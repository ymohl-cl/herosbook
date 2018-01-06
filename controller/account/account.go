package account

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gocql/gocql"
	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/models/users"
)

// liste of tag
const (
	Create = iota
	Update
	Delete
)

// Account manage create edit and delete user
type Account struct {
	user users.User
	Rec  func(*sql.DB, *gocql.Session) (int, error)
	Resp func() ([]byte, int, error)
}

// New : _
func New(tag int) (*Account, error) {
	a := Account{}

	switch tag {
	case Create:
		a.Rec = a.recordCreate
		a.Resp = a.returnUser
	case Update:
		a.Rec = a.recordUpdate
		a.Resp = a.returnUser
	case Delete:
		a.Rec = a.recordDelete
		a.Resp = a.returnID
	default:
		return nil, errors.New("tag used on account controller don't exist")
	}

	return &a, nil
}

// JSONParser : _
func (a *Account) JSONParser(c echo.Context) (int, error) {
	if err := c.Bind(&a.user); err != nil {
		return http.StatusPreconditionFailed, err
	}
	return 0, nil
}

// IsValid : _
func (a Account) IsValid(tag int) (int, error) {
	if ok, err := govalidator.ValidateStruct(a.user); !ok {
		return http.StatusBadRequest, err
	}
	if ok := govalidator.IsByteLength(a.user.Infos.Pseudo, users.PseudoSizeMin, users.PseudoSizeMax); !ok {
		return http.StatusBadRequest, errors.New("Inappropriate pseudo size")
	}

	if ok := govalidator.InRange(int(a.user.Infos.Age), int(users.AgeMin), int(users.AgeMax)); !ok {
		return http.StatusBadRequest, errors.New("age field must be in a range of 10 to 142")
	}
	if ok := govalidator.StringMatches(a.user.Pass.One, a.user.Pass.Two); !ok {
		return http.StatusBadRequest, errors.New("passwords differs")
	}
	return 0, nil
}

// Record is implementation interface controller
func (a *Account) Record(psql *sql.DB, cql *gocql.Session) (int, error) {
	return a.Rec(psql, cql)
}

// Response is implementation interface controller
func (a *Account) Response() ([]byte, int, error) {
	return a.Resp()
}
