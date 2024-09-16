package main

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/app"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage"
	"github.com/eqkez0r/lesta_matchmaker/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := logger.New("zap")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	store, err := storage.NewStorage(ctx, l)
	if err != nil {
		l.Errorf("Error creating storage: %v", err)
		os.Exit(1)
	}

	defer store.GracefulStop()

	a, err := app.New(ctx, l, store)
	if err != nil {
		l.Errorf("Error creating app: %v", err)
		os.Exit(1)
	}
	a.Run(ctx)
}
