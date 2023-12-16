CREATE TABLE quiz (
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	person_id BIGINT UNSIGNED NOT NULL,
        name VARCHAR(255) NOT NULL,
        label VARCHAR(255) NOT NULL,
        answers JSON NOT NULL,
        result JSON NOT NULL,
        score INT UNSIGNED NOT NULL,
	create_at TIMESTAMP NOT NULL,
	FOREIGN KEY (person_id) REFERENCES person(id),
	PRIMARY KEY (id)
);

CREATE INDEX idx_quiz_name ON quiz (name);
