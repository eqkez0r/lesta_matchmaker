package memory

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/object/player"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage/storageerrors"
	"github.com/eqkez0r/lesta_matchmaker/pkg/logger"
	"sync"
)

type StorageMemory struct {
	mu        sync.Mutex
	logger    logger.ILogger
	PlayerMap map[string]player.Player
}

func New(
	l logger.ILogger,
) *StorageMemory {
	return &StorageMemory{
		logger: l,
	}
}

func (s *StorageMemory) PutPlayer(ctx context.Context, player player.Player) error {
	_, ok := s.PlayerMap[player.Name]
	if ok {
		return storageerrors.ErrPlayerInQueue
	}
	s.mu.Lock()
	s.PlayerMap[player.Name] = player
	s.mu.Unlock()
	return nil
}

func (s *StorageMemory) DeleteGroupPlayer(ctx context.Context, players []player.Player) error {
	s.mu.Lock()
	for _, pl := range players {
		delete(s.PlayerMap, pl.Name)
	}
	s.mu.Unlock()
	return nil
}

func (s *StorageMemory) GetAllPlayers(ctx context.Context) ([]player.Player, error) {
	s.mu.Lock()
	var players []player.Player
	for _, pl := range s.PlayerMap {
		players = append(players, pl)
	}
	s.mu.Unlock()
	return players, nil
}

func (s *StorageMemory) GracefulStop() {

}
