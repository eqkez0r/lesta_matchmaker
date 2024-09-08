package app

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/app/config"
	"github.com/eqkez0r/lesta_matchmaker/internal/logger"
	"github.com/eqkez0r/lesta_matchmaker/internal/server"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage"
	"sync"
)

type App struct {
	logger     logger.ILogger
	store      storage.IStorage
	cfg        *config.Config
	httpserver *server.HTTPServer
	matchmaker interface{}
}

func New(
	ctx context.Context,
	l logger.ILogger,
	store storage.IStorage,
	cfg *config.Config,
) *App {
	ser := server.New(ctx, l, cfg.ServerConfig, store)
	return &App{
		logger:     l,
		store:      store,
		cfg:        cfg,
		httpserver: ser,
	}
}

func (app *App) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go app.httpserver.Start(ctx, wg)

	wg.Wait()
}
