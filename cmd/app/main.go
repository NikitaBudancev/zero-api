package main

import (
	"github.com/sirupsen/logrus"
	"zero_api/internal/config"
	"zero_api/internal/database"
	"zero_api/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		logrus.Fatalf("Could not load config: %v", err)
	}

	database.ConnectDatabase(&cfg.DBConfig)

	srv := server.NewServer(&cfg.AppConfig)

	if err := srv.Run(); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
