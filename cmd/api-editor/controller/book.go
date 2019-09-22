package controller

import (
	"github.com/lib/pq"
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
)

// RecordBook in sql database
func (c controller) RecordBook(b *model.Book) error {
	var err error
	var querySQL postgres.Query
	var tx postgres.Transaction
	var row postgres.ScanRow
	var boardID string

	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return nil
	}
	defer tx.Rollback()

	// create and get board id to the futur new book
	if querySQL, err = postgres.NewQuery(`INSERT INTO h_board DEFAULT VALUES RETURNING id;`); err != nil {
		return err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return err
	}
	if err = row.Scan(&boardID); err != nil {
		return err
	}
	// create book
	str := `INSERT INTO h_book(title, description, genre, publish, owner_id, creation_date, board_id) VALUES($1, $2, $3, $4, $5, 'now', $6) RETURNING id, creation_date`
	if querySQL, err = postgres.NewQuery(str,
		b.Title,
		b.Description,
		b.Genre,
		b.Publish,
		b.Owner,
		boardID,
	); err != nil {
		return err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return err
	}
	if err = row.Scan(&b.Identifier, &b.CreationDate); err != nil {
		return err
	}
	b.Board = boardID
	tx.Commit()
	return nil
}

// ReadBook return one book by id if user can acces it
func (c controller) ReadBook(bookID string, userID string) (model.Book, error) {
	var b model.Book
	/*	var err error
		var querySQL postgres.Query
		var scanner postgres.ScanRow

		if querySQL, err = postgres.NewQuery(`SELECT * FROM books WHERE id = $1 AND ( owner_id = $2 OR publish = $3 OR $2 = ANY (authors) )`,
			bookID,
			userID,
			true,
		); err != nil {
			return model.Book{}, err
		}
		if scanner, err = c.driverSQL.WithRow(querySQL); err != nil {
			return model.Book{}, err
		}
		defer scanner.Close()
		if _, err = scanner.Next(&b.Identifier, &b.Title, &b.Description, &b.Genre, &b.Publish, pq.Array(&b.Authors), &b.Owner, pq.Array(&b.NodeIDS), &b.CreationDate); err != nil {
			return model.Book{}, err
		}*/
	return b, nil
}

// ReadBooks return all book describe by the filter model and return it
func (c controller) ReadBooks(filters model.SearchBook, userID string) ([]model.Book, error) {
	var err error
	var strs []string
	var books []model.Book
	var querySQL postgres.Query
	var scanner postgres.ScanRows

	if filters.Title != "" {
		str := `title = ` + filters.Title
		strs = append(strs, str)
	}
	if filters.Genre != "" {
		str := `genre = ` + filters.Genre
		strs = append(strs, str)
	}
	if filters.Author != "" {
		str := filters.Author + ` = ANY (authors)`
		strs = append(strs, str)
	}
	if !filters.From.IsZero() {
		str := `creation date >= ` + filters.From.String()
		strs = append(strs, str)
	}
	if !filters.To.IsZero() {
		str := `creation date <= ` + filters.To.String()
		strs = append(strs, str)
	}
	querySTR := `SELECT * FROM h_book WHERE `
	if len(strs) == 0 {
		querySTR += `owner = ` + userID
	} else {
		for i, v := range strs {
			if i > 0 {
				querySTR += ` AND `
			}
			querySTR += v
		}
	}

	if querySQL, err = postgres.NewQuery(querySTR); err != nil {
		return nil, err
	}
	if scanner, err = c.driverSQL.WithRows(querySQL); err != nil {
		return nil, err
	}
	defer scanner.Close()
	ok := true
	for ok {
		b := model.Book{}
		if ok, err = scanner.Next(&b.Identifier, &b.Title, &b.Description, &b.Genre, &b.Publish, pq.Array(&b.Authors), &b.Owner, pq.Array(&b.NodeIDS), &b.CreationDate); err != nil {
			return nil, err
		}
		if ok {
			books = append(books, b)
		}
	}
	return books, nil
}

func (c controller) UpdateBook(b model.Book, userID string) (model.Book, error) {
	//	var err error
	//	var querySQL postgres.Query
	/*
		// record user in database
		if querySQL, err = postgres.NewQuery(`UPDATE books
			SET title = $1,
			describe = $2,
			genre = $3,
			publish = $4,
			authors = $5
			WHERE id = $6 AND publish = 'false' AND ( owner_id = $7 OR $7 = ANY (authors) )
			RETURNING owner_id, nodes, creation_date`,
			b.Title,
			b.Description,
			b.Genre,
			b.Publish,
			pq.Array(b.Authors),
			b.Identifier,
			userID,
		); err != nil {
			return model.Book{}, err
		}
		if err = c.driverSQL.Create(querySQL, &b.Owner, pq.Array(&b.NodeIDS), &b.CreationDate); err != nil {
			return model.Book{}, err
		}
	*/
	return b, nil
}
