package cassandra

import "github.com/gocql/gocql"

// Cassandra define the payload JSON matcher to instanciate a cassandra service
type Cassandra interface {
	Client() (d Driver, err error)
}

type json struct {
	Clusters []string `json:"clusters" valid:"required"`
	Keyspace string   `json:"keyspace" valid:"required"`
}

// Driver implement interface to *gocql.Session
type Driver interface {
	Close()
}

// New return a json struct to match the payload
func New() (c Cassandra) {
	return &json{}
}

// Client instanciate a new cassandra service
func (j json) Client() (d Driver, err error) {
	// create clusters
	cluster := gocql.NewCluster(j.Clusters...)
	cluster.Consistency = gocql.LocalOne
	cluster.ProtoVersion = 4
	cluster.Keyspace = j.Keyspace

	// get session
	if d, err = cluster.CreateSession(); err != nil {
		return nil, err
	}

	return d, nil
}
