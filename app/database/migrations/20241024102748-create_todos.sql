
-- +migrate Up
CREATE TABLE IF NOT EXISTS todos(
	id INT NOT NULL PRIMARY KEY,
	user_id INT NOT NULL,
	title VARCHAR(255) NOT NULL,
	content TEXT,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	index index_user_id (user_id)
);

-- +migrate Down
DROP TABLE IF EXISTS todos;
