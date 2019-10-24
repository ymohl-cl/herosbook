package controller

import (
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
)

// Controller interface to implement private manager resources
type Controller interface {
	RecordBook(b *model.Book) error
	ReadBook(bookID string, userID string) (model.Book, error)
	ReadBooks(filters model.SearchBook, userID string) ([]model.Book, error)
	UpdateBook(b model.Book, userID string) (model.Book, error)
	//	RecordNode(n model.Node, userID string) error
	//	ReadNode(nodeID string, userID string) (model.Node, error)
	//	ReadNodes(filters model.SearchNode, userID string) ([]model.Node, error)
	//	Update(n model.Node, userID string) (model.Node, error)
	RecordCategory(c model.Category, userID string, bookID string) (model.Category, error)
	UpdateCategory(c model.Category, userID string, bookID string) error
	DeleteCategory(categoryID, userID, bookID string) error
}

type controller struct {
	driverSQL postgres.Client
}

// New controller to crud resources with databases
func New(appName string) (Controller, error) {
	var c controller
	var err error

	if c.driverSQL, err = postgres.New(appName + "_app_postgres"); err != nil {
		return nil, err
	}
	return &c, nil
}
