CREATE TABLE test_5_57_kotenov (
	id BIGINT NOT NULL AUTO_INCREMENT,
	user_id BIGINT,
	answers JSON,
	ptsd INT,
	gsr INT,
	depression INT,
	lie_description text,
	ptsd_description text,
	gsr_description text,
	depression_description text,
	a1 INT,
	b_ INT,
	c_ INT,
	d_ INT,
	f_ INT,
	l INT,
	ag INT,
	di INT,
	b INT,
	c INT,
	d INT,
	e INT,
	f INT,
	create_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL,
	FOREIGN KEY (user_id) REFERENCES user(id),
	PRIMARY KEY (id)
);
