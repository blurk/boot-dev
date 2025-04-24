-- name: CreatePost :one
INSERT INTO posts (	id, created_at, updated_at, title, url, description, published_at, feed_id )
VALUES ($1, NOW(), NOW(), $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM posts
JOIN feeds ON feeds.id = posts.feed_id
WHERE feeds.user_id = $1
ORDER BY posts.created_at
LIMIT $2;