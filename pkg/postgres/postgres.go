package postgres

import (
	"database/sql"
)

// New Provide a simple new sql client
func New(host, port, dbName, password, user string) (driver *sql.DB, err error) {
	var connSTR string

	connSTR += "user=" + user + " "
	connSTR += "password=" + password + " "
	connSTR += "dbname=" + dbName + " "
	connSTR += "host=" + host + " "
	connSTR += "port=" + port
	if driver, err = sql.Open("postgres", connSTR); err != nil {
		return nil, err
	}

	return driver, nil
}
