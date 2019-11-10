package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/pkg/app/auth"
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/xerror"
)

// CreateNode record the node to the book id
func (h Handler) CreateNode(c echo.Context) error {
	var err error
	var xerr xerror.Xerror
	var user auth.User
	var node model.Node

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	if err = h.jsonValidator.Bind(c.Request(), &node); err != nil {
		return c.JSON(h.httpResponse(xerror.New(-1, err.Error())))
	}
	if node, xerr = h.controller.RecordNode(node, user.Identifier, bookID); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.JSON(http.StatusCreated, &node)
}

// UpdateNode to the book id
func (h Handler) UpdateNode(c echo.Context) error {
	var err error
	var xerr xerror.Xerror
	var user auth.User
	var node model.Node

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	if err = h.jsonValidator.Bind(c.Request(), &node); err != nil {
		return c.JSON(h.httpResponse(xerror.New(-1, err.Error())))
	}
	if xerr = h.controller.UpdateNode(node, user.Identifier, bookID); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.NoContent(http.StatusOK)
}

// GetNode to the book id with the node id given
func (h Handler) GetNode(c echo.Context) error {
	var user auth.User
	var node model.Node
	var xerr xerror.Xerror

	bookID := c.Param("id")
	nodeID := c.Param("id_node")
	user = auth.ParseToken(c)
	if node, xerr = h.controller.ReadNode(nodeID, bookID, user.Identifier); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.JSON(http.StatusOK, &node)
}

// GetNodes to the book id
func (h Handler) GetNodes(c echo.Context) error {
	var xerr xerror.Xerror
	var user auth.User
	var nodeIDS []string

	bookID := c.Param("id")
	user = auth.ParseToken(c)
	if nodeIDS, xerr = h.controller.ReadNodes(bookID, user.Identifier); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.JSON(http.StatusOK, &nodeIDS)
}

// RemoveNode to the book id
func (h Handler) RemoveNode(c echo.Context) error {
	var xerr xerror.Xerror
	var user auth.User

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	nodeID := c.Param("id_node")
	if xerr = h.controller.DeleteNode(nodeID, user.Identifier, bookID); xerr != nil {
		return c.JSON(h.httpResponse(xerr))
	}
	return c.NoContent(http.StatusOK)

}
