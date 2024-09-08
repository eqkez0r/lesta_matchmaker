package memory

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/logger"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage"
	"github.com/eqkez0r/lesta_matchmaker/pkg/object/player"
)

type StorageMemory struct {
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
		return storage.ErrPlayerInQueue
	}
	s.PlayerMap[player.Name] = player
	return nil
}

func (s *StorageMemory) DeleteGroupPlayer(ctx context.Context, players []player.Player) error {
	for _, pl := range players {
		delete(s.PlayerMap, pl.Name)
	}
	return nil
}

func (s *StorageMemory) GracefulStop() {

}
