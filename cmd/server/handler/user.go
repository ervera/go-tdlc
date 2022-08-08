package handler

import (
	"fmt"
	"io"
	"os"
	"strings"

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

func (c *userHandler) UploadUserImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramImage := ctx.Param("type")
		file, handler, err := ctx.Request.FormFile("image")
		if err != nil {
			fmt.Println("a")
			web.Error(ctx, 400, err.Error())
			return
		}
		var extensions = strings.Split(handler.Filename, ".")[1]
		var archivo string = "uploads/" + paramImage + "/" + jwt.UserID + "." + extensions

		f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("b")
			web.Error(ctx, 400, err.Error())
			return
		}
		_, err = io.Copy(f, file)
		if err != nil {
			web.Error(ctx, 400, err.Error()+"error al crear la copia")
			return
		}

		var user domain.User
		if paramImage == "avatar" {
			user.Avatar = jwt.UserID + "." + extensions
		}
		if paramImage == "banner" {
			user.Banner = jwt.UserID + "." + extensions
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
