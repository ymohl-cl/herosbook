package cassandra

// SaveSalt record user's salt on the cassandra's bdd
func (c client) SaveSalt(publicID string, salt []byte) (err error) {
	if err := c.driver.Query(`INSERT INTO herosbook.users(
		id_psql,
		salt) VALUES(?, ?)`,
		publicID, salt).Exec(); err != nil {
		return err
	}
	return nil
}

// DeleteSalt remove user's salt on the cassandra's bdd
func (c client) DeleteSalt(publicID string) (err error) {
	if err := c.driver.Query(`DELETE FROM herosbook.users WHERE
		id_psql) VALUES(?)`,
		publicID).Exec(); err != nil {
		return err
	}
	return nil

}
