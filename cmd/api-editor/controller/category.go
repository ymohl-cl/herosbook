package controller

import (
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
	"golang.org/x/xerrors"
)

// RecordCategory to the book id
func (c controller) RecordCategory(cat model.Category, userID, bookID string) (model.Category, error) {
	var err error
	var querySQL postgres.Query
	var tx postgres.Transaction
	var row postgres.ScanRow
	var str string

	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return model.Category{}, err
	}
	defer tx.Rollback()

	// check if user is owner book
	str = `SELECT id FROM h_book WHERE id = $1 AND owner_id = $2`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return model.Category{}, err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Category{}, err
	}
	check := ""
	if err = row.Scan(&check); err != nil {
		return model.Category{}, err
	}
	if check != bookID {
		return model.Category{}, xerrors.New("user can't add a category")
	}

	// insert category
	str = `INSERT INTO h_category(name_category, title, description, book_id) VALUES($1, $2, $3, $4) RETURNING id`
	if querySQL, err = postgres.NewQuery(str,
		cat.Type,
		cat.Title,
		cat.Description,
		bookID,
	); err != nil {
		return model.Category{}, err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Category{}, err
	}
	if err = row.Scan(&cat.Identifier); err != nil {
		return model.Category{}, err
	}

	tx.Commit()
	return cat, nil
}

// UpdateCategory to the book id
// check if user is book's owner
// check if title category is not used to this book
func (c controller) UpdateCategory(cat model.Category, userID, bookID string) error {
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
		return err
	}
	if nbAffectedRow, err = c.driverSQL.WithNoRow(querySQL); err != nil {
		return err
	}
	if nbAffectedRow != 1 {
		return xerrors.New("error to update the category")
	}
	return nil
}

// DeleteCategory from the database if:
// user_id is book's owner
// category exist to this book
func (c controller) DeleteCategory(categoryID, userID, bookID string) error {
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
		return err
	}
	if nbAffectedRow, err = c.driverSQL.WithNoRow(querySQL); err != nil {
		return err
	}
	if nbAffectedRow != 1 {
		return xerrors.New("error to update the category")
	}
	return nil
}
