package handler

import (
	"github.com/labstack/echo"
)

// Authenticated check if user has a right to ressources access
func (h handle) Authenticated(s string, c echo.Context) (ok bool, err error) {
	var PublicID string

	if c.Path() == "/" || c.Path() == "/login" || c.Path() == "/users" {
		return true, err
	}
	if err = c.Bind(&PublicID); err != nil {
		return false, err
	}
	/*	if ok, err = controller.CheckToken(h.cassandra, ID.Value, []byte(s)); err != nil {
		return false, err
	}*/
	if !ok {
		return false, nil
	}
	return true, nil
}

// Login : walktrougth to manager
/*
func (a *Authenticator) Login(psql *sql.DB, cql *gocql.Session) (int, error) {
	var err error
	var ok bool

	// get public id
	a.Log.ID.Value, err = controller.GetPublicIDByName(psql, a.Log.Pseudo)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	// check the non empty public id
	if a.Log.ID.Value == "" {
		return http.StatusBadRequest, errors.New("Pseudo not found")
	}
	// verify the password
	if ok, err = controller.CheckPassword(psql, cql, a.Log.Password, a.Log.ID.Value); err != nil {
		return http.StatusInternalServerError, err
	}
	if !ok {
		return http.StatusBadRequest, errors.New("Bad password")
	}

	// create and save the Token
	if a.Log.Token, err = controller.CreateToken(cql, a.Log.ID.Value); err != nil {
		return http.StatusInternalServerError, err
	}
	return 0, err
}*/

/*
	token := make([]byte, 32)

	// generate token
	if _, err = io.ReadFull(rand.Reader, token); err != nil {
		return token, err
	}

	// save token on cassandra with a time life
	err = cql.Query(`INSERT INTO herosbook.tokens(
		id_psql,
		value) VALUES(?, ?) USING TTL ?`,
		publicID, token, lifeToken).Exec()
return token, err

*/
