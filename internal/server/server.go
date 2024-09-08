package server

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/logger"
	servercfg "github.com/eqkez0r/lesta_matchmaker/internal/server/config"
	"github.com/eqkez0r/lesta_matchmaker/internal/server/handlers"
	"github.com/eqkez0r/lesta_matchmaker/internal/server/middleware"
	"github.com/eqkez0r/lesta_matchmaker/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type HTTPServer struct {
	server *http.Server
	engine *gin.Engine
	store  storage.IStorage
	logger logger.ILogger
}

func New(
	ctx context.Context,
	l logger.ILogger,
	config servercfg.ServerConfig,
	store storage.IStorage,

) *HTTPServer {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery(), middleware.Logger(l))
	router.Handle("POST", handlers.PutPlayerPath, handlers.AddPlayerHandler(ctx, l, store))

	return &HTTPServer{
		server: &http.Server{
			Addr:    config.Host,
			Handler: router,
		},
		engine: router,
		store:  store,
		logger: l,
	}
}

func (s *HTTPServer) Start(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			s.logger.Errorf("http server start error: %s", err.Error())
		}
	}()
	defer wg.Done()
	defer s.server.Shutdown(ctx)

	<-ctx.Done()
}
