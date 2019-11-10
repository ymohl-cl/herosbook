package controller

import (
	"fmt"

	"github.com/ymohl-cl/herosbook/pkg/dto"
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
	"github.com/ymohl-cl/herosbook/pkg/xerror"
	"golang.org/x/xerrors"
)

// RecordNode to the book id
// check if user is book's owner
// check if title category is not used to this book
func (c controller) RecordNode(node model.Node, userID, bookID string) (model.Node, xerror.Xerror) {
	var err error
	var querySQL postgres.Query
	var tx postgres.Transaction
	var row postgres.ScanRow
	var str string

	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	defer tx.Rollback()

	// check which user is book's owner
	str = `SELECT id FROM h_book WHERE id = $1 AND owner_id = $2`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	check := ""
	if err = row.Scan(&check); err != nil {
		return model.Node{}, catchErr(err)
	}

	// create node
	str = `INSERT INTO h_node(title, description, book_id, content) VALUES($1, $2, $3, $4) RETURNING id`
	if querySQL, err = postgres.NewQuery(str,
		node.Title,
		node.Description,
		bookID,
		node.Content,
	); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	if err = row.Scan(&node.Identifier); err != nil {
		return model.Node{}, catchErr(err)
	}

	// link node and category and check than categories already exists
	for _, categoryID := range node.CategoryIDS {
		str = `INSERT INTO h_link_node_category(id_category, id_node)
			SELECT id, '` + node.Identifier + `' FROM h_category WHERE id = $1 AND book_id = $2`
		if querySQL, err = postgres.NewQuery(str,
			categoryID,
			bookID,
		); err != nil {
			return model.Node{}, newInternalErr(err)
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return model.Node{}, newInternalErr(err)
		}
		if nbAffectedRow != 1 {
			return model.Node{}, newNoContentErr(xerrors.New(fmt.Sprintf("number rows affected by linked node and category in create node request: %d", nbAffectedRow)))
		}
	}
	tx.Commit()
	return node, nil
}

// UpdateNode and update the categories list attached
// check if user is book's owner
// check if title category is not used to this book
func (c controller) UpdateNode(node model.Node, userID, bookID string) xerror.Xerror {
	var err error
	var xerr xerror.Xerror
	var querySQL postgres.Query
	var tx postgres.Transaction
	var row postgres.ScanRow
	var str string

	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return newInternalErr(err)
	}
	defer tx.Rollback()
	// check which user is book's owner
	str = `SELECT id FROM h_book WHERE id = $1 AND owner_id = $2`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return newInternalErr(err)
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return newInternalErr(err)
	}
	check := ""
	if err = row.Scan(&check); err != nil {
		return catchErr(err)
	}

	// update node
	str = `UPDATE h_node SET title = $1,
		description = $2,
		content = $3 WHERE book_id = $4 AND id = $5`
	if querySQL, err = postgres.NewQuery(str,
		node.Title,
		node.Description,
		node.Content,
		bookID,
		node.Identifier,
	); err != nil {
		return newInternalErr(err)
	}
	nbAffectedRow := int64(0)
	if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
		return newInternalErr(err)
	}
	if nbAffectedRow != 1 {
		return newNoContentErr(xerrors.New(fmt.Sprintf("number rows affected by the update node request: %d", nbAffectedRow)))
	}
	if xerr = updateCategoriesNode(tx, node, bookID); xerr != nil {
		return xerr
	}
	if xerr = updateRelationNodes(tx, node.ChildIDS, node.Identifier, bookID, dto.TypeLink); xerr != nil {
		return xerr
	}
	if xerr = updateRelationNodes(tx, node.ContionnalIDS, node.Identifier, bookID, dto.TypeConditionnal); xerr != nil {
		return xerr
	}
	tx.Commit()
	return nil
}

func updateCategoriesNode(tx postgres.Transaction, node model.Node, bookID string) xerror.Xerror {
	var err error
	var querySQL postgres.Query
	var rows postgres.ScanRows
	var str string

	// update links node categories
	// get all categories
	str = `SELECT id_category FROM h_link_node_category WHERE id_node = $1`
	if querySQL, err = postgres.NewQuery(str,
		node.Identifier,
	); err != nil {
		return newInternalErr(err)
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return newInternalErr(err)
	}
	defer rows.Close()
	catIDS := []string{}
	str = ""
	ok := true
	for ok {
		if ok, err = rows.Next(&str); err != nil {
			return newInternalErr(err)
		}
		if ok {
			catIDS = append(catIDS, str)
		}
	}
	// browse the current node category's ids and remove from categories list, id which matches
	// remove the rest of catIDS
	// add the id not found in the current link
	newLink := []string{}
	for _, id := range node.CategoryIDS {
		find := false
		for index, catID := range catIDS {
			if id == catID {
				catIDS = append(catIDS[:index], catIDS[index+1:]...)
				find = true
				break
			}
		}
		if !find {
			newLink = append(newLink, id)
		}
	}
	// add new link
	for _, id := range newLink {
		str = `INSERT INTO h_link_node_category(id_category, id_node)
			SELECT id, $3 FROM h_category WHERE id = $1 AND book_id = $2`
		if querySQL, err = postgres.NewQuery(str,
			id,
			bookID,
			node.Identifier,
		); err != nil {
			return newInternalErr(err)
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return newInternalErr(err)
		}
		if nbAffectedRow != 1 {
			return newNoContentErr(xerrors.New(fmt.Sprintf("number rows affected add link category node: %d", nbAffectedRow)))
		}
	}
	// remove the rest of catIDS
	for _, catID := range catIDS {
		str = `DELETE FROM h_link_node_category WHERE id_node = $1 AND id_category = $2`
		if querySQL, err = postgres.NewQuery(str,
			node.Identifier,
			catID,
		); err != nil {
			return newInternalErr(err)
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return newInternalErr(err)
		}
		if nbAffectedRow != 1 {
			return newNoContentErr(xerrors.New(fmt.Sprintf("number rows affected by remove link category node: %d", nbAffectedRow)))
		}
	}
	return nil
}

func updateRelationNodes(tx postgres.Transaction, ids []string, nodeID, bookID, typeRelation string) xerror.Xerror {
	var err error
	var querySQL postgres.Query
	var rows postgres.ScanRows
	var str string

	// get all nodes
	str = `SELECT source_node FROM h_relation_node WHERE parent_node = $1 AND relation_type = $2`
	if querySQL, err = postgres.NewQuery(str,
		nodeID,
		typeRelation,
	); err != nil {
		return newInternalErr(err)
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return newInternalErr(err)
	}
	defer rows.Close()
	sourceIDS := []string{}
	for ok := true; ok; {
		str := ""
		if ok, err = rows.Next(&str); err != nil {
			return newInternalErr(err)
		}
		if ok {
			sourceIDS = append(sourceIDS, str)
		}
	}
	// browse currentIDS node list and remove from sourceIDS, id which matches
	// next, remove the rest of sourceIDS
	// add the id not found in the new relation link describe by typeRelation
	newLink := []string{}
	for _, id := range ids {
		find := false
		for index, sourceID := range sourceIDS {
			if id == sourceID {
				sourceIDS = append(sourceIDS[:index], sourceIDS[index+1:]...)
				find = true
				break
			}
		}
		if !find {
			newLink = append(newLink, id)
		}
	}
	for _, id := range newLink {
		str = `INSERT INTO h_relation_node(parent_node, source_node, relation_type)
			SELECT $1, id, $4 FROM h_node WHERE id = $2 AND book_id = $3`
		if querySQL, err = postgres.NewQuery(str,
			nodeID,
			id,
			bookID,
			typeRelation,
		); err != nil {
			return newInternalErr(err)
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return newInternalErr(err)
		}
		if nbAffectedRow != 1 {
			return newNoContentErr(xerrors.New(fmt.Sprintf("number rows affected to add links relation node: %d", nbAffectedRow)))
		}
	}
	// remove the rest of source IDS because is deprecated
	for _, id := range sourceIDS {
		str = `DELETE FROM h_relation_node WHERE source_node = $1 AND parent_node = $2 AND relation_type = $3`
		if querySQL, err = postgres.NewQuery(str,
			id,
			nodeID,
			typeRelation,
		); err != nil {
			return newInternalErr(err)
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return newInternalErr(err)
		}
		if nbAffectedRow != 1 {
			return newNoContentErr(xerrors.New(fmt.Sprintf("number rows affected to remove links relation node: %d", nbAffectedRow)))
		}
	}
	return nil
}

// ReadNode by id
func (c controller) ReadNode(nodeID, bookID, userID string) (model.Node, xerror.Xerror) {
	var node model.Node
	var querySQL postgres.Query
	var tx postgres.Transaction
	var err error
	var xerr xerror.Xerror
	var row postgres.ScanRow
	var rows postgres.ScanRows

	node.Identifier = nodeID
	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	defer tx.Rollback()
	// content node
	str := `SELECT title, description, content FROM h_node WHERE id = $1 AND book_id = $2 AND exists(SELECT id FROM h_book WHERE id = $2 AND owner_id = $3)`
	if querySQL, err = postgres.NewQuery(str,
		nodeID,
		bookID,
		userID,
	); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	if err = row.Scan(&node.Title, &node.Description, &node.Content); err != nil {
		return model.Node{}, catchErr(err)
	}
	// categories list
	str = `SELECT id_category FROM h_link_node_category WHERE id_node = $1`
	if querySQL, err = postgres.NewQuery(str, nodeID); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return model.Node{}, newInternalErr(err)
	}
	defer rows.Close()
	ok := true
	for ok {
		id := ""
		if ok, err = rows.Next(&id); err != nil {
			return model.Node{}, newInternalErr(err)
		}
		if ok {
			node.CategoryIDS = append(node.CategoryIDS, id)
		}
	}
	// child node list
	if node.ChildIDS, xerr = getRelationNode(tx, nodeID, dto.TypeLink); xerr != nil {
		return model.Node{}, xerr
	}
	// conditionnal node list
	if node.ContionnalIDS, xerr = getRelationNode(tx, nodeID, dto.TypeConditionnal); xerr != nil {
		return model.Node{}, xerr
	}

	tx.Commit()
	return node, nil
}

// Return slice of ids to the relation with the node
func getRelationNode(tx postgres.Transaction, nodeID, relationType string) ([]string, xerror.Xerror) {
	var querySQL postgres.Query
	var err error
	var rows postgres.ScanRows
	var str string
	var result []string

	str = `SELECT source_node FROM h_relation_node WHERE parent_node = $1 AND relation_type = $2`
	if querySQL, err = postgres.NewQuery(str,
		nodeID,
		relationType,
	); err != nil {
		return nil, newInternalErr(err)
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return nil, newInternalErr(err)
	}
	defer rows.Close()

	ok := true
	for ok {
		id := ""
		if ok, err = rows.Next(&id); err != nil {
			return nil, newInternalErr(err)
		}
		if ok {
			result = append(result, id)
		}
	}
	return result, nil
}

// ReadNodes to the bookID given and return list nodes ids
func (c controller) ReadNodes(bookID, userID string) ([]string, xerror.Xerror) {
	var querySQL postgres.Query
	var rows postgres.ScanRows
	var err error

	ids := []string{}
	str := `SELECT id FROM h_node WHERE book_id = $1 AND exists(SELECT id FROM h_book WHERE id = $1 AND owner_id = $2)`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return nil, newInternalErr(err)
	}
	if rows, err = c.driverSQL.WithRows(querySQL); err != nil {
		return nil, newInternalErr(err)
	}
	defer rows.Close()

	ok := true
	for ok {
		id := ""
		if ok, err = rows.Next(&id); err != nil {
			return nil, newInternalErr(err)
		}
		if ok {
			ids = append(ids, id)
		}
	}
	return ids, nil
}

// DeleteNode from the database
func (c controller) DeleteNode(nodeID, userID, bookID string) xerror.Xerror {
	var err error
	var querySQL postgres.Query
	var str string
	var nbAffectedRow int64

	str = `DELETE FROM h_node WHERE
		id = $1 AND
		book_id = $2 AND
		exists(SELECT id FROM h_book WHERE id = $2 AND owner_id = $3)`
	if querySQL, err = postgres.NewQuery(str,
		nodeID,
		bookID,
		userID,
	); err != nil {
		return newInternalErr(err)
	}
	if nbAffectedRow, err = c.driverSQL.WithNoRow(querySQL); err != nil {
		return newInternalErr(err)
	}
	if nbAffectedRow != 1 {
		return newNoContentErr(xerrors.New(fmt.Sprintf("number rows affected by the delete node request: %d", nbAffectedRow)))
	}
	return nil
}
