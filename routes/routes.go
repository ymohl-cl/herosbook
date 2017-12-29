package routes

import (
	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/handlers"
)

// Descriptions of all routes
func Descriptions(e *echo.Echo, h *handlers.Handler) {
	// Home page
	e.GET("/", h.Home)

	// connect managment
	e.POST("/login", h.Connect)
	e.DELETE("/login", h.Disconnect)

	// users managment
	e.GET("/users", h.GetUsers)
	e.GET("/users/:id", h.GetUser)
	e.POST("/users", h.CreateUser)
	e.PUT("/users/:id", h.EditUser)
	e.DELETE("/users/:id", h.DeleteUser)

	// book managment
	e.GET("/books", h.GetBooks)
	e.GET("/books/:id", h.GetBook)
	e.POST("/books", h.CreateBook)
	e.PUT("/books/:id", h.EditBook)
	e.DELETE("/books/:id", h.DeleteBook)

	// content book managment
	e.GET("/books/:id/contents", h.GetContentsBook)
	e.GET("/books/:id/contents/:idc", h.GetContentBook)
	e.POST("/books/:id/contents", h.CreateContentBook)
	e.PUT("/books/:id/contents/:idc", h.EditContentBook)
	e.DELETE("/books/:id/contents/:idc", h.DeleteContentBook)
}
