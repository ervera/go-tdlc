package handler

import (
	"strconv"

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

func (c *tweetHandler) GetTweetsByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pageString := ctx.Query("page")
		id := ctx.Param("id")
		pageInt, err := strconv.Atoi(pageString)
		page := int64(pageInt)
		if err != nil {
			page = 0
		}

		tweet, err := c.service.GetAllByUserId(ctx, id, page)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, tweet)

	}
}

func (c *tweetHandler) DeleteTweetsById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := c.service.DeleteOneById(ctx, id)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, nil)
	}
}

func NewHandlerTweet(p tweet.Service) *tweetHandler {
	return &tweetHandler{
		service: p,
	}
}
