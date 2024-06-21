package routes

import (
	"github.com/gofiber/fiber/v2"
	"zero_api/internal/config"
	"zero_api/internal/handlers"
	"zero_api/internal/middleware"
)

func SetupRoutes(app *fiber.App, cfg *config.AppConfig) {
	app.Get("/list", handlers.GetNews)

	auth := app.Group("", middleware.AuthRequired(cfg))
	auth.Post("/edit/:Id", handlers.EditNews)
}
