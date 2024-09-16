package matchmaker

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type config struct {
	GroupSize uint `json:"group_size" yaml:"group_size" env:"GROUP_SIZE"`
}

func initCfg() (*config, error) {
	const defaultPath = "./config.yaml"

	cfg := &config{}

	err := cleanenv.ReadConfig(defaultPath, cfg)
	if err != nil {
		return nil, err
	}

	log.Println("mm cfg", cfg)
	return cfg, nil
}
