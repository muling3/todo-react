-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id;

-- name: CreateTodo :exec
INSERT INTO todos (
  title, body, priority, due_date
) VALUES (
  ?, ?, ?, ?
);

-- name: UpdateTodo :exec
UPDATE todos SET body = ?, priority = ? WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;