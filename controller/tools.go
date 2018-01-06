package controller

import (
	"bytes"
	"database/sql"
	"errors"

	"github.com/gocql/gocql"
	"github.com/ymohl-cl/gen-pwd/generator"
)

// CheckPassword : _
func CheckPassword(psql *sql.DB, cql *gocql.Session, pass, ID string) (bool, error) {
	var err error
	var password []byte
	var passTest []byte
	salt := make([]byte, 32)

	// get salt from cassandra
	err = cql.Query(`SELECT salt FROM herosbook.users WHERE id_psql=?`,
		ID).Scan(&salt)
	if err != nil {
		return false, err
	}

	// get  encrypted password
	genPWD := generator.NewByDefault()
	if passTest, err = genPWD.GetEncryptedPassword(pass, salt); err != nil {
		return false, err
	}

	// get the current password from cassandra
	err = psql.QueryRow(`SELECT password FROM users WHERE id_public=$1`, ID).Scan(&password)
	if err != nil {
		return false, err
	}

	// test if passwords are equals
	if bytes.Compare(passTest, password) != 0 {
		return false, errors.New("bad password")
	}
	return true, nil
}
