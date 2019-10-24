package postgres

import (
	"database/sql"
)

// Transaction is an implementation mockable to transaction postgres sql
// Transaction is instanciable by NewTransaction method from Client
// Use defer rollback to undo the change after error occured
// Don't forget to call the commit method at the end
type Transaction interface {
	WithRows(q Query) (ScanRows, error)
	WithRow(q Query) (ScanRow, error)
	WithNoRow(q Query) (int64, error)
	Commit()
	Rollback()
}

type transaction struct {
	driver *sql.Tx
}

// Run the request with only one row wanted
// Return a scanner interface to read the row returned
func (t transaction) WithRow(q Query) (ScanRow, error) {
	var row *sql.Row

	row = t.driver.QueryRow(q.Content(), q.ARGS()...)
	return newScanRow(row), nil
}

// Run the request with any rows wanted
// Return a scanner interface to read the rows returned
func (t transaction) WithRows(q Query) (ScanRows, error) {
	var rows *sql.Rows
	var err error

	if rows, err = t.driver.Query(q.Content(), q.ARGS()...); err != nil {
		return nil, err
	}
	return newScanRows(rows)
}

// Execute query without get the result
// Return the number rows affected or error if occured
func (t transaction) WithNoRow(q Query) (int64, error) {
	var err error
	var result sql.Result
	var nbAffectedRows int64

	if result, err = t.driver.Exec(q.Content(), q.ARGS()...); err != nil {
		return 0, err
	}
	if nbAffectedRows, err = result.RowsAffected(); err != nil {
		return 0, err
	}
	return nbAffectedRows, nil
}

// Commit the end transaction
func (t transaction) Commit() {
	t.driver.Commit()
}

// Call rollback actions if error occured.
// No effect if commit is already called
func (t transaction) Rollback() {
	t.driver.Rollback()
}
