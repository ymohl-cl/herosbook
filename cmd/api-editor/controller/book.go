package controller

import (
	"fmt"

	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
	"golang.org/x/xerrors"
)

// RecordBook in sql database
func (c controller) RecordBook(b *model.Book) error {
	var err error
	var querySQL postgres.Query
	var tx postgres.Transaction
	var row postgres.ScanRow

	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return nil
	}
	defer tx.Rollback()

	// create book
	str := `INSERT INTO h_book(title, description, genre, publish, owner_id, creation_date) VALUES($1, $2, $3, $4, $5, 'now') RETURNING id, creation_date`
	if querySQL, err = postgres.NewQuery(str,
		b.Title,
		b.Description,
		b.Genre,
		b.Publish,
		b.Owner,
	); err != nil {
		return err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return err
	}
	if err = row.Scan(&b.Identifier, &b.CreationDate); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// ReadBook return one book by id if user can acces it
func (c controller) ReadBook(bookID string, userID string) (model.Book, error) {
	var b model.Book
	var err error
	var querySQL postgres.Query
	var row postgres.ScanRow
	var rows postgres.ScanRows
	var tx postgres.Transaction
	var str string

	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return model.Book{}, err
	}
	// get books
	str = `SELECT title, description, genre, publish, creation_date FROM h_book WHERE id = $1 AND owner_id = $2`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return model.Book{}, err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Book{}, err
	}
	if err = row.Scan(&b.Title, &b.Description, &b.Genre, &b.Publish, &b.CreationDate); err != nil {
		return model.Book{}, err
	}
	b.Identifier = bookID
	b.Owner = userID

	// get nodes
	str = `SELECT id FROM h_node WHERE book_id = $1`
	if querySQL, err = postgres.NewQuery(str, bookID); err != nil {
		return model.Book{}, err
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return model.Book{}, err
	}
	defer rows.Close()
	ok := true
	for ok {
		id := ""
		if ok, err = rows.Next(&id); err != nil {
			return model.Book{}, err
		}
		if ok {
			b.NodeIDS = append(b.NodeIDS, id)
		}
	}

	// get categories
	str = `SELECT id, name_category, title, description FROM h_category WHERE book_id = $1`
	if querySQL, err = postgres.NewQuery(str, bookID); err != nil {
		return model.Book{}, err
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return model.Book{}, err
	}
	defer rows.Close()
	ok = true
	for ok {
		cat := model.Category{}
		if ok, err = rows.Next(&cat.Identifier, &cat.Type, &cat.Title, &cat.Description); err != nil {
			return model.Book{}, err
		}
		if ok {
			cat.BookID = bookID
			b.Categories = append(b.Categories, cat)
		}
	}

	tx.Commit()
	return b, nil
}

// ReadBooks return all book describe by the filter model and return it
func (c controller) ReadBooks(filters model.SearchBook, userID string) ([]model.Book, error) {
	var err error
	var books []model.Book
	var querySQL postgres.Query
	var strs []string
	var rows postgres.ScanRows

	if filters.Title != "" {
		str := `title = '` + filters.Title + `'`
		strs = append(strs, str)
	}
	if filters.Genre != "" {
		str := `genre = '` + filters.Genre + `'`
		strs = append(strs, str)
	}

	str := `SELECT id, title, description, genre, publish, creation_date FROM h_book WHERE owner_id = $1`
	for _, v := range strs {
		str += ` AND ` + v
	}
	fmt.Println("query: ", str)
	if querySQL, err = postgres.NewQuery(str, userID); err != nil {
		return nil, err
	}
	if rows, err = c.driverSQL.WithRows(querySQL); err != nil {
		return nil, err
	}
	defer rows.Close()
	for ok := true; ok; {
		b := model.Book{}
		if ok, err = rows.Next(&b.Identifier, &b.Title, &b.Description, &b.Genre, &b.Publish, &b.CreationDate); err != nil {
			return nil, err
		}
		if ok {
			b.Owner = userID
			books = append(books, b)
		}
	}
	return books, nil
}

func (c controller) UpdateBook(b model.Book, userID string) (model.Book, error) {
	var err error
	var querySQL postgres.Query
	var row postgres.ScanRow

	str := `UPDATE h_book
			SET title = $1,
			description = $2,
			genre = $3,
			publish = $4 WHERE
			id = $5 AND
			owner_id = $6 RETURNING
			creation_date`
	if querySQL, err = postgres.NewQuery(str,
		b.Title,
		b.Description,
		b.Genre,
		b.Publish,
		b.Identifier,
		userID,
	); err != nil {
		return model.Book{}, err
	}
	if row, err = c.driverSQL.WithRow(querySQL); err != nil {
		return model.Book{}, err
	}
	if err = row.Scan(&b.CreationDate); err != nil {
		return model.Book{}, err
	}
	b.Owner = userID

	return b, nil
}

func (c controller) DeleteBook(bookID, userID string) error {
	var err error
	var querySQL postgres.Query
	var str string
	var nbAffectedRow int64

	str = `DELETE FROM h_book WHERE
		id = $1 AND
		owner_id = $2;`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return err
	}
	if nbAffectedRow, err = c.driverSQL.WithNoRow(querySQL); err != nil {
		return err
	}
	if nbAffectedRow != 1 {
		return xerrors.New("error to remove the book")
	}
	return nil
}
