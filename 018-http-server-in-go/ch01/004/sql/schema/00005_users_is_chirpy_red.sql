-- +goose Up
ALTER TABLE users
ADD COLUMN is_chirpy_red BOOLEAN DEFAULT false;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE users
DROP COLUMN is_chirpy_red;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
