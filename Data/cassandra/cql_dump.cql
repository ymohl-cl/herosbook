CREATE KEYSPACE IF NOT EXISTS herosbook WITH replication = {'class':'SimpleStrategy','replication_factor':'1'};
CREATE TABLE herosbook.users (
	id_psql uuid,
	salt blob,
	PRIMARY KEY (id_psql)
);
CREATE TABLE herosbook.tokens (
	id_psql uuid,
	value blob,
	PRIMARY KEY (id_psql)
);
