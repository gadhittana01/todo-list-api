CREATE TABLE todos(
	todo_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	activity_group_id INT,
	title VARCHAR(255),
	priority VARCHAR(255),
	is_active BOOLEAN,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(activity_group_id) REFERENCES activities(activity_id)
)