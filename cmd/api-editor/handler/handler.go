package handler

import (
	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/cmd/api-editor/controller"
	"github.com/ymohl-cl/herosbook/pkg/app/jsonvalidator"
)

// Handler data with drivers need to make the job
type Handler struct {
	controller    controller.Controller
	jsonValidator jsonvalidator.JSONValidator
}

// New handler object
// Describe all routes need to the app
func New(appName string, server *echo.Group) error {
	var h Handler
	var err error

	if h.controller, err = controller.New(appName); err != nil {
		return err
	}
	h.jsonValidator = jsonvalidator.New()
	// book
	server.POST("/books", h.CreateBook)
	server.GET("/books/:id", h.GetBook)
	server.POST("/books/_searches", h.SearchBooks)
	server.DELETE("/books/:id", h.RemoveBook)
	server.PUT("/books", h.UpdateBook)
	// category
	server.POST("/books/:id/category", h.CreateCategory)
	server.PUT("/books/:id/category", h.UpdateCategory)
	server.DELETE("/books/:id/category/:id_category", h.RemoveCategory)

	return nil
}
