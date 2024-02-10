package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	Config struct {
		APP_ENV         string
		APP_NAME        string
		APP_PORT        string
		DB_POSTGRES_DSN string
	}
)

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file", err)
	}

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
}

func GetConfig() Config {
	return Config{
		APP_ENV:         viper.GetString("APP_ENV"),
		APP_NAME:        viper.GetString("APP_NAME"),
		APP_PORT:        viper.GetString("APP_PORT"),
		DB_POSTGRES_DSN: viper.GetString("DB_POSTGRES_DSN"),
	}
}
