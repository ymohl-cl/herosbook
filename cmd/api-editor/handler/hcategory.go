package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/pkg/app/auth"
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/xerror"
)

// CreateCategory to the book
func (h Handler) CreateCategory(c echo.Context) error {
	var err error
	var xerr xerror.Xerror
	var user auth.User
	var cat model.Category

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	if err = h.jsonValidator.Bind(c.Request(), &cat); err != nil {
		return c.JSON(h.httpResponse(xerror.New(-1, err.Error())))
	}
	if cat, xerr = h.controller.RecordCategory(cat, user.Identifier, bookID); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.JSON(http.StatusCreated, &cat)
}

// UpdateCategory update the content category
func (h Handler) UpdateCategory(c echo.Context) error {
	var err error
	var xerr xerror.Xerror
	var user auth.User
	var cat model.Category

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	if err = h.jsonValidator.Bind(c.Request(), &cat); err != nil {
		return c.JSON(h.httpResponse(xerror.New(-1, err.Error())))
	}
	if xerr = h.controller.UpdateCategory(cat, user.Identifier, bookID); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.NoContent(http.StatusOK)
}

// RemoveCategory the category
func (h Handler) RemoveCategory(c echo.Context) error {
	var xerr xerror.Xerror
	var user auth.User

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	categoryID := c.Param("id_category")
	if xerr = h.controller.DeleteCategory(categoryID, user.Identifier, bookID); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.NoContent(http.StatusOK)
}
