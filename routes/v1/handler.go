package v1

import (
	"github.com/gofiber/fiber/v2"
	apikeyhandler "github.com/saqura-io/kagi/handlers"
)

func SetupRoutes(r fiber.Router) {
	r.Get("/generate", apikeyhandler.GenerateAPIKeyHandler)
	r.Post("/validate", apikeyhandler.ValidateAPIKeyHandler)
}
