package handler

import (
	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/jwt"
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

func (c *userHandler) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// id := ctx.Query("id")
		// fmt.Println(id)
		// fmt.Println("adsdsadsa")
		id := ctx.Param("id")
		user, err := c.service.GetUserById(ctx, id)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, user)

	}
}

func (c *userHandler) UpdateSelfUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user domain.User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		err = c.service.UpdateSelf(ctx, user, jwt.UserID)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, nil)

	}
}

func NewHandlerUser(p user.Service) *userHandler {
	return &userHandler{
		service: p,
	}
}
