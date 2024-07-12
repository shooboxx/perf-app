package main

import (
	routes "api/src"
	"os"

	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Info("Sentry initialization failed: %v\n", err)
	}
	app := fiber.New()
	app.Use(cors.New())

	sentryHandler := sentryfiber.New(sentryfiber.Options{
		Repanic:         true,
		WaitForDelivery: true,
	})
	app.Use(sentryHandler)

	routes.Register(app)

	app.Listen(":8080")
}
