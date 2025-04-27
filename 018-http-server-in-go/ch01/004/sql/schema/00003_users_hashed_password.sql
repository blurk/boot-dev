-- +goose Up
ALTER TABLE users
ADD COLUMN hashed_password TEXT NOT NULL DEFAULT 'unset';
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE users
DROP COLUMN hashed_password;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
