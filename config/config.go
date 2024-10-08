package config

import (
	"fmt"
	cfgutil "funnymovies/util/config"
	"os"
)

type Configuration struct {
	Port            int    `env:"PORT"`
	DbDsn           string `env:"DB_DSN"`
	JwtUserAlgo     string `env:"JWT_USER_ALGO"`
	JwtUserSecret   string `env:"JWT_USER_SECRET"`
	JwtUserDuration int    `env:"JWT_USER_DURATION"`
}

func Load() (*Configuration, error) {
	cfg := new(Configuration)
	stage := os.Getenv("STAGE")
	if stage == "" {
		stage = "local"
	}

	if err := cfgutil.Load(stage, cfg); err != nil {
		return nil, fmt.Errorf("error parsing environment config: %s", err)
	}
	return cfg, nil
}
