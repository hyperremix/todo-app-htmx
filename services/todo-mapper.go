package services

import (
	"github.com/hyperremix/todo-app-htmx/db"
	"github.com/hyperremix/todo-app-htmx/model"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func MapRowsToTodo(row []db.Todo) []model.Todo {
	seasons := make([]model.Todo, len(row))
	for i := range row {
		seasons[i] = MapRowToTodo(row[i])
	}
	return seasons
}

func MapRowToTodo(row db.Todo) model.Todo {
	return model.Todo{
		ID:          row.ID,
		Title:       row.Title.String,
		Description: row.Description.String,
		IsCompleted: row.IsCompleted.Bool,
	}
}

func MapRequestToTodo(c echo.Context) (model.Todo, error) {
	todo := model.Todo{}

	err := c.Bind(&todo)
	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func MapTodoToInsert(todo model.Todo) db.InsertTodoParams {
	return db.InsertTodoParams{
		Title:       pgtype.Text{String: todo.Title, Valid: true},
		Description: pgtype.Text{String: todo.Description, Valid: true},
		IsCompleted: pgtype.Bool{Bool: todo.IsCompleted, Valid: true},
	}
}
