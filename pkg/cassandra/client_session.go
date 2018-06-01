package cassandra

// Salt return salt from pubicID
func (c client) Salt(publicID string) (salt []byte, err error) {
	if err = c.driver.Query(`SELECT FROM herosbook.users(
		id_psql=$1)`, publicID).Scan(&salt); err != nil {
		return nil, err
	}
	return
}

// SaveSalt record user's salt on the cassandra's bdd
func (c client) SaveSalt(publicID string, salt []byte) (err error) {
	if err = c.driver.Query(`INSERT INTO herosbook.users(
		id_psql,
		salt) VALUES(?, ?)`,
		publicID, salt).Exec(); err != nil {
		return err
	}
	return nil
}

// DeleteSalt remove user's salt on the cassandra's bdd
func (c client) DeleteSalt(publicID string) (err error) {
	if err = c.driver.Query(`DELETE FROM herosbook.users WHERE
		id_psql = $1)`,
		publicID).Exec(); err != nil {
		return err
	}
	return nil

}

// SaveToken record user's token on the cassandra's bdd
func (c client) SaveToken(publicID string, token []byte, lifeToken int64) (err error) {
	if err = c.driver.Query(`INSERT INTO herosbook.tokens(
		id_psql,
		value) VALUES(?, ?) USING TTL ?`,
		publicID, token, lifeToken).Exec(); err != nil {
		return err
	}
	return nil
}

// Token return user's token which attached on the publicID parameter
func (c client) Token(publicID string) (myToken string, err error) {
	if err = c.driver.Query(`SELECT value FROM herosbook.tokens WHERE id_psql=?`,
		publicID).Scan(&myToken); err != nil {
		return "", err
	}
	return myToken, nil
}
