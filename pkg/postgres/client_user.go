package postgres

import "github.com/ymohl-cl/herosbook/pkg/model"

// CreateAccount on the sql DB
func (c client) CreateAccount(a *model.Account, password []byte) (err error) {
	if err = c.driver.QueryRow(`INSERT INTO users(
		pseudo,
		password,
		age,
		sex,
		email) VALUES($1, $2, $3, $4, $5) RETURNING id_public`,
		a.User.Pseudo,
		password,
		a.User.Age,
		a.User.Sex,
		a.User.Email).Scan(&a.User.PublicID); err != nil {
		return err
	}

	return nil
}

// DeleteAccount from sql DB
func (c client) DeleteAccount(a *model.Account) (err error) {
	if _, err = c.driver.Exec(`DELETE FROM users WHERE 
		public_id = $1`,
		a.User.PublicID); err != nil {
		return err
	}

	return nil
}

// UpdateUser on the sql DB
func (c client) UpdateUser(user *model.User) (err error) {
	if _, err = c.driver.Exec(`UPDATE users SET
		pseudo = $1,
		age = $2,
		sex = $3,
		email = $4 WHERE id_public = $5`,
		user.Pseudo,
		user.Age,
		user.Sex,
		user.Email,
		user.PublicID); err != nil {
		return err
	}

	return nil
}

// UpdatePassword on the sql DB
func (c client) UpdatePassword(user model.User, password []byte) (err error) {
	if _, err = c.driver.Exec(`UPDATE users SET
 		password = $1 WHERE id_public = $2`,
		password,
		user.PublicID); err != nil {
		return err
	}
	return nil
}
