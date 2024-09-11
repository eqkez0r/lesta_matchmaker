package config

type ServerConfig struct {
	Host string `json:"host" yaml:"host" env:"SERVER_HOST"`
}
