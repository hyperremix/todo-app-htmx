package handlers

import (
	"context"

	"github.com/hyperremix/todo-app-htmx/components/corecomponents"
	"github.com/hyperremix/todo-app-htmx/components/pages"
	"github.com/hyperremix/todo-app-htmx/components/partials"
	"github.com/hyperremix/todo-app-htmx/db"
	"github.com/hyperremix/todo-app-htmx/environment"
	"github.com/hyperremix/todo-app-htmx/model"
	"github.com/hyperremix/todo-app-htmx/services"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func DefineTodoRoutes(e *echo.Echo) {
	e.GET("/", getTodos)
	e.POST("/", createTodo)
	e.PUT("/:id", updateTodo)
}

func getTodos(echoCtx echo.Context) error {
	dbCtx := context.Background()
	conn, err := pgx.Connect(dbCtx, environment.DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer conn.Close(dbCtx)

	queries := db.New(conn)

	isHxRequest := echoCtx.Request().Header.Get("HX-Request")
	var request model.GetTodosRequest
	if err := echoCtx.Bind(&request); err != nil {
		return err
	}

	todoId, isModalVisible := services.MapGetTodosRequest(request)
	if isHxRequest == "true" {
		if request.IsCreateModalVisible {
			return services.Render(echoCtx, corecomponents.Modal(corecomponents.ModalProps{Content: partials.TodoForm(partials.TodoFormProps{Todo: model.Todo{}}), IsModalVisible: isModalVisible}))
		}

		if request.IsUpdateModalVisible {
			todoRow, err := queries.GetTodoById(dbCtx, todoId)
			if err != nil {
				return err
			}

			todo := services.MapRowToTodo(todoRow)

			return services.Render(echoCtx, corecomponents.Modal(corecomponents.ModalProps{Content: partials.TodoForm(partials.TodoFormProps{Todo: todo}), IsModalVisible: isModalVisible}))
		}

		return services.Render(echoCtx, corecomponents.Modal(corecomponents.ModalProps{IsModalVisible: false}))
	}

	todoRows, err := queries.ListTodos(dbCtx)
	if err != nil {
		return err
	}

	todos := services.MapRowsToTodo(todoRows)

	return services.Render(echoCtx, pages.TodosBase(todos, model.Todo{}, isModalVisible))
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

	todos, err := queries.ListTodos(dbCtx)
	if err != nil {
		return err
	}

	return services.Render(echoCtx, partials.TodosPartial(services.MapRowsToTodo(todos), model.Todo{}, false))
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

	todos, err := queries.ListTodos(dbCtx)
	if err != nil {
		return err
	}

	return services.Render(echoCtx, partials.TodosPartial(services.MapRowsToTodo(todos), model.Todo{}, false))
}
