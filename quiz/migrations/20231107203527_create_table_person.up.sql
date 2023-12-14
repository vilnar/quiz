CREATE TABLE person (
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	full_name VARCHAR(255) NOT NULL,
	military_name VARCHAR(255) NOT NULL,
	age INT,
	gender VARCHAR(255),
	unit VARCHAR(255),
	specialty VARCHAR(255),
	create_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL,
	PRIMARY KEY (id),
        FULLTEXT(full_name)
);

