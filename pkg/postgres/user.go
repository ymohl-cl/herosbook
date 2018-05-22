package postgres

import (
	"database/sql"

	"github.com/ymohl-cl/herosbook/pkg/model"
)

func (d Driver) CreateUser(user model.User, password []byte) (err error) {
	var psql *sql.DB

	psql = d.(*sql.DB)
	if err = psql.QueryRow(`INSERT INTO users(
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

func (d Driver) UpdateUser(user model.User) (err error) {
	var psql *sql.DB

	psql = d.(*sql.DB)
	if _, err = psql.Exec(`UPDATE users SET
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

func (d Driver) UpdatePassword(user model.User, password []byte) (err error) {
	if _, err = psql.Exec(`UPDATE users SET
 		password = $1 WHERE id_public = $2`,
		password,
		user.ID.Value); err != nil {
		return err
	}
	return nil
}
