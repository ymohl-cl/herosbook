package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/herosbook/pkg/app/auth"
	"github.com/ymohl-cl/herosbook/pkg/model"
)

// CreateNode record the node to the book id
func (h Handler) CreateNode(c echo.Context) error {
	var err error
	var user auth.User
	var node model.Node

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	if err = h.jsonValidator.Bind(c.Request(), &node); err != nil {
		fmt.Printf("CreateNode - jsonvalidator.Bind - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	if node, err = h.controller.RecordNode(node, user.Identifier, bookID); err != nil {
		fmt.Printf("CreateNode - h.controller.RecordNode - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	return c.JSON(http.StatusCreated, &node)
}

// UpdateNode to the book id
func (h Handler) UpdateNode(c echo.Context) error {
	var err error
	var user auth.User
	var node model.Node

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	if err = h.jsonValidator.Bind(c.Request(), &node); err != nil {
		fmt.Printf("UpdateNode - jsonvalidator.Bind - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}
	if err = h.controller.UpdateNode(node, user.Identifier, bookID); err != nil {
		fmt.Printf("UpdateNode - h.controller.UpdateNode - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	return c.NoContent(http.StatusOK)
}

// GetNode to the book id with the node id given
func (h Handler) GetNode(c echo.Context) error {
	var err error
	var user auth.User
	var node model.Node

	bookID := c.Param("id")
	nodeID := c.Param("id_node")
	user = auth.ParseToken(c)
	if node, err = h.controller.ReadNode(nodeID, bookID, user.Identifier); err != nil {
		fmt.Printf("GetNode - h.controller.ReadNode - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	if node.Identifier == "" {
		return c.JSON(http.StatusNoContent, map[string]string{
			"message": "node not found",
		})
	}
	return c.JSON(http.StatusOK, &node)
}

// GetNodes to the book id
func (h Handler) GetNodes(c echo.Context) error {
	var err error
	var user auth.User
	var nodeIDS []string

	bookID := c.Param("id")
	user = auth.ParseToken(c)
	if nodeIDS, err = h.controller.ReadNodes(bookID, user.Identifier); err != nil {
		fmt.Printf("GetNode - h.controller.ReadNode - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	if len(nodeIDS) == 0 {
		return c.JSON(http.StatusNoContent, map[string]string{
			"message": "node not found",
		})
	}
	return c.JSON(http.StatusOK, &nodeIDS)
}

// RemoveNode to the book id
func (h Handler) RemoveNode(c echo.Context) error {
	var err error
	var user auth.User

	user = auth.ParseToken(c)
	bookID := c.Param("id")
	nodeID := c.Param("id_node")
	if err = h.controller.DeleteNode(nodeID, user.Identifier, bookID); err != nil {
		fmt.Printf("RemoveNode - h.controller.DeleteNode - %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "internal error",
		})
	}
	return c.NoContent(http.StatusOK)

}
