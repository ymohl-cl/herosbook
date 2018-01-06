package account

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gocql/gocql"
	"github.com/ymohl-cl/gen-pwd/generator"
	"github.com/ymohl-cl/herosbook/controller"
)

func (a *Account) recordCreate(psql *sql.DB, cql *gocql.Session) (int, error) {
	var err error
	var password, salt []byte

	// get password encrypted
	genPWD := generator.NewByDefault()
	password, salt, err = genPWD.CreateNewPassword(a.user.Pass.One)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// save new user on Psql
	if err = a.newUser(psql, password); err != nil {
		return http.StatusInternalServerError, err
	}

	// save salt on cassandra
	if err = a.newSalt(cql, salt); err != nil {
		return http.StatusInternalServerError, err
	}
	return 0, nil
}

func (a *Account) recordUpdate(psql *sql.DB, cql *gocql.Session) (int, error) {
	var err error
	var ok bool
	var password, salt []byte

	if ok, err = controller.CheckPassword(psql, cql, a.user.Pass.Old, a.user.ID.Value); err != nil {
		return http.StatusInternalServerError, err
	}
	// check the good password to update user
	if !ok {
		return http.StatusBadRequest, errors.New("bad password")
	}

	if err = a.updateUser(psql); err != nil {
		return http.StatusInternalServerError, err
	}

	// check if new password
	if strings.Compare(a.user.Pass.Old, a.user.Pass.One) != 0 {
		// get new encrypted password
		genPWD := generator.NewByDefault()
		password, salt, err = genPWD.CreateNewPassword(a.user.Pass.One)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		if err = a.updatePassword(psql, password); err != nil {
			return http.StatusInternalServerError, err
		}
		if err = a.updateSalt(cql, salt); err != nil {
			return http.StatusInternalServerError, err
		}
	}
	return 0, nil
}

func (a *Account) recordDelete(psql *sql.DB, cql *gocql.Session) (int, error) {
	var err error

	if err = a.deleteUser(psql); err != nil {
		return http.StatusInternalServerError, err
	}

	if err = a.deleteSalt(cql); err != nil {
		return http.StatusInternalServerError, err
	}
	return 0, nil
}

// returnUser delete Pass data and return struct user
func (a *Account) returnUser() ([]byte, int, error) {
	a.user.ClearPasswords()

	// encode struct
	m, err := json.Marshal(a.user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// get a map from the interface
	var tmp interface{}
	json.Unmarshal(m, &tmp)
	b := tmp.(map[string]interface{})

	// delete key Passwords
	delete(b, "passwords")

	ret, err := json.Marshal(b)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return ret, 0, nil
}

func (a *Account) returnID() ([]byte, int, error) {
	// encode struct
	m, err := json.Marshal(a.user.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// get a map from the interface
	var tmp interface{}
	json.Unmarshal(m, &tmp)
	b := tmp.(map[string]interface{})

	ret, err := json.Marshal(b)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return ret, 0, nil
}
