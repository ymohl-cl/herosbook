package handler

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/cmd/api-editor/controller"
	"github.com/ymohl-cl/herosbook/pkg/app/jsonvalidator"
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/xerror"
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
	// node
	server.POST("/books/:id/node", h.CreateNode)
	server.GET("/books/:id/node", h.GetNodes)
	server.GET("/books/:id/node/:id_node", h.GetNode)
	server.PUT("/books/:id/node", h.UpdateNode)
	server.DELETE("/books/:id/node/:id_node", h.RemoveNode)

	return nil
}

func (h Handler) httpResponse(xerr xerror.Xerror) (int, interface{}) {
	var code int

	m := model.Message{ID: xerr.ID()}
	fmt.Println(fmt.Sprintf("{[O;O] id: %s - [stack: %s %s]}", xerr.ID(), getFormatStack(), xerr.Error()))
	switch xerr.Code() {
	case -1:
		m.Message = "invalid input, you can read the api guide on http://heroesbook/apiguide.fr"
		code = http.StatusBadRequest
	case controller.ErrDuplicateKey:
		m.Message = "duplicate resources are prohibited"
		code = http.StatusBadRequest
	case controller.ErrNoResult:
		code = http.StatusNoContent
	default:
		m.Message = fmt.Sprintf("oups, it would seem that something supernatural is happening. Are you a Heroes ? You can help us with report the next code error: %s", xerr.ID())
		code = http.StatusInternalServerError
	}
	return code, &m
}

func getFormatStack() string {
	var format string

	if _, file, line, ok := runtime.Caller(2); ok {
		format += fmt.Sprintf("in %s:%d -> ", file, line)
	}
	if _, file, line, ok := runtime.Caller(3); ok {
		format += fmt.Sprintf("in %s:%d -> ", file, line)
	}
	format += "detail error: "
	return format
}
