package app

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/matchmaker"
	"github.com/eqkez0r/lesta_matchmaker/internal/object/player"
	"github.com/eqkez0r/lesta_matchmaker/internal/server"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage"
	"github.com/eqkez0r/lesta_matchmaker/pkg/logger"
	"sync"
)

type App struct {
	logger     logger.ILogger
	store      storage.IStorage
	httpserver *server.HTTPServer
	matchmaker *matchmaker.Matchmaker
	playerChan chan player.Player
}

func New(
	ctx context.Context,
	l logger.ILogger,
	store storage.IStorage,
) (*App, error) {
	pch := make(chan player.Player, 100)
	ser, err := server.New(ctx, l, store, pch)
	if err != nil {
		return nil, err
	}
	mm, err := matchmaker.NewMatchmaker(l, store)
	if err != nil {
		return nil, err
	}
	return &App{
		logger:     l,
		store:      store,
		httpserver: ser,
		matchmaker: mm,
		playerChan: pch,
	}, nil
}

func (app *App) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go app.httpserver.Start(ctx, wg)
	go app.matchmaker.Run(ctx, wg, app.playerChan)

	defer app.Stop()
	wg.Wait()
}

func (app *App) Stop() {
	close(app.playerChan)

}
