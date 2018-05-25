package postgres

import "database/sql"

// Conf postgres provide a set parameters to instanciate a new client
// Conf define the tad to your configuration file
type Conf struct {
	DriverName string `json:"driver" valid:"required"`
	User       string `json:"user" valid:"required"`
	Password   string `json:"password" valid:"required"`
	DBName     string `json:"db" valid:"required"`
	Ssl        string `json:"sslmode" valid:"required"`
	HostName   string `json:"host" valid:"required"`
	Port       string `json:"port" valid:"required"`
}

// New provide a postgres client
func (c Conf) New() (driver ClientI, err error) {
	var connSTR string
	var cl client

	connSTR += "user=" + c.User + " "
	connSTR += "password=" + c.Password + " "
	connSTR += "dbname=" + c.DBName + " "
	connSTR += "sslmode=" + c.Ssl + " "
	connSTR += "host=" + c.HostName + " "
	connSTR += "port=" + c.Port
	if cl.driver, err = sql.Open(c.DriverName, connSTR); err != nil {
		return nil, err
	}
	return &cl, nil
}
