package auth

// User api
type User struct {
	Identifier string `json:"identifier"`
	Pseudo     string `json:"pseudo" validate:"required"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Age        int16  `json:"age"`
	Genre      string `json:"genre"`
	Email      string `json:"email" validate:"required,email"`
}

// Validate implementation to use in jsonvalidator service
func (u User) Validate() error {
	return nil
}
