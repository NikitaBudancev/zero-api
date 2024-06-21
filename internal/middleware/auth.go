package middleware

import (
	"github.com/gofiber/fiber/v2"
	"zero_api/internal/config"
	"zero_api/internal/utils"
)

func AuthRequired(cfg *config.AppConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" || tokenStr != "Bearer "+cfg.BearerToken {
			return utils.HandleError(c, fiber.StatusUnauthorized, "Missing or invalid token")
		}
		return c.Next()
	}
}
