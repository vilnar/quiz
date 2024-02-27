CREATE TABLE person (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	last_name VARCHAR(255) NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	patronymic VARCHAR(255) NOT NULL,
	military_name VARCHAR(255) NOT NULL,
	age INTEGER NOT NULL,
	gender VARCHAR(255) NOT NULL,
	unit VARCHAR(255) NOT NULL,
	specialty VARCHAR(255) NOT NULL,
	create_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_person_last_name ON person (last_name);
CREATE INDEX idx_person_first_name ON person (first_name);
CREATE INDEX idx_person_patronymic ON person (patronymic);
CREATE INDEX idx_person_unit ON person (unit);
