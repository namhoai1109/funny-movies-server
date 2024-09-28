package cfgutil

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Load loads configuration from local .env file
func Load(stage string, out interface{}) error {
	if stage == "local" {
		return LoadLocal(out)
	}

	return LoadENV(out)
}

// Load loads configuration from local .env file
func LoadLocal(out interface{}) error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	if err := env.Parse(out); err != nil {
		return err
	}

	return nil
}

func LoadENV(out interface{}) error {

	return env.Parse(out)
}
