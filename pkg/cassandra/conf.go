package cassandra

import "github.com/gocql/gocql"

// Conf cassandra provide a set parameters to instanciate a new client
// Conf define the tags to your configuration file
type Conf struct {
	Clusters []string `json:"clusters" valid:"required"`
	Keyspace string   `json:"keyspace" valid:"required"`
}

// New provide a cassandra client
func (c Conf) New() (driver ClientI, err error) {
	var cl client

	// create clusters
	cluster := gocql.NewCluster(c.Clusters...)
	cluster.Consistency = gocql.LocalOne
	cluster.ProtoVersion = 4
	cluster.Keyspace = c.Keyspace

	// get session
	if cl.driver, err = cluster.CreateSession(); err != nil {
		return nil, err
	}

	return &cl, nil
}
