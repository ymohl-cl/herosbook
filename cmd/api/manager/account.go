package manager

import (
	"bytes"
	"crypto/rand"
	"io"

	"github.com/ymohl-cl/gen-pwd/generator"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// TODO : TERMINATED CONNECTACCOUNT (see comment on middleware)
func (m manage) ConnectAccount(a *model.Account) (token string, err error) {
	var requester *model.Account
	var salt []byte

	// get user's information
	if requester, err = m.clientSQL.Account(a.User.Pseudo); err != nil {
		return token, err
	}

	// get salt password
	if salt, err = m.clientCQL.Salt(requester.User.PublicID); err != nil {
		return token, err
	}

	// get  encrypted password
	genPWD := generator.NewByDefault()
	var pass []byte
	if pass, err = genPWD.GetEncryptedPassword(a.Passwords.One, salt); err != nil {
		return token, err
	}

	if bytes.Compare(pass, requester.Passwords.One) != 0 {
		return token, err
	}

	// generate token
	t := make([]byte, 32)
	if _, err = io.ReadFull(rand.Reader, t); err != nil {
		return token, err
	}

	// save it on cql cassandra bdd
	if err = m.clientCQL.SaveToken(requester.User.PublicID, t); err != nil {
		return token, err
	}

	token = string.(t)
	return token, nil
}

func (m manage) Disconnect() (err error) {
	return nil
}

// CreateAccount provide a new account on sql db and cql db to the salt saving
func (m manage) CreateAccount(a *model.Account) (err error) {
	var password, salt []byte

	// get password encrypted
	genPWD := generator.NewByDefault()
	password, salt, err = genPWD.CreateNewPassword(a.Passwords.One)
	if err != nil {
		return err
	}

	// save new account on Psql
	if err = m.clientSQL.CreateAccount(a, password); err != nil {
		return err
	}

	// save salt on cassandra
	if err = m.clientCQL.SaveSalt(a.User.PublicID, salt); err != nil {
		return err
	}
	return nil
}

// UpdateUser on sql db
// Note which UpdateUser not perform the password update (see updatePassword request)
func (m manage) UpdateUser(u *model.User) (err error) {
	if err = m.clientSQL.UpdateUser(u); err != nil {
		return err
	}

	return nil
}

// DeleteAccount remove the current user from sql db
// Too remove the salt saved on cql server
func (m manage) DeleteAccount(a *model.Account) (err error) {
	if err = m.clientSQL.DeleteAccount(a); err != nil {
		return err
	}

	if err = m.clientCQL.DeleteSalt(a.User.PublicID); err != nil {
		return err
	}

	// remove the current instance to provide a fresh account
	a = new(model.Account)
	return nil
}

// UpdatePassword udpate the old password on from sql bdd
// Update too a new salt on cql bdd
func (m manage) UpdatePassword(a *model.Account) (err error) {
	var password, salt []byte

	// get password encrypted
	genPWD := generator.NewByDefault()
	password, salt, err = genPWD.CreateNewPassword(a.Passwords.One)
	if err != nil {
		return err
	}

	// save new account on Psql
	if err = m.clientSQL.UpdatePassword(a.User, password); err != nil {
		return err
	}

	// save salt on cassandra
	if err = m.clientCQL.SaveSalt(a.User.PublicID, salt); err != nil {
		return err
	}
	return nil
}
