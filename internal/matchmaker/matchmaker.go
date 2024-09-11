package matchmaker

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/object/player"
	"github.com/eqkez0r/lesta_matchmaker/pkg/logger"
	"math"
	"sync"
)

type ClearPlayersProvider interface {
	GetAllPlayers(ctx context.Context) ([]player.Player, error)
	DeleteGroupPlayer(ctx context.Context, players []player.Player) error
}

type Matchmaker struct {
	logger    logger.ILogger
	groupSize uint
	m         map[float32]skillbucket
	store     ClearPlayersProvider

	groupCounter uint64
}

func NewMatchmaker(
	logger logger.ILogger,
	cfg MatchmakerConfig,
	store ClearPlayersProvider,
) *Matchmaker {
	return &Matchmaker{
		m:         make(map[float32]skillbucket),
		groupSize: cfg.GroupSize,
		logger:    logger,
		store:     store,
	}
}

func (m *Matchmaker) Run(
	ctx context.Context,
	wg *sync.WaitGroup,
	playerPooler chan player.Player,
) {
	players, err := m.store.GetAllPlayers(ctx)
	if err != nil {
		m.logger.Error(err)
	}
	for _, pl := range players {
		playerPooler <- pl
	}
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
				m.logger.Infof("matchmaker got player %v", pl)
				skill := float32(math.Floor(float64(pl.Skill)))
				pls := m.m[skill]
				pls.PutPlayer(pl)
				if len(pls.Players()) >= int(m.groupSize) {
					pls.SortByLatency()
					err = m.store.DeleteGroupPlayer(ctx, pls.Players())
					if err != nil {
						m.logger.Error(err)
						continue
					}
					st := pls.Stat(m.groupSize)
					m.logger.Infof("registered match %d "+
						"skill  min/max/avg %f/%f/%f "+
						"latency min/max/avg %f/%f/%f "+
						"spent time in queue min/max/avg %t/%t/%t "+
						"players %v",
						m.groupCounter,
						st.minSkill, st.maxSkill, st.avgSkill,
						st.minLatency, st.maxLatency, st.avgLatency,
						st.minLatency, st.maxLatency, st.avgLatency,
						st.playersList)
					pls.Reset(m.groupSize)
					//здесь какая-то регистрация матча
					m.groupCounter++
				}
				m.m[skill] = pls
			}
		default:
			{

			}
		}

	}
}
