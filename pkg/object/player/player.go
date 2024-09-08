package player

type Player struct {
	Name    string  `json:"name"`
	Skill   float32 `json:"skill"`
	Latency float32 `json:"latency"`
}
