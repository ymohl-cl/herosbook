package cassandra

import "github.com/gocql/gocql"

// New provide a basic client to perform request with cassanda driver
func New(clusters []string, keyspace string) (driver *gocql.Session, err error) {
	// create clusters
	cluster := gocql.NewCluster(clusters...)
	cluster.Consistency = gocql.LocalOne
	cluster.ProtoVersion = 4
	cluster.Keyspace = keyspace

	// get session
	if driver, err = cluster.CreateSession(); err != nil {
		return nil, err
	}

	return driver, nil
}
