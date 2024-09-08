package config

type DatabaseConfig struct {
	DatabaseType string `json:"DatabaseType" env:"DB_TYPE"`
	DatabaseURL  string `json:"DatabaseURL" env:"DB_DSN"`
}
