package config

import (
	"fmt"
	cfgutil "funnymovies/util/config"
)

type Configuration struct {
	Port  int    `env:"PORT"`
	DbDsn string `env:"DB_DSN"`
}

func Load() (*Configuration, error) {
	cfg := new(Configuration)
	if err := cfgutil.Load(cfg); err != nil {
		return nil, fmt.Errorf("error parsing environment config: %s", err)
	}
	return cfg, nil
}
