-- +goose Up
CREATE TABLE feeds (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name TEXT NOT NULl,
	url TEXT NOT NULL UNIQUE ,
	user_id UUID NOT NULl REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE feeds;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
