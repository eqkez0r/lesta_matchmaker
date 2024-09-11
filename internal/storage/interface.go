package storage

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/pkg/object/player"
)

type IStorage interface {
	PutPlayer(ctx context.Context, player player.Player) error
	DeleteGroupPlayer(ctx context.Context, players []player.Player) error
	GetAllPlayers(ctx context.Context) ([]player.Player, error)
	GracefulStop()
}
