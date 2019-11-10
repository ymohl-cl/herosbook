package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/pkg/app/auth"
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/xerror"
)

// CreateBook to the user given
func (h Handler) CreateBook(c echo.Context) error {
	var err error
	var xerr xerror.Xerror
	var user auth.User
	var b model.Book

	user = auth.ParseToken(c)
	if err = h.jsonValidator.Bind(c.Request(), &b); err != nil {
		return c.JSON(h.httpResponse(xerror.New(-1, err.Error())))
	}
	b.Owner = user.Identifier
	b.Publish = false
	if xerr = h.controller.RecordBook(&b); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}

	return c.JSON(http.StatusOK, &b)
}

// GetBook search the book by id
// Need be book's owner or book's author to get the book or
// The book must be published
func (h Handler) GetBook(c echo.Context) error {
	var xerr xerror.Xerror
	var user auth.User
	var b model.Book

	bookID := c.Param("id")
	user = auth.ParseToken(c)
	if b, xerr = h.controller.ReadBook(bookID, user.Identifier); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.JSON(http.StatusOK, &b)
}

// SearchBooks analyze the search filter object to create a sql request and provide all books associated
func (h Handler) SearchBooks(c echo.Context) error {
	var err error
	var xerr xerror.Xerror
	var user auth.User
	var filter model.SearchBook
	var books []model.Book

	user = auth.ParseToken(c)
	if err = h.jsonValidator.Bind(c.Request(), &filter); err != nil {
		return c.JSON(h.httpResponse(xerror.New(-1, err.Error())))
	}
	if books, xerr = h.controller.ReadBooks(filter, user.Identifier); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.JSON(http.StatusOK, &books)
}

// UpdateBook targeted by the bookID
func (h Handler) UpdateBook(c echo.Context) error {
	var err error
	var xerr xerror.Xerror
	var user auth.User
	var b model.Book

	user = auth.ParseToken(c)
	if err = h.jsonValidator.Bind(c.Request(), &b); err != nil {
		return c.JSON(h.httpResponse(xerror.New(-1, err.Error())))
	}
	if b, xerr = h.controller.UpdateBook(b, user.Identifier); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.JSON(http.StatusOK, &b)
}

// RemoveBook by id
func (h Handler) RemoveBook(c echo.Context) error {
	var xerr xerror.Xerror
	var user auth.User

	bookID := c.Param("id")
	user = auth.ParseToken(c)
	if xerr = h.controller.DeleteBook(bookID, user.Identifier); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.NoContent(http.StatusOK)
}
