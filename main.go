package main

import (
	"context"
	"os"

	"github.com/hyperremix/todo-app-htmx/handler"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	godotenv.Load(".env")
	e := echo.New()

	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, os.Getenv("TODO_APP_DB_CONNECTION_STRING"))
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer connPool.Close()

	logger := zerolog.New(os.Stdout)

	e.Use(
		middleware.Recover(),
		middleware.Static(os.Getenv("TODO_APP_ASSETS_PATH")),
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogStatus:  true,
			LogLatency: true,
			LogMethod:  true,
			LogURI:     true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				logger.Info().
					Timestamp().
					Int("status", v.Status).
					Dur("latency", v.Latency).
					Str("method", v.Method).
					Str("URI", v.URI).
					Msg("request")

				return nil
			},
			HandleError: true,
		}))

	e.Static("/assets", os.Getenv("TODO_APP_ASSETS_PATH"))

	handler.RegisterHandlerRoutes(e, connPool)

	e.Logger.Fatal(e.Start(":8080"))
}
