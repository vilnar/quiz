CREATE TABLE person (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	last_name VARCHAR(255) NOT NULL COLLATE NOCASE,
	first_name VARCHAR(255) NOT NULL COLLATE NOCASE,
	patronymic VARCHAR(255) NOT NULL COLLATE NOCASE,
	military_name VARCHAR(255) NOT NULL,
	age INT,
	gender VARCHAR(255),
	unit VARCHAR(255) COLLATE NOCASE,
	specialty VARCHAR(255),
	create_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_person_last_name ON person (last_name COLLATE NOCASE);
CREATE INDEX idx_person_first_name ON person (first_name COLLATE NOCASE);
CREATE INDEX idx_person_patronymic ON person (patronymic COLLATE NOCASE);
CREATE INDEX idx_person_unit ON person (unit);
