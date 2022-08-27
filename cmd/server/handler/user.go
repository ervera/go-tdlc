package handler

import (
	"encoding/json"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/media"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/sendgrid"
	"github.com/ervera/tdlc-gin/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	userService     user.Service
	mediaService    media.Service
	sendgridService sendgrid.Service
}

func (c *userHandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user domain.User

		formValue := ctx.Request.FormValue("data")
		err := json.Unmarshal([]byte(formValue), &user)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}

		file, handler, errImage := ctx.Request.FormFile("image")
		if errImage == nil && handler != nil && handler.Size != 0 {
			url, err := c.mediaService.UploadMedia(ctx, file, handler)
			if err != nil {
				web.Error(ctx, 400, err.Error())
				return
			}
			user.Avatar = url
		}

		userCreated, err := c.userService.CreateUser(ctx, user)

		if err != nil {
			if handler != nil && handler.Size != 0 {
				c.mediaService.DeleteMedia(ctx, user.Avatar)
			}
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Response(ctx, 200, userCreated)

	}
}

func (c *userHandler) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := uuid.Must(uuid.Parse(ctx.Param("id")))
		user, err := c.userService.GetById(ctx, userId)

		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Response(ctx, 200, user)
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
		err = c.userService.Update(ctx, user)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Response(ctx, 200, nil)
	}
}

func (c *userHandler) SendMail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.sendgridService.SendMail()
		web.Response(ctx, 200, nil)
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

func NewHandlerUser(p user.Service, m media.Service, s sendgrid.Service) *userHandler {
	return &userHandler{
		userService:     p,
		mediaService:    m,
		sendgridService: s,
	}
}
