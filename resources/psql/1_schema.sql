CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE h_user
(
	id UUID						DEFAULT uuid_generate_v4(),
	pseudo VARCHAR (56)			UNIQUE NOT NULL,
	last_name VARCHAR(255)  	NOT NULL,
	first_name VARCHAR(255) 	NOT NULL,
	hashpass BYTEA				NOT NULL,
	age INT						NOT NULL,
	genre VARCHAR(10)			NOT NULL,
	email VARCHAR(245)			UNIQUE NOT NULL,
	PRIMARY KEY (id)
);
CREATE TABLE h_node
(
	id UUID						DEFAULT uuid_generate_v4(),
	title VARCHAR(255)			,
	description TEXT			,
	owner_id UUID				REFERENCES h_user NOT NULL,
	content TEXT				,
	labels TEXT[]				,
	PRIMARY KEY (id)
);
CREATE TABLE h_board
(
	id UUID						DEFAULT uuid_generate_v4(),
	labels TEXT[]				,
	PRIMARY KEY (id)				
);
CREATE TABLE h_book
(
	id UUID						DEFAULT uuid_generate_v4(),
	title VARCHAR(255)			UNIQUE NOT NULL,
	description TEXT			NOT NULL,
	genre VARCHAR(255) 			,
	publish BOOLEAN				NOT NULL,
	owner_id UUID 				REFERENCES h_user NOT NULL,
	node_id UUID				REFERENCES h_node,
	creation_date TIMESTAMP		NOT NULL,
	board_id UUID				REFERENCES h_board NOT NULL,
	PRIMARY KEY (id)
);
CREATE TABLE h_removed_book
(
	id UUID						DEFAULT uuid_generate_v4(),
	title VARCHAR(255)			UNIQUE NOT NULL,
	describe TEXT				NOT NULL,
	genre VARCHAR(255) 			,
	publish BOOLEAN				NOT NULL,
	owner_id UUID 				REFERENCES h_user NOT NULL,
	node_id UUID				REFERENCES h_node,
	creation_date TIMESTAMP		NOT NULL,
	board_id UUID				REFERENCES h_board,
	PRIMARY KEY (id)
);
CREATE TABLE h_relation_node
(
	parent_node UUID			REFERENCES h_node NOT NULL,
	source_node UUID			REFERENCES h_node NOT NULL
);
CREATE TABLE h_conditionnal_node
(
	destination_node UUID		REFERENCES h_node NOT NULL,
	check_node UUID				REFERENCES h_node NOT NULL
);
CREATE TABLE h_category
(
	id UUID						DEFAULT uuid_generate_v4(),
	name_category VARCHAR(255)	NOT NULL,
	title VARCHAR(255)			NOT NULL,
	description TEXT,
	PRIMARY KEY (id)
);
CREATE TABLE h_link_book_category
(
	id_category UUID			REFERENCES h_category ON DELETE CASCADE NOT NULL,
	id_book	UUID				REFERENCES h_book ON DELETE CASCADE NOT NULL
)