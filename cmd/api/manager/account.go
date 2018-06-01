package manager

import (
	"bytes"
	"crypto/rand"
	"io"

	"github.com/ymohl-cl/gen-pwd/generator"
	"github.com/ymohl-cl/herosbook/cmd/api/constant"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// ConnectAccount get the current account and extract password and salt user.
// If password matches, create a new token and return it
func (m manage) ConnectAccount(a *model.Account) (err error) {
	var cmp *model.Account
	var salt []byte

	if cmp, err = m.clientSQL.Account(a.User.Pseudo); err != nil {
		return err
	}
	if salt, err = m.clientCQL.Salt(cmp.User.PublicID); err != nil {
		return err
	}
	genPWD := generator.NewByDefault()
	var pass []byte
	if pass, err = genPWD.GetEncryptedPassword(string(a.Passwords.One), salt); err != nil {
		return err
	}

	if bytes.Compare(pass, []byte(cmp.Passwords.One)) != 0 {
		return err
	}
	token := make([]byte, 32)
	if _, err = io.ReadFull(rand.Reader, token); err != nil {
		return err
	}
	if err = m.clientCQL.SaveToken(cmp.User.PublicID, token, constant.Public.LifeAPIToken); err != nil {
		return err
	}

	// attach on the pointer a the account getted on cmp
	cmp.Token = string(token)
	a = cmp
	return nil
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
