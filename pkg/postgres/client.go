package postgres

import (
	"database/sql"

	"github.com/ymohl-cl/herosbook/pkg/model"
)

// ClientI perform requests to the application
// Has a driver sql on his instance
type ClientI interface {
	Close() (err error)
	CreateAccount(a *model.Account, password []byte) (err error)
	DeleteAccount(a *model.Account) (err error)
	UpdateUser(user *model.User) (err error)
	UpdatePassword(user model.User, password []byte) (err error)
}

type client struct {
	driver *sql.DB
}

// Close fd used by sql.DB driver
func (c client) Close() (err error) {
	return c.driver.Close()
}
