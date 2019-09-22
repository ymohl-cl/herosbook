package model

import "time"

// Book api model
type Book struct {
	Identifier   string    `json:"identifier"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Genre        string    `json:"genre"`
	Publish      bool      `json:"publish"`
	Authors      []string  `json:"authors"`
	Owner        string    `json:"owner"`
	NodeIDS      []string  `json:"nodeIds"`
	CreationDate time.Time `json:"creationDate"`
	Board        string    `json:"boardId"`
}

// IsEditable check if the book can be edit
// Check if the book is not published
// Check if the user is owner or author to update it
func (b Book) IsEditable(user string) bool {
	if b.Owner == user {
		return true
	}
	for _, v := range b.Authors {
		if v == user {
			return true
		}
	}
	return false
}

// SearchBook filter
// All criterias are cumulative
type SearchBook struct {
	Title  string    `json:"title"`
	Genre  string    `json:"genre"`
	Author string    `json:"authors"`
	From   time.Time `json:"fromDate"`
	To     time.Time `json:"toDate"`
}

// Node api model
type Node struct {
	Identifier        string   `json:"identifier"`
	Authors           []string `json:"authors"`
	Owner             string   `json:"owner"`
	ChildNodes        []string `json:"childNodes"`
	ConditionnalNodes []string `json:"conditionnalNodes"`
	Content           string   `json:"content"`
}
