package server

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type config struct {
	Host string `json:"host" yaml:"host" env:"SERVER_HOST"`
}

func initCfg() (*config, error) {
	const defaultPath = "./config.yaml"

	cfg := &config{}

	err := cleanenv.ReadConfig(defaultPath, cfg)
	if err != nil {
		return nil, err
	}

	log.Println("server cfg", cfg)

	return cfg, nil
}
