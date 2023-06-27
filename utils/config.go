package utils

import "github.com/kelseyhightower/envconfig"

type Config struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"` // Allows to be read by .env
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	SecretKey        string `envconfig:"SECRET_KEY"`
}

// Global configuration
var appConfig Config

func SetConfig() error {
	// Loading environment variables into the config struct
	err := envconfig.Process("", &appConfig)

	if err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	// Returns the global configuration
	return appConfig
}
