package auth

import (
	"github.com/ymohl-cl/herosbook/pkg/app/jsonvalidator"
	"github.com/ymohl-cl/herosbook/pkg/cassandra"
	"github.com/ymohl-cl/herosbook/pkg/hasher"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
	"golang.org/x/xerrors"
)

// Auth management
type Auth struct {
	userDB        postgres.Client
	saltDB        cassandra.Client
	jsonValidator jsonvalidator.JSONValidator
	hasher        hasher.Hasher
}

// New authenticator service
func New(appName string) (Auth, error) {
	var a Auth
	var err error

	if a.userDB, err = postgres.New(appName + "_auth_postgres"); err != nil {
		return Auth{}, err
	}
	if a.saltDB, err = cassandra.New(appName + "_auth_cassandra"); err != nil {
		return Auth{}, err
	}
	a.jsonValidator = jsonvalidator.New()
	a.hasher = hasher.New()
	return a, nil
}

// Create account service
func (a Auth) Create(user User, password string) (User, error) {
	var err error
	var querySQL postgres.Query
	var queryCQL cassandra.Query
	var salt []byte
	var hashPass []byte
	var row postgres.ScanRow

	if hashPass, salt, err = a.hasher.Hash(password); err != nil {
		return User{}, err
	}

	// record user in database
	if querySQL, err = postgres.NewQuery(`INSERT INTO h_user(
		pseudo,
		last_name,
		first_name,
		hashpass,
		age,
		genre,
		email) VALUES($1, $2, $3, ($4::bytea), $5, $6, $7) RETURNING id`,
		user.Pseudo,
		user.LastName,
		user.Name,
		hashPass,
		user.Age,
		user.Genre,
		user.Email,
	); err != nil {
		return User{}, err
	}
	if row, err = a.userDB.WithRow(querySQL); err != nil {
		return User{}, err
	}
	if err = row.Scan(&user.Identifier); err != nil {
		return User{}, err
	}

	// save the user's salt
	if queryCQL, err = cassandra.NewQuery(`INSERT INTO heroesbook.salts(
		pseudo,
		salt) VALUES(?, ?)`,
		user.Pseudo,
		salt,
	); err != nil {
		return User{}, err
	}
	if err = a.saltDB.Create(queryCQL); err != nil {
		return User{}, err
	}
	return user, nil
}

// User getter with username and password parameters
func (a Auth) User(username, password string) (User, error) {
	var queryCQL cassandra.Query
	var scanCQL cassandra.Scanner
	var querySQL postgres.Query
	var scanSQL postgres.ScanRow
	var u User
	var salt []byte
	var err error
	var ok bool

	if queryCQL, err = cassandra.NewQuery(`SELECT salt FROM heroesbook.salts WHERE pseudo = ?`,
		username,
	); err != nil {
		return User{}, err
	}
	scanCQL = a.saltDB.Read(queryCQL)
	defer scanCQL.Close()
	if ok = scanCQL.Next(&salt); !ok {
		return User{}, xerrors.Errorf("get user's salt failed for the user %s, no result found !", username)
	}

	var pass []byte
	if pass, err = a.hasher.Decrypt(password, salt); err != nil {
		return User{}, err
	}

	if querySQL, err = postgres.NewQuery(`SELECT
		id,
		pseudo,
		last_name,
		first_name,
		age,
		genre,
		email FROM h_user WHERE pseudo = $1 AND hashpass = ($2::bytea)`,
		username,
		pass,
	); err != nil {
		return User{}, err
	}
	if scanSQL, err = a.userDB.WithRow(querySQL); err != nil {
		return User{}, err
	}
	if err = scanSQL.Scan(&u.Identifier, &u.Pseudo, &u.LastName, &u.Name, &u.Age, &u.Genre, &u.Email); err != nil {
		return User{}, err
	}
	return u, nil
}
