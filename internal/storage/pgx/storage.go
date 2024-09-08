package pgx

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/logger"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage/config"
	"github.com/eqkez0r/lesta_matchmaker/pkg/object/player"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	createTableQuery = `CREATE TABLE IF NOT EXISTS players (
    name text primary key not null,
    skill double precision,
    latency double precision
     )`
	putPlayerQuery = `INSERT INTO players(name, skill, latency) VALUES ($1, $2, $3)`
	deletePlayers  = `DELETE FROM players WHERE name = $1`
)

type PgxStorage struct {
	logger logger.ILogger
	conn   *pgxpool.Pool
}

func NewPgxStorage(
	ctx context.Context,
	logger logger.ILogger,
	cfg config.DatabaseConfig,
) (*PgxStorage, error) {
	conn, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	store := &PgxStorage{
		logger: logger,
		conn:   conn,
	}
	return store, nil
}

func (p *PgxStorage) PutPlayer(ctx context.Context, player player.Player) error {
	_, err := p.conn.Exec(ctx, putPlayerQuery, player.Name, player.Skill, player.Latency)
	if err != nil {
		return err
	}
	return nil
}

func (p *PgxStorage) DeleteGroupPlayer(ctx context.Context, players []player.Player) error {
	for _, pl := range players {
		_, err := p.conn.Exec(ctx, deletePlayers, pl.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PgxStorage) GracefulStop() {
	p.conn.Close()
}
