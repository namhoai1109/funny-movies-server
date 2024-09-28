package cfgutil

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Load loads configuration from local .env file
func Load(out interface{}) error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	if err := env.Parse(out); err != nil {
		return err
	}

	return nil
}
