package config

import (
	"errors"
	"flag"
	servercfg "github.com/eqkez0r/lesta_matchmaker/internal/server/config"
	dbcfg "github.com/eqkez0r/lesta_matchmaker/internal/storage/config"
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	// Значение хоста по умолчанию
	defaultAddr = "localhost:8080"
)

var (
	// Объявление ошибки об недопустимых переменных
	ErrUnexpectedArguments = errors.New("unexpected arguments")
)

type Config struct {
	ServerConfig   servercfg.ServerConfig `json:"server_config"`
	DatabaseConfig dbcfg.DatabaseConfig   `json:"database_config"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.ServerConfig.Host, "a", defaultAddr, "server host")
	flag.StringVar(&cfg.DatabaseConfig.DatabaseURL, "d", "", "database uri")
	flag.StringVar(&cfg.DatabaseConfig.DatabaseType, "t", "pgx", "database type")
	flag.Parse()

	if len(flag.Args()) != 0 {
		return nil, ErrUnexpectedArguments
	}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
