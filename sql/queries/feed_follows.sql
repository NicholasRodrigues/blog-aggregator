-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, feed_id, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
--

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE id = $1 AND user_id = $2;
--

-- name: GetFeedFollowsByUserId :many
SELECT * FROM feed_follows
WHERE user_id = $1;
--
