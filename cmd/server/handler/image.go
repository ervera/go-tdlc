package handler

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/jwt"
	"github.com/ervera/tdlc-gin/pkg/web"
	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	service user.Service
}

func (c *ImageHandler) UploadUserImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramImage := ctx.Param("type")
		file, _, _ := ctx.Request.FormFile("image")
		cld, _ := cloudinary.NewFromParams(cloudname, cloudapikey, cloudinarysecret)
		options := uploader.UploadParams{}
		resp, _ := cld.Upload.Upload(ctx, file, options)

		var user domain.User
		if paramImage == "avatar" {
			user.Avatar.ImgUrl = resp.URL
			user.Avatar.PublicID = resp.PublicID
		}
		if paramImage == "banner" {
			user.Banner.ImgUrl = resp.URL
			user.Banner.PublicID = resp.PublicID
		}

		err := c.service.UpdateSelf(ctx, user, jwt.UserID)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, nil)
	}
}

func (c *ImageHandler) DeleteUserImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramImage := ctx.Param("type")
		user, err := c.service.GetUserById(ctx, jwt.UserID)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		cld, _ := cloudinary.NewFromParams(cloudname, cloudapikey, cloudinarysecret)

		var userUpdate domain.User
		if paramImage == "avatar" {
			resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: user.Avatar.PublicID})
			if err != nil {
				web.Error(ctx, 400, err.Error())
				return
			}
			fmt.Println(resp)
			userUpdate.Avatar.ImgUrl = "null"
			userUpdate.Avatar.PublicID = "null"
		}
		if paramImage == "banner" {
			resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: user.Banner.PublicID})
			if err != nil {
				web.Error(ctx, 400, err.Error())
				return
			}
			fmt.Println(resp)
			userUpdate.Banner.ImgUrl = "null"
			userUpdate.Banner.PublicID = "null"
		}
		err = c.service.UpdateSelf(ctx, userUpdate, jwt.UserID)
		if err != nil {
			web.Error(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, nil)
	}
}

func NewImageHandler(p user.Service) *ImageHandler {
	return &ImageHandler{
		service: p,
	}
}
