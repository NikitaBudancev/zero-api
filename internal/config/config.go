package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

type AppConfig struct {
	BearerToken string
	AppPort     string
}

type Config struct {
	DBConfig  DBConfig
	AppConfig AppConfig
}

func LoadConfig() (*Config, error) {

	viper.AutomaticEnv()

	config := &Config{
		DBConfig: DBConfig{
			DBHost:     viper.GetString("DB_HOST"),
			DBUser:     viper.GetString("POSTGRES_USER"),
			DBPassword: viper.GetString("POSTGRES_PASSWORD"),
			DBName:     viper.GetString("POSTGRES_DB"),
			DBPort:     viper.GetString("DB_PORT"),
		},
		AppConfig: AppConfig{
			BearerToken: viper.GetString("BEARER_TOKEN"),
			AppPort:     viper.GetString("API_PORT"),
		},
	}

	return config, nil
}
