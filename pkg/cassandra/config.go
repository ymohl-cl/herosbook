package cassandra

// Config datas to instanciate a cql client type Cassandra
// - Hosts accept the comma separator to any clusters
type Config struct {
	Hosts    string `required:"true"`
	Keyspace string `required:"true"`
	Port     int    `required:"true"`
}
