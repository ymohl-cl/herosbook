package controller

import (
	"strconv"

	"github.com/ymohl-cl/herosbook/pkg/dto"
	"github.com/ymohl-cl/herosbook/pkg/model"
	"github.com/ymohl-cl/herosbook/pkg/postgres"
	"golang.org/x/xerrors"
)

func (c controller) RecordNode(node model.Node, userID, bookID string) (model.Node, error) {
	var err error
	var querySQL postgres.Query
	var tx postgres.Transaction
	var row postgres.ScanRow
	var str string

	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return model.Node{}, err
	}
	defer tx.Rollback()

	// check which user is book's owner
	str = `SELECT id FROM h_book WHERE id = $1 AND owner_id = $2`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return model.Node{}, err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Node{}, err
	}
	check := ""
	if err = row.Scan(&check); err != nil {
		return model.Node{}, err
	}
	if check != bookID {
		return model.Node{}, xerrors.New("user can't add a category")
	}

	// create node
	str = `INSERT INTO h_node(title, description, book_id, content) VALUES($1, $2, $3, $4) RETURNING id`
	if querySQL, err = postgres.NewQuery(str,
		node.Title,
		node.Description,
		bookID,
		node.Content,
	); err != nil {
		return model.Node{}, err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Node{}, err
	}
	if err = row.Scan(&node.Identifier); err != nil {
		return model.Node{}, err
	}

	// link node and category and check than categories already exists
	for _, categoryID := range node.CategoryIDS {
		str = `INSERT INTO h_link_node_category(id_category, id_node)
			SELECT id, '` + node.Identifier + `' FROM h_category WHERE id = $1 AND book_id = $2`
		if querySQL, err = postgres.NewQuery(str,
			categoryID,
			bookID,
		); err != nil {
			return model.Node{}, err
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return model.Node{}, err
		}
		if nbAffectedRow != 1 {
			return model.Node{}, xerrors.New("linked node and category failed because rows affected: " + strconv.Itoa(int(nbAffectedRow)))
		}
	}
	tx.Commit()
	return node, nil
}

// UpdateNode and update the categories list attached
func (c controller) UpdateNode(node model.Node, userID, bookID string) error {
	var err error
	var querySQL postgres.Query
	var tx postgres.Transaction
	var row postgres.ScanRow
	var str string

	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return err
	}
	defer tx.Rollback()
	// check which user is book's owner
	str = `SELECT id FROM h_book WHERE id = $1 AND owner_id = $2`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return err
	}
	check := ""
	if err = row.Scan(&check); err != nil {
		return err
	}
	if check != bookID {
		return xerrors.New("user can't add a category")
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
		return err
	}
	nbAffectedRow := int64(0)
	if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
		return err
	}
	if nbAffectedRow != 1 {
		return xerrors.New("update node failed")
	}
	if err = updateCategoriesNode(tx, node, bookID); err != nil {
		return err
	}
	if err = updateRelationNodes(tx, node.ChildIDS, node.Identifier, bookID, dto.TypeLink); err != nil {
		return err
	}
	if err = updateRelationNodes(tx, node.ContionnalIDS, node.Identifier, bookID, dto.TypeConditionnal); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func updateCategoriesNode(tx postgres.Transaction, node model.Node, bookID string) error {
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
		return err
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return err
	}
	defer rows.Close()
	catIDS := []string{}
	str = ""
	ok := true
	for ok {
		if ok, err = rows.Next(&str); err != nil {
			return err
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
			return err
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return err
		}
		if nbAffectedRow != 1 {
			return xerrors.New("linked node and category failed")
		}
	}
	// remove the rest of catIDS
	for _, catID := range catIDS {
		str = `DELETE FROM h_link_node_category WHERE id_node = $1 AND id_category = $2`
		if querySQL, err = postgres.NewQuery(str,
			node.Identifier,
			catID,
		); err != nil {
			return err
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return err
		}
		if nbAffectedRow != 1 {
			return xerrors.New("error to remove link node and category")
		}
	}
	return nil
}

func updateRelationNodes(tx postgres.Transaction, ids []string, nodeID, bookID, typeRelation string) error {
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
		return err
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return err
	}
	defer rows.Close()
	sourceIDS := []string{}
	for ok := true; ok; {
		str := ""
		if ok, err = rows.Next(&str); err != nil {
			return err
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
			return err
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return err
		}
		if nbAffectedRow != 1 {
			return xerrors.New("linked node failed")
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
			return err
		}
		nbAffectedRow := int64(0)
		if nbAffectedRow, err = tx.WithNoRow(querySQL); err != nil {
			return err
		}
		if nbAffectedRow != 1 {
			return xerrors.New("error to remove link conditionnal node")
		}
	}
	return nil
}

// ReadNode by id
func (c controller) ReadNode(nodeID, bookID, userID string) (model.Node, error) {
	var node model.Node
	var querySQL postgres.Query
	var tx postgres.Transaction
	var err error
	var row postgres.ScanRow
	var rows postgres.ScanRows

	node.Identifier = nodeID
	if tx, err = c.driverSQL.NewTransaction(); err != nil {
		return model.Node{}, err
	}
	defer tx.Rollback()
	// content node
	str := `SELECT title, description, content FROM h_node WHERE id = $1 AND book_id = $2 AND exists(SELECT id FROM h_book WHERE id = $2 AND owner_id = $3)`
	if querySQL, err = postgres.NewQuery(str,
		nodeID,
		bookID,
		userID,
	); err != nil {
		return model.Node{}, err
	}
	if row, err = tx.WithRow(querySQL); err != nil {
		return model.Node{}, err
	}
	if err = row.Scan(&node.Title, &node.Description, &node.Content); err != nil {
		return model.Node{}, err
	}
	// categories list
	str = `SELECT id_category FROM h_link_node_category WHERE id_node = $1`
	if querySQL, err = postgres.NewQuery(str, nodeID); err != nil {
		return model.Node{}, err
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return model.Node{}, err
	}
	defer rows.Close()
	ok := true
	for ok {
		id := ""
		if ok, err = rows.Next(&id); err != nil {
			return model.Node{}, err
		}
		if ok {
			node.CategoryIDS = append(node.CategoryIDS, id)
		}
	}
	// child node list
	if node.ChildIDS, err = getRelationNode(tx, nodeID, dto.TypeLink); err != nil {
		return model.Node{}, err
	}
	// conditionnal node list
	if node.ContionnalIDS, err = getRelationNode(tx, nodeID, dto.TypeConditionnal); err != nil {
		return model.Node{}, err
	}

	tx.Commit()
	return node, nil
}

// Return slice of ids to the relation with the node
func getRelationNode(tx postgres.Transaction, nodeID, relationType string) ([]string, error) {
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
		return nil, err
	}
	if rows, err = tx.WithRows(querySQL); err != nil {
		return nil, err
	}
	defer rows.Close()

	ok := true
	for ok {
		id := ""
		if ok, err = rows.Next(&id); err != nil {
			return nil, err
		}
		if ok {
			result = append(result, id)
		}
	}
	return result, nil
}

// ReadNodes to the bookID given and return list nodes ids
func (c controller) ReadNodes(bookID, userID string) ([]string, error) {
	var ids []string
	var querySQL postgres.Query
	var rows postgres.ScanRows
	var err error

	str := `SELECT id FROM h_node WHERE book_id = $1 AND exists(SELECT id FROM h_book WHERE id = $1 AND owner_id = $2)`
	if querySQL, err = postgres.NewQuery(str,
		bookID,
		userID,
	); err != nil {
		return nil, err
	}
	if rows, err = c.driverSQL.WithRows(querySQL); err != nil {
		return nil, err
	}
	defer rows.Close()

	ok := true
	for ok {
		id := ""
		if ok, err = rows.Next(&id); err != nil {
			return nil, err
		}
		if ok {
			ids = append(ids, id)
		}
	}
	return ids, nil
}

// DeleteNode from the database
func (c controller) DeleteNode(nodeID, userID, bookID string) error {
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
		return err
	}
	if nbAffectedRow, err = c.driverSQL.WithNoRow(querySQL); err != nil {
		return err
	}
	if nbAffectedRow != 1 {
		return xerrors.New("error to update the category")
	}
	return nil
}
