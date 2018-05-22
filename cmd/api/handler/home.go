package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
** Home page
 */

// Home page, provide the books being read and books by themes and new books
func (h handle) Home(c echo.Context) (err error) {

	// Temporary to provide connection token
	/*	publicID := "90b69dee-b881-4da9-9356-7917c95c250b"
		token := make([]byte, 32)

		// generate token
		if _, err = io.ReadFull(rand.Reader, token); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// save token on cassandra to 30 seconds
		err = h.cassandra.Query(`INSERT INTO herosbook.tokens(id_psql, value) VALUES(?, ?) USING TTL 30`,
			publicID, token).Exec()
		if err != nil {
			fmt.Println("Here :()")
			return c.JSON(http.StatusInternalServerError, err.Error())
		}*/
	return c.String(http.StatusOK, "Hello, World! \n It's a home page")
}
