package handler

import (
	"github.com/ervera/tdlc-gin/internal/tweet"
	"github.com/ervera/tdlc-gin/pkg/web"
	"github.com/gin-gonic/gin"
)

type tweetHandler struct {
	service tweet.Service
}

type message struct {
	Mensaje string `json:"mensaje,omitempty"`
}

func (c *tweetHandler) CreateTweet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var m message
		err := ctx.ShouldBindJSON(&m)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		tweet, err := c.service.Save(ctx, m.Mensaje)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, tweet)

	}
}

func NewHandlerTweet(p tweet.Service) *tweetHandler {
	return &tweetHandler{
		service: p,
	}
}