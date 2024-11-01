package handlers

import (
	"context"

	"github.com/hyperremix/todo-app-htmx/components/corecomponents"
	"github.com/hyperremix/todo-app-htmx/components/pages"
	"github.com/hyperremix/todo-app-htmx/components/partials"
	"github.com/hyperremix/todo-app-htmx/db"
	"github.com/hyperremix/todo-app-htmx/environment"
	"github.com/hyperremix/todo-app-htmx/services"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func DefineTodoRoutes(e *echo.Echo) {
	e.GET("/", getTodos)
	e.POST("/todos", createTodo)
}

func getTodos(c echo.Context) error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, environment.DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)
	todoRows, err := queries.ListTodos(ctx)
	if err != nil {
		return err
	}

	todos := services.MapRowsToTodo(todoRows)

	isHxRequest := c.Request().Header.Get("HX-Request")

	if isHxRequest == "true" {
		return services.Render(c, partials.TodosPartial(todos))
	}

	return services.Render(c, pages.TodosBase(todos))
}

func createTodo(c echo.Context) error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, environment.DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	todoInsert, err := services.MapRequestToTodo(c)
	if err != nil {
		return err
	}

	todoRow, err := queries.InsertTodo(ctx, services.MapTodoToInsert(todoInsert))
	if err != nil {
		return err
	}

	return services.Render(c, corecomponents.Todo(corecomponents.TodoProps{Todo: services.MapRowToTodo(todoRow)}))
}
