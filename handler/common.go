package handler

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func RegisterHandlerRoutes(e *echo.Echo, connPool *pgxpool.Pool) {
	registerTodoRoutes(e, connPool)
}
