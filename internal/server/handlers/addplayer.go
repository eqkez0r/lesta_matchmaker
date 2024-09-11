package handlers

import (
	"context"
	"github.com/eqkez0r/lesta_matchmaker/internal/object/player"
	"github.com/eqkez0r/lesta_matchmaker/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	PutPlayerPath = "/users"
)

type PutPlayerQueueProvider interface {
	PutPlayer(ctx context.Context, player player.Player) error
}

func AddPlayerHandler(
	ctx context.Context,
	l logger.ILogger,
	p PutPlayerQueueProvider,
	pch chan<- player.Player,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pl player.Player
		if err := c.ShouldBindJSON(&pl); err != nil {
			l.Errorf("parsing body error: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err := p.PutPlayer(ctx, pl); err != nil {
			l.Errorf("put player error: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		pch <- pl
		l.Infof("player was added to queue %v", pl)
	}
}
