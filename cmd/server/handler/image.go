package handler

// import (
// 	"io"
// 	"os"
// 	"strings"

// 	"github.com/ervera/tdlc-gin/internal/domain"
// 	"github.com/ervera/tdlc-gin/internal/user"
// 	"github.com/ervera/tdlc-gin/pkg/jwt"
// 	"github.com/ervera/tdlc-gin/pkg/web"
// 	"github.com/gin-gonic/gin"
// )

// type imageHandler struct {
// 	service user.Service
// }

// func (c *imageHandler) UploadUserImage() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		file, handler, err := ctx.Request.FormFile("avatar")
// 		var extensions = strings.Split(handler.Filename, ".")[1]
// 		var archivo string = "uploads/avatar" + jwt.UserID + "." + extensions

// 		f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		_, err = io.Copy(f, file)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error()+"error al crear la copia")
// 			return
// 		}

// 		var user domain.User
// 		user.Avatar = jwt.UserID + "." + extensions
// 		err = c.service.UpdateSelf(ctx, user, jwt.UserID)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Success(ctx, 200, nil)
// 	}
// }
