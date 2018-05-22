package postgres

import (
	"database/sql"

	"github.com/ymohl-cl/herosbook/pkg/model"
)

// Postgres define the payload JSON matcher to instanciate a postgres service
type Postgres interface {
	Client() (d Driver, err error)
}

type json struct {
	DriverName string `json:"driver" valid:"required"`
	User       string `json:"user" valid:"required"`
	Password   string `json:"password" valid:"required"`
	DbName     string `json:"db" valid:"required"`
	Ssl        string `json:"sslmode" valid:"required"`
	HostName   string `json:"host" valid:"required"`
	Port       string `json:"port" valid:"required"`
}

// Driver implement interface to *sql.DB
type Driver interface {
	Close() error
	NewUser(user model.User, password []byte) (err error)
}

// New return a json struct to match the payload
func New() (p Postgres) {
	return &json{}
}

// Client instanciate a new postgres service
func (j json) Client() (d Driver, err error) {
	var connSTR string

	connSTR += "user=" + j.User + " "
	connSTR += "password=" + j.Password + " "
	connSTR += "dbname=" + j.DbName + " "
	connSTR += "sslmode=" + j.Ssl + " "
	connSTR += "host=" + j.HostName + " "
	connSTR += "port=" + j.Port
	if d, err = sql.Open(j.DriverName, connSTR); err != nil {
		return nil, err
	}
	return d, nil
}
