package config

type MatchmakerConfig struct {
	GroupSize uint `json:"group_size" yaml:"group_size" env:"GROUP_SIZE"`
}
