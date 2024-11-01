// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: todos.queries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertTodo = `-- name: InsertTodo :one
INSERT INTO
    todos (title, description, is_completed)
VALUES ($1, $2, $3) RETURNING id, title, description, is_completed, created_at, updated_at, deleted_at
`

type InsertTodoParams struct {
	Title       pgtype.Text
	Description pgtype.Text
	IsCompleted pgtype.Bool
}

func (q *Queries) InsertTodo(ctx context.Context, arg InsertTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, insertTodo, arg.Title, arg.Description, arg.IsCompleted)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsCompleted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, description, is_completed, created_at, updated_at, deleted_at FROM todos ORDER BY created_at DESC
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.Query(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.IsCompleted,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE
    todos
SET
    title = $1,
    description = $2,
    is_completed = $3,
    updated_at = NOW()
WHERE
    id = $4 RETURNING id, title, description, is_completed, created_at, updated_at, deleted_at
`

type UpdateTodoParams struct {
	Title       pgtype.Text
	Description pgtype.Text
	IsCompleted pgtype.Bool
	ID          int64
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateTodo,
		arg.Title,
		arg.Description,
		arg.IsCompleted,
		arg.ID,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsCompleted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
