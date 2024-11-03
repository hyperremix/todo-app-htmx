package handlers

import (
	"context"
	"strconv"

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
	e.GET("/todos/:id", getTodo)
	e.PUT("/todos/:id", updateTodo)
}

func getTodos(echoCtx echo.Context) error {
	dbCtx := context.Background()
	conn, err := pgx.Connect(dbCtx, environment.DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer conn.Close(dbCtx)

	queries := db.New(conn)
	return renderTodos(echoCtx, dbCtx, queries)
}

func getTodo(echoCtx echo.Context) error {
	dbCtx := context.Background()
	conn, err := pgx.Connect(dbCtx, environment.DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer conn.Close(dbCtx)

	queries := db.New(conn)

	id, err := strconv.ParseInt(echoCtx.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	todoRow, err := queries.GetTodoById(dbCtx, id)
	if err != nil {
		return err
	}

	todo := services.MapRowToTodo(todoRow)
	isHxRequest := echoCtx.Request().Header.Get("HX-Request")

	if isHxRequest == "true" {
		return services.Render(echoCtx, partials.TodoPartial(todo))
	}

	return services.Render(echoCtx, pages.TodoBase(todo))
}

func createTodo(echoCtx echo.Context) error {
	dbCtx := context.Background()
	conn, err := pgx.Connect(dbCtx, environment.DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer conn.Close(dbCtx)

	queries := db.New(conn)

	todo, err := services.MapRequestToTodo(echoCtx)
	if err != nil {
		return err
	}

	_, err = queries.InsertTodo(dbCtx, services.MapTodoToInsert(todo))
	if err != nil {
		return err
	}

	return renderTodos(echoCtx, dbCtx, queries)
}

func updateTodo(echoCtx echo.Context) error {
	dbCtx := context.Background()
	conn, err := pgx.Connect(dbCtx, environment.DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer conn.Close(dbCtx)

	queries := db.New(conn)

	todo, err := services.MapRequestToTodo(echoCtx)
	if err != nil {
		return err
	}

	_, err = queries.UpdateTodo(dbCtx, services.MapTodoToUpdate(todo))
	if err != nil {
		return err
	}

	return renderTodos(echoCtx, dbCtx, queries)
}

func renderTodos(echoCtx echo.Context, dbCtx context.Context, queries *db.Queries) error {
	todoRows, err := queries.ListTodos(dbCtx)
	if err != nil {
		return err
	}

	todos := services.MapRowsToTodo(todoRows)

	isHxRequest := echoCtx.Request().Header.Get("HX-Request")

	if isHxRequest == "true" {
		return services.Render(echoCtx, partials.TodosPartial(todos))
	}

	return services.Render(echoCtx, pages.TodosBase(todos))
}
