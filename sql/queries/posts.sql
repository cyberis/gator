-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostByURL :one
SELECT * FROM posts
WHERE url = $1;

-- name: GetPostsForUser :many
SELECT posts.*, feeds.name AS feed_name, users.name AS user_name
FROM posts
JOIN feeds ON posts.feed_id = feeds.id
JOIN feed_follows ON feeds.id = feed_follows.feed_id
JOIN users ON feed_follows.user_id = users.id
WHERE users.name = $1
ORDER BY published_at DESC
LIMIT $2 OFFSET $3;