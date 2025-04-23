-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
	INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *
)
SELECT inserted_feed_follow.*, users.name AS user_name, feeds.name AS feed_name
FROM inserted_feed_follow
JOIN users ON inserted_feed_follow.user_id = users.id
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowForUser :many
SELECT feed_follows.*, users.name AS creator, feeds.name AS feedName FROM feed_follows
JOIN feeds ON feed_follows.feed_id = feeds.id
JOIN users ON feeds.user_id = users.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.user_id = $1 AND feeds.url = $2;