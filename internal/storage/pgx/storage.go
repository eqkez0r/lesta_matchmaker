package pgx

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/object/player"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage"
	"github.com/eqkez0r/lesta_matchmaker/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxStorage struct {
	logger logger.ILogger
	conn   *pgxpool.Pool
}

func NewPgxStorage(
	ctx context.Context,
	logger logger.ILogger,
	cfg storage.DatabaseConfig,
) (*PgxStorage, error) {
	const createTableQuery = `CREATE TABLE IF NOT EXISTS players (
    name text primary key not null,
    skill double precision,
    latency double precision
     )`
	conn, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec(ctx, createTableQuery)
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
	const putPlayerQuery = `INSERT INTO players(name, skill, latency) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`
	_, err := p.conn.Exec(ctx, putPlayerQuery, player.Name, player.Skill, player.Latency)
	if err != nil {
		return err
	}
	return nil
}

func (p *PgxStorage) DeleteGroupPlayer(ctx context.Context, players []player.Player) error {
	const deletePlayers = `DELETE FROM players WHERE name = ANY($1)`
	nameList := make([]string, len(players))
	for i, pl := range players {
		nameList[i] = pl.Name
	}
	_, err := p.conn.Exec(ctx, deletePlayers, nameList)
	if err != nil {
		return err
	}

	return nil
}

func (p *PgxStorage) GetAllPlayers(ctx context.Context) ([]player.Player, error) {
	const getAllPlayers = `SELECT * FROM players`
	rows, err := p.conn.Query(ctx, getAllPlayers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var players []player.Player
	for rows.Next() {
		var player player.Player
		err = rows.Scan(&player.Name, &player.Skill, &player.Latency)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

func (p *PgxStorage) GracefulStop() {
	p.conn.Close()
}
