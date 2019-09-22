package dto

// postgres tables
const (
	TableUser             = "h_user"
	TableNode             = "h_node"
	TableBoard            = "h_board"
	TableBook             = "h_book"
	TableRemovedBook      = "h_removed_book"
	TableAuthor           = "h_author"
	TableRelationNode     = "h_relation_node"
	TableConditionnalNode = "h_conditionnal_node"
)

// User's keys from user table
const (
	UserID        = "id"
	UserPseudo    = "pseudo"
	UserLastName  = "last_name"
	UserFirstName = "first_name"
	UserHash      = "hashpass"
	UserAge       = "age"
	UserGenre     = "genre"
	UserEMail     = "email"
)

// Node's keys from node table
const (
	NodeID          = "id"
	NodeTitle       = "title"
	NodeDescription = "description"
	NodeOwner       = "owner_id"
	NodeContent     = "content"
	NodeLabels      = "labels"
)

// Board's keys from board table
const (
	BoardID    = "id"
	BoardLabel = "labels"
)

// Book's keys from book table
const (
	BookID           = "id"
	BookTitle        = "title"
	BookDescription  = "description"
	BookGenre        = "genre"
	BookPublish      = "publish"
	BookOwner        = "owner_id"
	BookNode         = "node_id"
	BookCreationDate = "creation_date"
	BookBoard        = "board_id"
)

// Book's keys from removed_book table
const (
	RBookID           = "id"
	RBookTitle        = "title"
	RBookDescription  = "description"
	RBookGenre        = "genre"
	RBookPublish      = "publish"
	RBookOwner        = "owner_id"
	RBookNode         = "node_id"
	RBookCreationDate = "creation_date"
	RBookBoard        = "board_id"
)

// Author's keys from author table
const (
	AuthorBook = "book_id"
	AuthorUser = "user_id"
)

// Relation nodes's key from relation_node table
const (
	RelationNodeParent = "parent_node"
	RelationNodeSource = "source_node"
)

// Conditionnal nodes's key from conditionnal_node table
const (
	ConditionnalNodeDest  = "destination_node"
	ConditionnalNodeCheck = "check_node"
)
