package main

import (
	"os"

	"github.com/hyperremix/todo-app-htmx/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	e := echo.New()

	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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
	}))
	e.Static("/assets", "./assets")

	handlers.DefineTodoRoutes(e)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
