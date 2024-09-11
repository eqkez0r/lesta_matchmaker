package storage

type DatabaseConfig struct {
	DatabaseType string `json:"database_type" yaml:"database_type" env:"DB_TYPE"`
	DatabaseURL  string `json:"database_url" yaml:"database_url" env:"DB_DSN"`
}
