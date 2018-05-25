package postgres

import "github.com/ymohl-cl/herosbook/pkg/model"

// CreateUser on the sql DB
func (c client) CreateUser(user model.User, password []byte) (err error) {
	if err = c.driver.QueryRow(`INSERT INTO users(
		pseudo,
		password,
		age,
		sex,
		email) VALUES($1, $2, $3, $4, $5) RETURNING id_public`,
		user.Infos.Pseudo,
		password,
		user.Infos.Age,
		user.Infos.Sex,
		user.Infos.Email).Scan(&user.ID.Value); err != nil {
		return err
	}

	return nil
}

// UpdareUser on the sql DB
func (c client) UpdateUser(user model.User) (err error) {
	if _, err = c.driver.Exec(`UPDATE users SET
		pseudo = $1,
		age = $2,
		sex = $3,
		email = $4 WHERE id_public = $5`,
		user.Infos.Pseudo,
		user.Infos.Age,
		user.Infos.Sex,
		user.Infos.Email,
		user.ID.Value); err != nil {
		return err
	}

	return nil
}

// UpdatePassword on the sql DB
func (c client) UpdatePassword(user model.User, password []byte) (err error) {
	if _, err = c.driver.Exec(`UPDATE users SET
 		password = $1 WHERE id_public = $2`,
		password,
		user.ID.Value); err != nil {
		return err
	}
	return nil
}
