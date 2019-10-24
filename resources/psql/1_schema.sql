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
CREATE TABLE h_book
(
	id UUID						DEFAULT uuid_generate_v4(),
	title VARCHAR(255)			UNIQUE NOT NULL,
	description TEXT			NOT NULL,
	genre VARCHAR(255) 			,
	publish BOOLEAN				NOT NULL,
	owner_id UUID 				REFERENCES h_user ON DELETE CASCADE NOT NULL,
	creation_date TIMESTAMP		NOT NULL,
	PRIMARY KEY (id)
);
CREATE TABLE h_node
(
	id UUID						DEFAULT uuid_generate_v4(),
	title VARCHAR(255)			,
	description TEXT			,
	book_id UUID				REFERENCES h_book ON DELETE CASCADE NOT NULL,
	content TEXT				,
	PRIMARY KEY (id)
);
CREATE TABLE h_category
(
	id UUID						DEFAULT uuid_generate_v4(),
	name_category VARCHAR(255)	NOT NULL,
	title VARCHAR(255)			NOT NULL,
	description TEXT			,
	book_id UUID				REFERENCES h_book ON DELETE CASCADE NOT NULL,
	UNIQUE(title, book_id)		,
	PRIMARY KEY (id)
);
CREATE TABLE h_relation_node
(
	parent_node UUID			REFERENCES h_node ON DELETE CASCADE NOT NULL,
	source_node UUID			REFERENCES h_node ON DELETE CASCADE NOT NULL,
	relation_type VARCHAR(255)
);
CREATE TABLE h_link_node_category
(
	id_category UUID			REFERENCES h_category ON DELETE CASCADE NOT NULL,
	id_node UUID				REFERENCES h_node ON DELETE CASCADE NOT NULL
)