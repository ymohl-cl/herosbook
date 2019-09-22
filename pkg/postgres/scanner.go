package postgres

import "database/sql"

// ScanRows descibe an interface to read any rows from postgres
type ScanRows interface {
	Next(models ...interface{}) (bool, error)
	Close()
}

// ScanRow descibe an interface to read only on row from postgres
type ScanRow interface {
	Scan(models ...interface{}) error
}

type scanRows struct {
	rows *sql.Rows
}

type scanRow struct {
	row *sql.Row
}

func newScanRows(rows *sql.Rows) (ScanRows, error) {
	var err error
	var s scanRows

	if err = rows.Err(); err != nil {
		rows.Close()
		return nil, err
	}
	s.rows = rows
	return &s, nil
}

func newScanRow(row *sql.Row) ScanRow {
	var s scanRow

	s.row = row
	return &s
}

// Next get the next result from the rows scanner
// If error occurred, return it
// If no result, the booleen will set to false, else to true
// You need call the close method before forget the scanner
func (s scanRows) Next(models ...interface{}) (bool, error) {
	var err error

	if ok := s.rows.Next(); !ok {
		if err = s.rows.Err(); err != nil {
			return false, err
		}
		return false, nil
	}
	if err = s.rows.Scan(models...); err != nil {
		return false, err
	}
	return true, nil
}

// Close sql rows
func (s scanRows) Close() {
	s.rows.Close()
}

// Scan column in the returned row
func (s scanRow) Scan(models ...interface{}) error {
	if err := s.row.Scan(models...); err != nil {
		return err
	}
	return nil
}
