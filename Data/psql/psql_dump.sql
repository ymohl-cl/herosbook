CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
	id UUID					DEFAULT uuid_generate_v4(),
	id_public UUID			DEFAULT uuid_generate_v4(),
	pseudo VARCHAR (56)		UNIQUE NOT NULL,
	password BYTEA			NOT NULL,
	age INT					NOT NULL,
	sex VARCHAR(10)			NOT NULL,
	email VARCHAR(245)		UNIQUE NOT NULL,
	PRIMARY KEY (id)
)
