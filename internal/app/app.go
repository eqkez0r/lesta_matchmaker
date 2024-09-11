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
	cfg        *Config
	httpserver *server.HTTPServer
	matchmaker *matchmaker.Matchmaker
	playerChan chan player.Player
}

func New(
	ctx context.Context,
	l logger.ILogger,
	cfg *Config,
	store storage.IStorage,

) *App {
	pch := make(chan player.Player, 100)
	ser := server.New(ctx, l, cfg.ServerConfig, store, pch)
	return &App{
		logger:     l,
		store:      store,
		cfg:        cfg,
		httpserver: ser,
		matchmaker: matchmaker.NewMatchmaker(l, cfg.MatchmakerConfig, store),
		playerChan: pch,
	}
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
