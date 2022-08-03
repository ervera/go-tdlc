package handler

import (
	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/web"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func (c *userHandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var formattedBody domain.User
		err := ctx.ShouldBindJSON(&formattedBody)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		user, err := c.service.CreateUser(ctx, formattedBody)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, user)

	}
}

func NewHandlerUser(p user.Service) *userHandler {
	return &userHandler{
		service: p,
	}
}
