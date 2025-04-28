-- +goose Up
CREATE TABLE chirps (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	body TEXT NOT NULL,
	user_id UUID NOT NULl REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE chirps;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
