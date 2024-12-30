-- name: ListTodos :many
SELECT * FROM todos ORDER BY created_at DESC;

-- name: ListOpenTodos :many
SELECT * FROM todos WHERE is_completed = FALSE ORDER BY created_at DESC;

-- name: GetTodoById :one
SELECT * FROM todos WHERE id = $1 LIMIT 1;

-- name: InsertTodo :one
INSERT INTO
    todos (title, description, is_completed)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateTodo :one
UPDATE
    todos
SET
    title = $1,
    description = $2,
    is_completed = $3,
    updated_at = NOW()
WHERE
    id = $4 RETURNING *;