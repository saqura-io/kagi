package apikeyhandler

import (
	"github.com/gofiber/fiber/v2"
	apikey "github.com/saqura-io/kagi/pkg"
)

type apiKeyRequest struct {
	APIKey string `json:"api_key"`
}

func GenerateAPIKeyHandler(c *fiber.Ctx) error {
	apiKey := apikey.GenerateAPIKey()

	return c.JSON(fiber.Map{"api_key": apiKey})
}

func ValidateAPIKeyHandler(c *fiber.Ctx) error {
	var req apiKeyRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	valid := apikey.ValidateAPIKey(req.APIKey)

	return c.JSON(fiber.Map{"valid": valid})
}
