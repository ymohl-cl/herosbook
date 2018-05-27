package manager

import (
	"github.com/ymohl-cl/herosbook/pkg/cassandra"
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
)

// Manager implement type to manage all queries from client
type Manager interface {
	ConnectAccount(a *model.Account) (err error)
	//	Disconnect() (err error)
	CreateAccount(a *model.Account) (err error)
	UpdateUser(user *model.User) (err error)
	DeleteAccount(a *model.Account) (err error)
	UpdatePassword(a *model.Account) (err error)
}

type manage struct {
	clientSQL postgres.ClientI
	clientCQL cassandra.ClientI
}

// New provide the Manager interface to execute the action relative at client's query
func New(clientSQL postgres.ClientI, clientCQL cassandra.ClientI) (m Manager) {
	m = &manage{
		clientCQL: clientCQL,
		clientSQL: clientSQL,
	}
	return m
}
