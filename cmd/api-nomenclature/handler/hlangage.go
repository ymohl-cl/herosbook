package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/ymohl-cl/herosbook/cmd/api-nomenclature/controller"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// Langage return the nomenclature category to books
func (h Handler) Langage(c echo.Context) error {
	var langage controller.Langage
	var paramLang string
	var output model.Nomenclature

	paramLang = c.Param("langage")
	langage = h.controller.NewLangage(paramLang)
	output.List = h.controller.Langage(langage)
	output.Number = len(output.List)
	return c.JSON(http.StatusOK, &output)
}
