package cassandra

import (
	"strings"

	"github.com/gocql/gocql"
	"github.com/ymohl-cl/herosbook/pkg/config"
)

// Client is an implementation mockable to gocql server
// Provide an easy crud methods
// Client is instanciable by New method calling
// Don't forget to call the close method at the end
type Client interface {
	Read(q Query) Scanner
	Create(q Query) error
	Update(q Query) error
	Delete(q Query) error
	Close()

	exec(q Query) error
}

type client struct {
	driver *gocql.Session
}

// New provide a cassandra driver and load the configuration from vars environments
// The searches vars are the next:
// Clusters
// Keyspace
func New(appName string) (Client, error) {
	var c Config
	var err error

	if err = config.ParseEnv(appName, &c); err != nil {
		return nil, err
	}
	return NewWithConfig(c)
}

// NewWithConfig provide a cassandra driver without load the configuration
// We assum that the config is correctly setup
func NewWithConfig(conf Config) (Client, error) {
	var err error
	var c client

	clusters := strings.Split(conf.Hosts, ",")
	clusterConfig := gocql.NewCluster(clusters...)
	clusterConfig.Consistency = gocql.LocalOne
	clusterConfig.ProtoVersion = 3
	clusterConfig.Keyspace = conf.Keyspace
	clusterConfig.Port = conf.Port

	if c.driver, err = clusterConfig.CreateSession(); err != nil {
		return nil, err
	}
	return &c, nil
}

// Close the fd used by cql driver
func (c client) Close() {
	c.driver.Close()
}

// Create data describe in the query
func (c client) Create(q Query) error {
	return c.exec(q)
}

// Read data describe by query from the postgres database
// Return a scanner interface to get each item returned
func (c client) Read(q Query) Scanner {
	var rows *gocql.Iter
	var scan Scanner

	rows = c.driver.Query(q.Content(), q.ARGS()...).Iter()
	scan = newScanner(rows)
	return scan
}

// Update data describe by the query parameter
func (c client) Update(q Query) error {
	return c.exec(q)
}

// Delete data describe by the qury parameter
func (c client) Delete(q Query) error {
	return c.exec(q)
}

func (c client) exec(q Query) error {
	var err error

	if err = c.driver.Query(q.Content(), q.ARGS()...).Exec(); err != nil {
		return err
	}
	return nil
}
