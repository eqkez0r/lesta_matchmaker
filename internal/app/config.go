package app

import (
	"errors"
	"flag"
	mmcfg "github.com/eqkez0r/lesta_matchmaker/internal/matchmaker"
	servercfg "github.com/eqkez0r/lesta_matchmaker/internal/server"
	dbcfg "github.com/eqkez0r/lesta_matchmaker/internal/storage"
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	// Значение хоста по умолчанию
	defaultAddr = "localhost:8080"

	defaultPath = "./config.yaml"
)

var (
	// Объявление ошибки об недопустимых переменных
	ErrUnexpectedArguments = errors.New("unexpected arguments")
)

type Config struct {
	ServerConfig     servercfg.ServerConfig `json:"server_config" yaml:"server_config"`
	DatabaseConfig   dbcfg.DatabaseConfig   `json:"database_config" yaml:"database_config"`
	MatchmakerConfig mmcfg.MatchmakerConfig `json:"matchmaker_config" yaml:"matchmaker_config"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if len(flag.Args()) != 0 {
		return nil, ErrUnexpectedArguments
	}

	err := cleanenv.ReadConfig(defaultPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
