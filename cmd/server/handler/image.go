package handler

import (
	"github.com/ervera/tdlc-gin/internal/user"
)

type ImageHandler struct {
	service user.Service
}

const (
	png  string = "png"
	jpg  string = "jpg"
	jpeg string = "jpeg"
)

var imgType = map[string]struct{}{
	png:  {},
	jpg:  {},
	jpeg: {},
}

const (
	cloudname        = "dangvuvyq"
	cloudapikey      = "754218821349648"
	cloudinarysecret = "rjiWDoS5G0yNdiY4NZkEXtvit8k"
)

// func (c *ImageHandler) UploadUserImage() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		paramImage := ctx.Param("type")
// 		file, handler, err := ctx.Request.FormFile("image")
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		format := strings.Split(handler.Filename, ".")[1]
// 		_, exist := imgType[format]
// 		if !exist {
// 			web.Error(ctx, 400, "image format is not correct")
// 			return
// 		}
// 		cld, err := cloudinary.NewFromParams(cloudname, cloudapikey, cloudinarysecret)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		options := uploader.UploadParams{}
// 		resp, err := cld.Upload.Upload(ctx, file, options)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		var user domain.User
// 		if paramImage == "avatar" {
// 			user.Avatar.ImgUrl = resp.URL
// 			user.Avatar.PublicID = resp.PublicID
// 		}
// 		if paramImage == "banner" {
// 			user.Banner.ImgUrl = resp.URL
// 			user.Banner.PublicID = resp.PublicID
// 		}
// 		err = c.service.UpdateSelf(ctx, user, jwt.UserID)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Success(ctx, 200, nil)
// 	}
// }

// func (c *ImageHandler) DeleteUserImage() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		paramImage := ctx.Param("type")
// 		user, err := c.service.GetUserById(ctx, jwt.UserID)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		cld, _ := cloudinary.NewFromParams(cloudname, cloudapikey, cloudinarysecret)

// 		var userUpdate domain.User
// 		if paramImage == "avatar" {
// 			resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: user.Avatar.PublicID})
// 			if err != nil {
// 				web.Error(ctx, 400, err.Error())
// 				return
// 			}
// 			fmt.Println(resp)
// 			userUpdate.Avatar.ImgUrl = "null"
// 			userUpdate.Avatar.PublicID = "null"
// 		}
// 		if paramImage == "banner" {
// 			resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: user.Banner.PublicID})
// 			if err != nil {
// 				web.Error(ctx, 400, err.Error())
// 				return
// 			}
// 			fmt.Println(resp)
// 			userUpdate.Banner.ImgUrl = "null"
// 			userUpdate.Banner.PublicID = "null"
// 		}
// 		err = c.service.UpdateSelf(ctx, userUpdate, jwt.UserID)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Success(ctx, 200, nil)
// 	}
// }

func NewImageHandler(p user.Service) *ImageHandler {
	return &ImageHandler{
		service: p,
	}
}
