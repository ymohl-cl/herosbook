package cassandra

import (
	"github.com/gocql/gocql"
)

// Scanner descibe an interface to read each response from postgres
type Scanner interface {
	Next(models ...interface{}) bool
	Close()
}

type scanner struct {
	rows *gocql.Iter
}

func newScanner(rows *gocql.Iter) Scanner {
	return &scanner{rows: rows}
}

// Next get the next result from the rows scanner
// If no result, the booleen will set to false, else to true
// You need call the close method before forget the scanner
func (s scanner) Next(models ...interface{}) bool {
	return s.rows.Scan(models...)
}

// Close sql rows
func (s scanner) Close() {
	s.rows.Close()
}
