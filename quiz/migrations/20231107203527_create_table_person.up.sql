CREATE TABLE person (
	id BIGINT NOT NULL AUTO_INCREMENT,
	full_name varchar(255),
	military_name varchar(255),
	age int,
	gender varchar(255),
	unit varchar(255),
	specialty varchar(255),
	create_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL,
	PRIMARY KEY (id)
);
