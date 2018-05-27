package model

// Account manage create edit and delete user
type Account struct {
	User      User
	Passwords Password
	Token     string
}

// Validate
func (a Account) Validate() (err error) {
	if err = a.User.Validate(); err != nil {
		return err
	}
	if err = a.Passwords.Validate(); err != nil {
		return err
	}
	// TODO: Valide token (len and format) with govalidator
	return nil
}
