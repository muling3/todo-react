-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: LoginUser :one
SELECT * FROM users
WHERE username = ? AND password = ? LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :exec
INSERT INTO users (
  username, password
) VALUES (
  ?, ?
);

-- name: UpdateUser :exec
UPDATE users SET username = ?, password = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;