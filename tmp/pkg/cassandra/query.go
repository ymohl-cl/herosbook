package cassandra

/*
func (a *Account) newUser(psql *sql.DB, password []byte) error {
	if err := psql.QueryRow(`INSERT INTO users(
		pseudo,
		password,
		age,
		sex,
		email) VALUES($1, $2, $3, $4, $5) RETURNING id_public`,
		a.user.Infos.Pseudo,
		password,
		a.user.Infos.Age,
		a.user.Infos.Sex,
		a.user.Infos.Email).Scan(&a.user.ID.Value); err != nil {
		return err
	}

	return nil
}

func (a *Account) newSalt(cql *gocql.Session, salt []byte) error {
	if err := cql.Query(`INSERT INTO herosbook.users(
		id_psql,
		salt) VALUES(?, ?)`,
		a.user.ID.Value, salt).Exec(); err != nil {
		return err
	}
	return nil
}


func (a Account) updateSalt(cql *gocql.Session, salt []byte) error {
	if err := cql.Query(`UPDATE herosbook.users SET
		salt = ? WHERE id_psql = ?`,
		salt, a.user.ID.Value).Exec(); err != nil {
		fmt.Println("SALT ERROR")
		return err
	}
	return nil
}

func (a Account) deleteUser(psql *sql.DB) error {
	if _, err := psql.Exec(`DELETE FROM users WHERE
		id_public = $1`,
		a.user.ID.Value); err != nil {
		return err
	}
	return nil
}

func (a Account) deleteSalt(cql *gocql.Session) error {
	if err := cql.Query(`DELETE FROM herosbook.users WHERE
		id_psql=?`,
		a.user.ID.Value).Exec(); err != nil {
		return err
	}
	return nil
}
*/
