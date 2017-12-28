package routes

import (
	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/handlers"
)

// Descriptions of all routes
func Descriptions(e *echo.Echo) {
	// Home page
	e.GET("/", handlers.Home)

	// connect managment
	e.POST("/login", handlers.Connect)
	e.DELETE("/login", handlers.Disconnect)

	// users managment
	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/:id", handlers.EditUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

	// book managment
	e.GET("/books", handlers.GetBooks)
	e.GET("/books/:id", handlers.GetBook)
	e.POST("/books", handlers.CreateBook)
	e.PUT("/books/:id", handlers.EditBook)
	e.DELETE("/books/:id", handlers.DeleteBook)

	// content book managment
	e.GET("/books/:id/contents", handlers.GetContentsBook)
	e.GET("/books/:id/contents/:idc", handlers.GetContentBook)
	e.POST("/books/:id/contents", handlers.CreateContentBook)
	e.PUT("/books/:id/contents/:idc", handlers.EditContentBook)
	e.DELETE("/books/:id/contents/:idc", handlers.DeleteContentBook)
}
