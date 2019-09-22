package postgres

import (
	"database/sql"

	// need to use postges sql with database sql
	_ "github.com/lib/pq"
	"github.com/ymohl-cl/herosbook/pkg/config"
)

// Client is an implementation mockable to database/sql server
// Client is instanciable by New method calling
// Don't forget to call the close method at the end
type Client interface {
	WithRows(q Query) (ScanRows, error)
	WithRow(q Query) (ScanRow, error)
	WithNoRow(q Query) error
	NewTransaction() (Transaction, error)
	Close() error
}

type client struct {
	driver *sql.DB
}

// New provide a postgres driver and load the configuration from vars environments
// The searches vars are the next:
// User
// Password
// DBName
// SSL
// Host
// Port
func New(appName string) (Client, error) {
	var conf Config
	var err error

	if err = config.ParseEnv(appName, &conf); err != nil {
		return nil, err
	}
	return NewWithConfig(conf)
}

// NewWithConfig provide a postgres driver without load the configuration
// We assum that the config is correctly setup
func NewWithConfig(conf Config) (Client, error) {
	var err error
	var connSTR string
	var c client

	connSTR = "user=" + conf.User + " "
	connSTR += "password=" + conf.Password + " "
	connSTR += "dbname=" + conf.DbName + " "
	connSTR += "sslmode=" + conf.Ssl + " "
	connSTR += "host=" + conf.Host + " "
	connSTR += "port=" + conf.Port
	if c.driver, err = sql.Open("postgres", connSTR); err != nil {
		return nil, err
	}
	if err = c.driver.Ping(); err != nil {
		return nil, err
	}
	return &c, nil
}

// Close the fd used by sql.DB driver
func (c client) Close() error {
	return c.driver.Close()
}

// Run the request with only one row wanted
// Return a scanner interface to read the row returned
func (c client) WithRow(q Query) (ScanRow, error) {
	var row *sql.Row

	row = c.driver.QueryRow(q.Content(), q.ARGS()...)
	return newScanRow(row), nil
}

// Run the request with any rows wanted
// Return a scanner interface to read the rows returned
func (c client) WithRows(q Query) (ScanRows, error) {
	var rows *sql.Rows
	var err error

	if rows, err = c.driver.Query(q.Content(), q.ARGS()...); err != nil {
		return nil, err
	}
	return newScanRows(rows)
}

// Execute query without get the result
func (c client) WithNoRow(q Query) error {
	if _, err := c.driver.Exec(q.Content(), q.ARGS()...); err != nil {
		return err
	}
	return nil
}

// NewTransaction return a transaction driver to exec transac sql requests
func (c client) NewTransaction() (Transaction, error) {
	var err error
	var t transaction

	if t.driver, err = c.driver.Begin(); err != nil {
		return nil, err
	}
	return &t, nil
}
