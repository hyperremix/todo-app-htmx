package handlers

import (
	"context"

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
}

func getTodos(c echo.Context) error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, environment.DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)
	_, err = queries.ListTodos(ctx)
	if err != nil {
		return err
	}

	isHxRequest := c.Request().Header.Get("HX-Request")

	if isHxRequest == "true" {
		return services.Render(c, partials.TodosPartial())
	}

	return services.Render(c, pages.TodosBase())
}
