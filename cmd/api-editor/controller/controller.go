package controller

import (
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
)

// Controller interface to implement private manager resources
type Controller interface {
	// book controller
	RecordBook(b *model.Book) error
	ReadBook(bookID string, userID string) (model.Book, error)
	ReadBooks(filters model.SearchBook, userID string) ([]model.Book, error)
	UpdateBook(b model.Book, userID string) (model.Book, error)
	DeleteBook(bookID string, userID string) error
	// node controller
	RecordNode(n model.Node, userID, bookID string) (model.Node, error)
	UpdateNode(n model.Node, userID, bookID string) error
	ReadNode(nodeID, bookID, userID string) (model.Node, error)
	ReadNodes(bookID, userID string) ([]string, error)
	DeleteNode(nodeID, userID, bookID string) error
	// category controller
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
