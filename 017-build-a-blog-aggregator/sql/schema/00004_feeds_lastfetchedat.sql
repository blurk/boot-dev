-- +goose Up
ALTER TABLE feeds
ADD COLUMN last_fetched_at TIMESTAMP;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE feeds
DROP COLUMN last_fetched_at;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
