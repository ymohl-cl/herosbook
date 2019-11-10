package controller

import (
	"fmt"

	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
	"github.com/ymohl-cl/herosbook/pkg/xerror"
	"golang.org/x/xerrors"
)

// RecordCategory to the book id
// check is use only to veriff than user is book's owner. It's not a case, Scan return an error and catch the not found error
func (c controller) RecordCategory(cat model.Category, userID, bookID string) (model.Category, xerror.Xerror) {
	var err error
	var querySQL postgres.Query
	var tx postgres.Transaction
	var row postgres.ScanRow
	var str string

	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return model.Category{}, newInternalErr(err)
	}
	defer tx.Rollback()

	// check if user is owner book
	str = `SELECT id FROM h_book WHERE id = $1 AND owner_id = $2`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return model.Category{}, newInternalErr(err)
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Category{}, newInternalErr(err)
	}
	check := ""
	if err = row.Scan(&check); err != nil {
		return model.Category{}, catchErr(err)
	}

	// insert category
	str = `INSERT INTO h_category(name_category, title, description, book_id) VALUES($1, $2, $3, $4) RETURNING id`
	if querySQL, err = postgres.NewQuery(str,
		cat.Type,
		cat.Title,
		cat.Description,
		bookID,
	); err != nil {
		return model.Category{}, newInternalErr(err)
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Category{}, newInternalErr(err)
	}
	if err = row.Scan(&cat.Identifier); err != nil {
		return model.Category{}, catchErr(err)
	}

	tx.Commit()
	return cat, nil
}

// UpdateCategory to the book id
// check if user is book's owner
// check if title category is not used to this book
func (c controller) UpdateCategory(cat model.Category, userID, bookID string) xerror.Xerror {
	var err error
	var querySQL postgres.Query
	var str string
	var nbAffectedRow int64

	str = `UPDATE h_category SET title = $1,
		name_category = $2,
		description = $3 WHERE
		book_id = $4 AND
		EXISTS(SELECT id FROM h_book WHERE id = $4 AND owner_id = $5) AND
		id = $6;`
	if querySQL, err = postgres.NewQuery(str,
		cat.Title,
		cat.Type,
		cat.Description,
		bookID,
		userID,
		cat.Identifier,
	); err != nil {
		return newInternalErr(err)
	}
	if nbAffectedRow, err = c.driverSQL.WithNoRow(querySQL); err != nil {
		return newInternalErr(err)
	}
	if nbAffectedRow != 1 {
		return newNoContentErr(xerrors.New(fmt.Sprintf("number rows affected by the update category request: %d", nbAffectedRow)))
	}
	return nil
}

// DeleteCategory from the database if:
// user_id is book's owner
// category exist to this book
func (c controller) DeleteCategory(categoryID, userID, bookID string) xerror.Xerror {
	var err error
	var querySQL postgres.Query
	var str string
	var nbAffectedRow int64

	str = `DELETE FROM h_category WHERE
		exists(SELECT id FROM h_book WHERE id = $1 AND owner_id = $2) AND
		book_id = $1 AND
		id = $3;`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
		categoryID,
	); err != nil {
		return newInternalErr(err)
	}
	if nbAffectedRow, err = c.driverSQL.WithNoRow(querySQL); err != nil {
		return newInternalErr(err)
	}
	if nbAffectedRow != 1 {
		return newNoContentErr(xerrors.New(fmt.Sprintf("number rows affected by the delete category request: %d", nbAffectedRow)))
	}
	return nil
}
