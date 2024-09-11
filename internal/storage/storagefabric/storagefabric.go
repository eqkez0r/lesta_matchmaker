package storagefabric

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/logger"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage/config"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage/memory"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage/pgx"
)

const (
	memoryType   = "memory"
	postgresType = "pgx"
)

func NewStorage(
	ctx context.Context,
	l logger.ILogger,
	cfg config.DatabaseConfig,
) (storage.IStorage, error) {
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
			return nil, storage.ErrUnknownStorageType
		}
	}
}
