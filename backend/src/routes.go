package routes

import (
	"api/src/business"

	"github.com/gofiber/fiber/v2"
)

type healthCheck struct {
	Good bool
}

func Register(c *fiber.App) {
	// Prefix all routes with /api
	api := c.Group("/api")
	// Add a health check
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(healthCheck{true})
	})

	business.RegisterRoutes(api)
}
