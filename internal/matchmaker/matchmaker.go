package matchmaker

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/logger"
	"github.com/eqkez0r/lesta_matchmaker/internal/matchmaker/config"
	"github.com/eqkez0r/lesta_matchmaker/pkg/object/player"
	"math"
	"sync"
)

type ClearPlayersProvider interface {
	DeleteGroupPlayer(ctx context.Context, players []player.Player) error
}

type Matchmaker struct {
	logger    logger.ILogger
	groupSize uint
	mu        sync.RWMutex // for highload it change to rw
	m         map[float32]skillbucket
	store     ClearPlayersProvider

	groupCounter uint64
}

func NewMatchmaker(
	logger logger.ILogger,
	cfg config.MatchmakerConfig,
	store ClearPlayersProvider,
) *Matchmaker {
	return &Matchmaker{
		groupSize: cfg.GroupSize,
		logger:    logger,
		store:     store,
	}
}

func (m *Matchmaker) Run(
	ctx context.Context,
	wg *sync.WaitGroup,
	playerPooler <-chan player.Player,
) {
	for {
		select {
		case <-ctx.Done():
			{
				m.logger.Info("matchmaker shutting down")
				wg.Done()
				return
			}
		case pl := <-playerPooler:
			{
				m.mu.Lock()
				skill := float32(math.Floor(float64(pl.Skill)))
				b := m.m[skill]
				b.PutPlayer(pl)
				pls := m.m[skill]
				if len(pls.Players()) >= int(m.groupSize) {
					b.SortByLatency()
					err := m.store.DeleteGroupPlayer(ctx, pls.Players())
					if err != nil {
						m.logger.Error(err)
						m.mu.Unlock()
						continue
					}
					st := b.Stat(m.groupSize)
					m.logger.Infof("registered match %d \n"+
						"skill  min/max/avg %f/%f/%f \n"+
						"latency min/max/avg %f/%f/%f \n"+
						"spent time in queue min/max/avg %t/%t/%t \n"+
						"players %v",
						m.groupCounter,
						st.minSkill, st.maxSkill, st.avgSkill,
						st.minLatency, st.maxLatency, st.avgLatency,
						st.minLatency, st.maxLatency, st.avgLatency,
						st.playersList)
					//здесь какая-то регистрация матча
				}
				m.mu.Unlock()
			}
		default:
			{

			}
		}

	}
}
