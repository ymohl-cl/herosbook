package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/ymohl-cl/herosbook/cmd/api-nomenclature/controller"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// BookCategory return the nomenclature category to books
func (h Handler) BookCategory(c echo.Context) error {
	var langage controller.Langage
	var paramLang string
	var output model.Nomenclature

	paramLang = c.Param("langage")
	langage = h.controller.NewLangage(paramLang)
	output.List = h.controller.BookCategory(langage)
	output.Number = len(output.List)
	return c.JSON(http.StatusOK, &output)
}

// BookTheme return the nomenclature theme to books
func (h Handler) BookTheme(c echo.Context) error {
	var langage controller.Langage
	var paramLang string
	var output model.Nomenclature

	paramLang = c.Param("langage")
	langage = h.controller.NewLangage(paramLang)
	output.List = h.controller.BookTheme(langage)
	output.Number = len(output.List)
	return c.JSON(http.StatusOK, &output)
}

// Book return the nomenclature to the books
func (h Handler) Book(c echo.Context) error {
	var langage controller.Langage
	var paramLang string
	var output model.NomenclatureBookOutput

	paramLang = c.Param("langage")
	langage = h.controller.NewLangage(paramLang)
	output.BookCategory.List = h.controller.BookCategory(langage)
	output.BookCategory.Number = len(output.BookCategory.List)
	output.BookTheme.List = h.controller.BookTheme(langage)
	output.BookTheme.Number = len(output.BookTheme.List)

	return c.JSON(http.StatusOK, &output)
}
