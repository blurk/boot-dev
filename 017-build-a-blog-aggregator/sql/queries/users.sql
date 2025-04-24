-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2
)
RETURNING *;
-- $1, $2, $3, and $4 are parameters that we'll be able to pass into the query in our Go code. The :one at the end of the query name tells SQLC that we expect to get back a single row (the created user).

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;