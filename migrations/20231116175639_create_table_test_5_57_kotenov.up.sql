CREATE TABLE test_5_57_kotenov (
	id BIGINT NOT NULL AUTO_INCREMENT,
	user_id BIGINT,
	answers JSON,
	ptsd varchar(255),
	gsr varchar(255),
	depression varchar(255),
	conclusion text,
	create_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL,
	FOREIGN KEY (user_id) REFERENCES user(id),
	PRIMARY KEY (id)
);
