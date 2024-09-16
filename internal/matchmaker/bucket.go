package matchmaker

import (
	"github.com/eqkez0r/lesta_matchmaker/internal/object/player"
	"math"
	"sort"
	"time"
)

type skillbucket struct {
	players []player.Player
}

type stat struct {
	minSkill, maxSkill, avgSkill                   float32
	minLatency, maxLatency, avgLatency             float32
	minTimeInQueue, maxTimeInQueue, avgTimeInQueue time.Duration
	playersList                                    []string
}

func (b *skillbucket) PutPlayer(player player.Player) {
	b.players = append(b.players, player)
}

func (b *skillbucket) Players() []player.Player {
	return b.players
}

func (b *skillbucket) Reset(groupSize uint) {
	b.players = b.players[groupSize:]
}

func (b *skillbucket) SortByLatency() {
	sort.Slice(b.players, func(i, j int) bool {
		return b.players[i].Latency < b.players[j].Latency
	})
}

func (b *skillbucket) Stat(groupSize uint) stat {
	n := time.Now()
	s := stat{
		minSkill:       math.MaxFloat32,
		maxSkill:       0.0,
		avgSkill:       0.0,
		minLatency:     math.MaxFloat32,
		maxLatency:     0.0,
		avgLatency:     0.0,
		playersList:    make([]string, 0, groupSize),
		maxTimeInQueue: n.Sub(b.players[0].InQueueFrom),
		minTimeInQueue: n.Sub(b.players[0].InQueueFrom),
	}

	for i, pl := range b.players {
		if i == int(groupSize) {
			break
		}
		s.playersList = append(s.playersList, pl.Name)
		s.avgSkill += pl.Skill
		s.avgLatency += pl.Latency
		dur := n.Sub(pl.InQueueFrom)
		s.avgTimeInQueue += dur

		if pl.Skill < s.minSkill {
			s.minSkill = pl.Skill
		}
		if pl.Skill > s.maxSkill {
			s.maxSkill = pl.Skill
		}

		if pl.Latency < s.minLatency {
			s.minLatency = pl.Latency
		}
		if pl.Latency > s.maxLatency {
			s.maxLatency = pl.Latency
		}

		if dur < s.minTimeInQueue {
			s.minTimeInQueue = n.Sub(pl.InQueueFrom)
		}
		if dur > s.maxTimeInQueue {
			s.minTimeInQueue = n.Sub(pl.InQueueFrom)
		}

	}

	s.avgSkill /= float32(groupSize)
	s.avgLatency /= float32(groupSize)
	s.avgLatency /= float32(groupSize)

	return s
}
