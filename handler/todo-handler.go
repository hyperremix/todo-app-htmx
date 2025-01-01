package handler

import (
	"strconv"

	"github.com/hyperremix/todo-app-htmx/components/corecomponents"
	"github.com/hyperremix/todo-app-htmx/components/pages"
	"github.com/hyperremix/todo-app-htmx/components/partials"
	"github.com/hyperremix/todo-app-htmx/db"
	"github.com/hyperremix/todo-app-htmx/mapper"
	"github.com/hyperremix/todo-app-htmx/model"
	"github.com/hyperremix/todo-app-htmx/template"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func registerTodoRoutes(e *echo.Echo, connPool *pgxpool.Pool) {
	e.GET("/", getTodos(connPool))
	e.POST("/", createTodo(connPool))
	e.PUT("/:id", updateTodo(connPool))
	e.DELETE("/:id", deleteTodo(connPool))
}

func getTodos(connPool *pgxpool.Pool) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx := echoCtx.Request().Context()
		conn, err := connPool.Acquire(ctx)

		if err != nil {
			return err
		}
		defer conn.Release()

		queries := db.New(conn)

		isHxRequest := echoCtx.Request().Header.Get("HX-Request") == "true"

		var request model.GetTodosRequest
		if err := echoCtx.Bind(&request); err != nil {
			return err
		}

		todo := model.Todo{}
		if request.TodoID != 0 {
			todoRow, err := queries.GetTodoById(ctx, request.TodoID)
			if err != nil {
				return err
			}

			todo = mapper.MapRowToTodo(todoRow)
		}

		if isHxRequest && request.IsUpdateModalVisible {
			return template.Render(echoCtx, partials.UpdateTodoModal(todo, true))
		}

		if isHxRequest && request.IsDeleteModalVisible {
			return template.Render(echoCtx, partials.DeleteTodoModal(todo, true))
		}

		if isHxRequest {
			return template.Render(echoCtx, corecomponents.Modal(corecomponents.ModalProps{IsModalVisible: false}))
		}

		todoRows, err := queries.ListOpenTodos(ctx)
		if err != nil {
			return err
		}

		todos := mapper.MapRowsToTodo(todoRows)

		return template.Render(echoCtx, pages.TodosBase(partials.TodosProps{Todos: todos, Todo: todo, IsUpdateModalVisible: request.IsUpdateModalVisible, IsDeleteModalVisible: request.IsDeleteModalVisible}))
	}
}

func createTodo(connPool *pgxpool.Pool) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx := echoCtx.Request().Context()
		conn, err := connPool.Acquire(ctx)
		if err != nil {
			return err
		}
		defer conn.Release()

		queries := db.New(conn)

		todo, err := mapper.MapRequestToTodo(echoCtx)
		if err != nil {
			return err
		}

		_, err = queries.InsertTodo(ctx, mapper.MapTodoToInsert(todo))
		if err != nil {
			return err
		}

		todoRows, err := queries.ListOpenTodos(ctx)
		if err != nil {
			return err
		}

		todos := mapper.MapRowsToTodo(todoRows)
		return template.Render(echoCtx, partials.TodosPartial(partials.TodosProps{Todos: todos}))
	}
}

func updateTodo(connPool *pgxpool.Pool) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx := echoCtx.Request().Context()
		conn, err := connPool.Acquire(ctx)
		if err != nil {
			return err
		}
		defer conn.Release()

		queries := db.New(conn)

		todo, err := mapper.MapRequestToTodo(echoCtx)
		if err != nil {
			return err
		}

		_, err = queries.UpdateTodo(ctx, mapper.MapTodoToUpdate(todo))
		if err != nil {
			return err
		}

		todoRows, err := queries.ListOpenTodos(ctx)
		if err != nil {
			return err
		}

		todos := mapper.MapRowsToTodo(todoRows)
		return template.Render(echoCtx, partials.TodosPartial(partials.TodosProps{Todos: todos}))
	}
}

func deleteTodo(connPool *pgxpool.Pool) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx := echoCtx.Request().Context()
		conn, err := connPool.Acquire(ctx)
		if err != nil {
			return err
		}
		defer conn.Release()

		queries := db.New(conn)

		idParam := echoCtx.Param("id")
		todoID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			return err
		}

		_, err = queries.DeleteTodoById(ctx, todoID)
		if err != nil {
			return err
		}

		todoRows, err := queries.ListOpenTodos(ctx)
		if err != nil {
			return err
		}

		todos := mapper.MapRowsToTodo(todoRows)
		return template.Render(echoCtx, partials.TodosPartial(partials.TodosProps{Todos: todos}))
	}
}
