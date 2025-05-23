-- +goose Up
CREATE TABLE feed_follows (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id UUID NOT NULl REFERENCES users(id) ON DELETE CASCADE,
	feed_id UUID NOT NULl REFERENCES feeds(id) ON DELETE CASCADE,
  UNIQUE(user_id, feed_id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE feed_follows;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
