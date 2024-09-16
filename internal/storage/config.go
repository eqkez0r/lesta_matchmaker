package storage

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type config struct {
	DatabaseType string `json:"database_type" yaml:"database_type" env:"DB_TYPE"`
	DatabaseURL  string `json:"database_url" yaml:"database_url" env:"DB_DSN"`
}

func initCfg() (*config, error) {
	const defaultPath = "./config.yaml"

	cfg := &config{}

	err := cleanenv.ReadConfig(defaultPath, cfg)
	if err != nil {
		return nil, err
	}
	log.Println("cfg db", cfg)
	return cfg, nil
}
