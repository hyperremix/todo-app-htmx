-- name: ListTodos :many
SELECT * FROM todos ORDER BY created_at DESC;
