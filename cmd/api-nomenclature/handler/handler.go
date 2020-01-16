package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/ymohl-cl/herosbook/cmd/api-nomenclature/controller"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// Handler data with drivers need to make the job
type Handler struct {
	controller controller.Controller
}

// New handler object
// Describe all routes need to the app
func New(appName string, server *echo.Group) error {
	var h Handler

	h.controller = controller.New(appName)
	server.GET("/", h.Nomenclature)
	server.GET("/user", h.GenreUser)
	server.GET("/langage", h.Langage)
	server.GET("/book", h.Book)
	server.GET("/book/category", h.BookCategory)
	server.GET("/book/theme", h.BookTheme)

	return nil
}

// Nomenclature return all data api with the adaptative langage
func (h Handler) Nomenclature(c echo.Context) error {
	var langage controller.Langage
	var paramLang string
	var output model.NomenclatureOutput

	paramLang = c.Param("langage")
	langage = h.controller.NewLangage(paramLang)
	output.User.List = h.controller.UserGenre(langage)
	output.User.Number = len(output.User.List)
	output.BookCategory.List = h.controller.BookCategory(langage)
	output.BookCategory.Number = len(output.BookCategory.List)
	output.BookTheme.List = h.controller.BookTheme(langage)
	output.BookTheme.Number = len(output.BookTheme.List)
	output.Langage.List = h.controller.Langage(langage)
	output.Langage.Number = len(output.Langage.List)

	return c.JSON(http.StatusOK, &output)
}
