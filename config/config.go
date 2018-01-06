package config

import (
	"database/sql"
	"encoding/json"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/gocql/gocql"
)

// Config : content file config_example.json
type Config struct {
	API  Service   `json:"api" valid:"required"`
	Psql Postgres  `json:"psql" valid:"required"`
	Cass Cassandra `json:"cassandra" valid:"required"`
}

// Service : infos on the service
type Service struct {
	Cert   string `json:"certificat" valid:"required"`
	Key    string `json:"key" valid:"required"`
	Domain string `json:"domain" valid:"required"`
}

// Postgres : infos on driver postgres sql
type Postgres struct {
	DriverName string `json:"driver" valid:"required"`
	User       string `json:"user" valid:"required"`
	Password   string `json:"password" valid:"required"`
	DbName     string `json:"db" valid:"required"`
	Ssl        string `json:"sslmode" valid:"required"`
	HostName   string `json:"host" valid:"required"`
	Port       string `json:"port" valid:"required"`
}

// Cassandra : infos on driver cassandra cql
type Cassandra struct {
	Clusters []string `json:"clusters" valid:"required"`
	Keyspace string   `json:"keyspace" valid:"required"`
}

// New : instance configuration
func New() (*Config, error) {
	var c Config

	// open config file
	file, err := os.Open("/bin/config_example.json")
	if err != nil {
		return nil, err
	}

	// get infos file
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// create buffer file
	data := make([]byte, stat.Size())

	// read file
	if _, err := file.Read(data); err != nil {
		return nil, err
	}

	// unmarshall JSON
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	// check ValidateStruct
	if ok, err := govalidator.ValidateStruct(&c); !ok {
		return nil, err
	}

	return &c, nil
}

// Init Driver postgres sql
func (p Postgres) Init() (*sql.DB, error) {
	var psql *sql.DB
	var err error
	var connStr string
	esp := " "

	connStr += "user=" + p.User + esp
	connStr += "password=" + p.Password + esp
	connStr += "dbname=" + p.DbName + esp
	connStr += "sslmode=" + p.Ssl + esp
	connStr += "host=" + p.HostName + esp
	connStr += "port=" + p.Port

	if psql, err = sql.Open(p.DriverName, connStr); err != nil {
		return nil, err
	}
	return psql, nil
}

// Init Cassandra cql
func (c Cassandra) Init() (*gocql.Session, error) {
	var session *gocql.Session
	var err error

	// create clusters
	cluster := gocql.NewCluster(c.Clusters...)
	cluster.Consistency = gocql.LocalOne
	cluster.ProtoVersion = 4
	cluster.Keyspace = c.Keyspace

	// get session
	if session, err = cluster.CreateSession(); err != nil {
		return nil, err
	}

	return session, nil
}