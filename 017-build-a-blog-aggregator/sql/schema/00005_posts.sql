-- +goose Up
CREATE TABLE posts (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULl,
	updated_at TIMESTAMP NOT NULL,
	title TEXT NOT NULL,
	url TEXT NOT NULL UNIQUE,
	description TEXT,
	published_at TIMESTAMP,
	feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE posts;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
