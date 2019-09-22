package postgres

import "errors"

// Query describe the postgres query
type Query interface {
	Content() string
	ARGS() []interface{}
}

type query struct {
	content string
	args    []interface{}
}

// NewQuery provide an helper to construct a postgres sql query
func NewQuery(content string, args ...interface{}) (Query, error) {
	var q query

	if content == "" {
		return nil, errors.New("the content query can't be empty")
	}
	q.content = content
	q.args = append(q.args, args...)
	return &q, nil
}

// Content getter to the query
func (q *query) Content() string {
	return q.content
}

// ARGS getter to the query
func (q *query) ARGS() []interface{} {
	return q.args
}
