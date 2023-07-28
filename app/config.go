// IMPORTANT: THIS FILE CAN BE EDITED TO FIT YOUR NEEDS

package app

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDB       string `env:"POSTGRES_DB,required"` // Allows to be read by .env
	PostgresHost     string `env:"POSTGRES_HOST,required"`
	PostgresUser     string `env:"POSTGRES_USER,required"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required"`
	SecretKey        string `env:"SECRET_KEY,required"`
}

// Global configuration
var appConfig Config

// Loads the configuration from environment variables and sets it to the global config variable
func setConfig() error {
	// Loading environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		return err
	}

	// Loading environment variables into the config struct
	err = env.Parse(&appConfig)

	if err != nil {
		return err
	}

	return nil
}

// Returns the global configuration
func GetConfig() Config {
	return appConfig
}
