package config

type ServerConfig struct {
	Host string `yaml:"host" env:"SERVER_HOST"`
}
