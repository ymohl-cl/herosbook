package dto

// postgres tables
const (
	TableUser             = "h_user"
	TableBook             = "h_book"
	TableNode             = "h_node"
	TableCategory         = "h_category"
	TableRelationNode     = "h_relation_node"
	TableLinkNodeCategory = "h_link_node_category"
)

// relation node detail
// Conditionnal attach node to access parent's node (source must be read to read parent)
// Link attach node to access parent's node (source is the next node after read parent)
const (
	TypeConditionnal = "conditionnal"
	TypeLink         = "link"
)
