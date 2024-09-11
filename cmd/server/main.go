package main

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/app"
	"github.com/eqkez0r/lesta_matchmaker/internal/app/config"
	zaplogger "github.com/eqkez0r/lesta_matchmaker/internal/logger/zap"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage/storagefabric"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := zaplogger.New()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.NewConfig()
	if err != nil {
		l.Errorf("Error loading config: %v", err)
		os.Exit(1)
	}

	l.Infof("app started with config %v", cfg)

	store, err := storagefabric.NewStorage(ctx, l, cfg.DatabaseConfig)
	if err != nil {
		l.Errorf("Error creating storage: %v", err)
		os.Exit(1)
	}

	defer store.GracefulStop()

	a := app.New(ctx, l, cfg, store)
	a.Run(ctx)
}
