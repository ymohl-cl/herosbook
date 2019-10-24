package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/pkg/app/auth"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// CreateCategory to the book
func (h Handler) CreateCategory(c echo.Context) error {
	var err error
	var user auth.User
	var cat model.Category

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	if err = h.jsonValidator.Bind(c.Request(), &cat); err != nil {
		fmt.Printf("CreateCategory - jsonvalidator.Bind - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	if cat, err = h.controller.RecordCategory(cat, user.Identifier, bookID); err != nil {
		fmt.Printf("CreateCategory - h.controller.RecordCategory - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	return c.JSON(http.StatusCreated, &cat)
}

// UpdateCategory update the content category
func (h Handler) UpdateCategory(c echo.Context) error {
	var err error
	var user auth.User
	var cat model.Category

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	if err = h.jsonValidator.Bind(c.Request(), &cat); err != nil {
		fmt.Printf("UpdateCategory - jsonvalidator.Bind - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	if err = h.controller.UpdateCategory(cat, user.Identifier, bookID); err != nil {
		fmt.Printf("UpdateCategory - h.controller.UpdateCategory - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	return c.NoContent(http.StatusOK)
}

// RemoveCategory the category
func (h Handler) RemoveCategory(c echo.Context) error {
	var err error
	var user auth.User

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	categoryID := c.Param("id_category")
	if err = h.controller.DeleteCategory(categoryID, user.Identifier, bookID); err != nil {
		fmt.Printf("RemoveCategory - h.controller.DeleteCategory - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	return c.NoContent(http.StatusOK)
}
