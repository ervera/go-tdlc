package handler

import (
	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		userId := uuid.Must(uuid.Parse(ctx.Param("id")))
		user, err := c.service.GetById(ctx, userId)

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
		err = c.service.Update(ctx, user)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, nil)
	}
}

// func (c *userHandler) CreateUserRelation() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		id := ctx.Param("id")
// 		err := c.service.SaveUserRelation(ctx, id)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Success(ctx, 200, nil)
// 	}
// }

func NewHandlerUser(p user.Service) *userHandler {
	return &userHandler{
		service: p,
	}
}
