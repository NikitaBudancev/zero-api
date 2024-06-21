package server

import (
	"github.com/gofiber/fiber/v2"
	"zero_api/internal/config"
	"zero_api/internal/routes"
)

type Server struct {
	app *fiber.App
	cfg *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	app := fiber.New()
	return &Server{app: app, cfg: cfg}
}

func (s *Server) Run() error {
	routes.SetupRoutes(s.app, s.cfg)
	return s.app.Listen(":" + s.cfg.AppPort)
}
