package utils

import (
	"github.com/gofiber/fiber/v2"
)

func HandleError(c *fiber.Ctx, statusCode int, errMsg string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"error":   errMsg,
	})
}

func HandleSuccess(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(data)
}
