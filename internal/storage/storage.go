package storage

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/object/player"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage/memory"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage/pgx"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage/storageerrors"
	"github.com/eqkez0r/lesta_matchmaker/pkg/logger"
)

const (
	memoryType   = "memory"
	postgresType = "pgx"
)

type IStorage interface {
	PutPlayer(ctx context.Context, player player.Player) error
	DeleteGroupPlayer(ctx context.Context, players []player.Player) error
	GetAllPlayers(ctx context.Context) ([]player.Player, error)
	GracefulStop()
}

func NewStorage(
	ctx context.Context,
	l logger.ILogger,
	cfg DatabaseConfig,
) (IStorage, error) {
	//TODO init and parse cfg
	switch cfg.DatabaseType {
	case memoryType:
		{
			return memory.New(l), nil
		}
	case postgresType:
		{
			return pgx.NewPgxStorage(ctx, l, cfg)
		}
	default:
		{
			return nil, storageerrors.ErrUnknownStorageType
		}
	}
}
