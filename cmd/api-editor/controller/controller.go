package controller

import (
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
	"github.com/ymohl-cl/herosbook/pkg/xerror"
)

// Controller interface to implement private manager resources
type Controller interface {
	// book controller
	RecordBook(b *model.Book) xerror.Xerror
	ReadBook(bookID string, userID string) (model.Book, xerror.Xerror)
	ReadBooks(filters model.SearchBook, userID string) ([]model.Book, xerror.Xerror)
	UpdateBook(b model.Book, userID string) (model.Book, xerror.Xerror)
	DeleteBook(bookID string, userID string) xerror.Xerror
	// node controller
	RecordNode(n model.Node, userID, bookID string) (model.Node, xerror.Xerror)
	UpdateNode(n model.Node, userID, bookID string) xerror.Xerror
	ReadNode(nodeID, bookID, userID string) (model.Node, xerror.Xerror)
	ReadNodes(bookID, userID string) ([]string, xerror.Xerror)
	DeleteNode(nodeID, userID, bookID string) xerror.Xerror
	// category controller
	RecordCategory(c model.Category, userID string, bookID string) (model.Category, xerror.Xerror)
	UpdateCategory(c model.Category, userID string, bookID string) xerror.Xerror
	DeleteCategory(categoryID, userID, bookID string) xerror.Xerror
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
