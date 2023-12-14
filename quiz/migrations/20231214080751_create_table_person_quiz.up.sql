CREATE TABLE person_quiz (
	id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	person_id BIGINT UNSIGNED NOT NULL,
        quiz_id BIGINT UNSIGNED NOT NULL,
        name VARCHAR(255) NOT NULL,
	create_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL,
	FOREIGN KEY (person_id) REFERENCES person(id),
	PRIMARY KEY (id)
);

CREATE INDEX idx_person_quiz_id ON person_quiz (quiz_id);
CREATE INDEX idx_person_quiz_name ON person_quiz (name);
