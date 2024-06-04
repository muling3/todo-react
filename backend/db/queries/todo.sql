-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: GetUuserTodo :one
SELECT * FROM todos
WHERE user_id = ? AND id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id;

-- name: ListUserTodos :many
SELECT * FROM todos
WHERE user_id = ?
ORDER BY id;

-- name: CreateTodo :exec
INSERT INTO todos (
  user_id, title, body, priority, due_date
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: UpdateTodo :exec
UPDATE todos SET body = ?, priority = ? WHERE id = ?;

-- name: UpdateUserTodo :exec
UPDATE todos SET body = ?, priority = ? WHERE user_id = ? AND id = ?;

-- name: DeleteUserTodo :exec
DELETE FROM todos
WHERE user_id = ? AND id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;