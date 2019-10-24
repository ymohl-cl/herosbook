package model

// Node api model
type Node struct {
	Identifier    string   `json:"identifier" validate:"omitempty,uuid4"`
	Title         string   `json:"title" validates:"max=255"`
	Description   string   `json:"description"`
	Content       string   `json:"content"`
	ChildIDS      []string `json:"childs"`
	ContionnalIDS []string `json:"conditionnals"`
	CategoryIDS   []string `json:"categories"`
}

// Validate implementation of jsonvalidator.Model
func (n Node) Validate() error {
	return nil
}
