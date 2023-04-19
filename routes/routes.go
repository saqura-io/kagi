package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saqura-io/kagi/routes/v1"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1Group := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("x-kagi-version", "v1")

		return c.Next()
	})

	v1.SetupRoutes(v1Group)
}
