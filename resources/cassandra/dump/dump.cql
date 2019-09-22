CREATE KEYSPACE IF NOT EXISTS heroesbook WITH replication = {'class':'SimpleStrategy','replication_factor':'1'};
CREATE TABLE heroesbook.salts (
	pseudo text,
	salt blob,
	PRIMARY KEY (pseudo)
);
