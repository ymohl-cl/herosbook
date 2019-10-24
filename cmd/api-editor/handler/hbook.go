package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/pkg/app/auth"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// CreateBook to the user given
func (h Handler) CreateBook(c echo.Context) error {
	var err error
	var user auth.User
	var b model.Book

	user = auth.ParseToken(c)
	if err = h.jsonValidator.Bind(c.Request(), &b); err != nil {
		fmt.Printf("CreateBook - jsonvalidator.Bind - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	b.Owner = user.Identifier
	b.Publish = false
	if err = h.controller.RecordBook(&b); err != nil {
		fmt.Printf("CreateBook - h.controller.RecordBook - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}

	return c.JSON(http.StatusOK, &b)
}

// GetBook search the book by id
// Need be book's owner or book's author to get the book or
// The book must be published
func (h Handler) GetBook(c echo.Context) error {
	var err error
	var user auth.User
	var b model.Book

	bookID := c.Param("id")
	user = auth.ParseToken(c)
	if b, err = h.controller.ReadBook(bookID, user.Identifier); err != nil {
		fmt.Printf("GetBook - h.controller.ReadBook - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	if b.Identifier == "" {
		return c.JSON(http.StatusNoContent, map[string]string{
			"message": "book not found",
		})
	}
	return c.JSON(http.StatusOK, &b)
}

// SearchBooks analyze the search filter object to create a sql request and provide all books associated
func (h Handler) SearchBooks(c echo.Context) error {
	var err error
	var user auth.User
	var filter model.SearchBook
	var books []model.Book

	user = auth.ParseToken(c)
	if err = h.jsonValidator.Bind(c.Request(), &filter); err != nil {
		fmt.Printf("SearchBooks - jsonvalidator.Bind - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	if books, err = h.controller.ReadBooks(filter, user.Identifier); err != nil {
		fmt.Printf("SearchBooks - h.controller.ReadBooks - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	if len(books) == 0 {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, &books)
}

// UpdateBook targeted by the bookID
func (h Handler) UpdateBook(c echo.Context) error {
	var err error
	var user auth.User
	var b model.Book

	user = auth.ParseToken(c)
	if err = h.jsonValidator.Bind(c.Request(), &b); err != nil {
		fmt.Printf("UpdateBook - jsonvalidator.Bind - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	if b, err = h.controller.UpdateBook(b, user.Identifier); err != nil {
		fmt.Printf("UpdateBook - h.controller.UpdateBook - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal-error",
		})
	}

	return c.JSON(http.StatusOK, &b)
}

// RemoveBook by id
func (h Handler) RemoveBook(c echo.Context) error {
	var err error
	var user auth.User

	bookID := c.Param("id")
	user = auth.ParseToken(c)
	if err = h.controller.DeleteBook(bookID, user.Identifier); err != nil {
		fmt.Printf("RemoveBook - h.controller.DeleteBook - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"identifier": bookID,
	})
}
