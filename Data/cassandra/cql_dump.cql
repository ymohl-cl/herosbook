CREATE KEYSPACE IF NOT EXISTS herosbook WITH replication = {'class':'SimpleStrategy','replication_factor':'1'};
CREATE TABLE herosbook.users (
	id text,
	id_psql text,
	salt int,
	PRIMARY KEY (id)
);
