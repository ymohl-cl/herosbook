package model

// Node api model
type Node struct {
	Identifier        string   `json:"identifier"`
	ChildNodes        []string `json:"childNodes"`
	ConditionnalNodes []string `json:"conditionnalNodes"`
	Content           string   `json:"content"`
}
