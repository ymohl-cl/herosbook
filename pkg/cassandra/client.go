package cassandra

import "github.com/gocql/gocql"

// ClientI perform requests to the application
// Has a driver cql on his instance
type ClientI interface {
	Close() (err error)
	Salt(publicID string) (salt []byte, err error)
	SaveSalt(publicID string, salt []byte) (err error)
	DeleteSalt(publicID string) (err error)
	SaveToken(publicID string, token []byte, lifeToken int64) (err error)
	Token(publicID string) (myToken string, err error)
}

type client struct {
	driver *gocql.Session
}

// Close fd used by sql.DB driver
func (c client) Close() (err error) {
	c.driver.Close()
	return nil
}
