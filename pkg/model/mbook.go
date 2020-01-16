package model

import "time"

// Book api model
type Book struct {
	Identifier   string     `json:"identifier"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Genre        string     `json:"genre"`
	Publish      bool       `json:"publish"`
	Owner        string     `json:"owner"`
	NodeIDS      []string   `json:"nodeIds"`
	CreationDate time.Time  `json:"creationDate"`
	Categories   Categories `json:"categories"`
}

// Validate book json model
func (b Book) Validate() error {
	return nil
}

// SearchBook filter
// All criterias are cumulative
type SearchBook struct {
	Title string `json:"title"`
	Genre string `json:"genre"`
}

// Validate implentation of jsonvalidator.Model
func (sb SearchBook) Validate() error {
	return nil
}
