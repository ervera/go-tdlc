package handler

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strings"

// 	"github.com/ervera/tdlc-gin/internal/domain"
// 	"github.com/ervera/tdlc-gin/internal/media"
// 	"github.com/ervera/tdlc-gin/internal/user"
// 	"github.com/ervera/tdlc-gin/pkg/web"
// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type password struct {
// 	Password string `json:"password"`
// }

// type email struct {
// 	Email string `json:"email"`
// }

// type UserHandler struct {
// 	userService  user.Service
// 	mediaService media.Service
// }

// func (c *UserHandler) CreateUser() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var user domain.User

// 		formValue := ctx.Request.FormValue("data")
// 		err := json.Unmarshal([]byte(formValue), &user)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}

// 		file, handler, errImage := ctx.Request.FormFile("image")
// 		if errImage == nil && handler != nil && handler.Size != 0 {
// 			url, err := c.mediaService.UploadMedia(ctx, file, handler)
// 			if err != nil {
// 				web.Error(ctx, 400, err.Error())
// 				return
// 			}
// 			user.Avatar = strings.Replace(url, "upload", "upload/c_fill,g_auto,h_800,w_800", 1)
// 		}

// 		userCreated, err := c.userService.CreateUser(ctx, user)

// 		if err != nil {
// 			if handler != nil && handler.Size != 0 {
// 				c.mediaService.DeleteMedia(ctx, user.Avatar)
// 			}
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Response(ctx, 200, userCreated)
// 	}
// }

// func (c *UserHandler) GetUserById() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		userId := uuid.Must(uuid.Parse(ctx.Param("id")))
// 		user, err := c.userService.GetById(ctx, userId)

// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Response(ctx, 200, user)
// 	}
// }

// func (c *UserHandler) UpdateSelfUser() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var user domain.User
// 		err := ctx.ShouldBindJSON(&user)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		err = c.userService.Update(ctx, user)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Response(ctx, 200, user)
// 	}
// }

// func (c *UserHandler) ForgotPassword() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var email email
// 		err := ctx.ShouldBind(&email)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		err = c.userService.SendEmailWithPassword(ctx, email.Email)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Response(ctx, 200, nil)
// 	}
// }

// func (c *UserHandler) NewPassword() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var password password
// 		err := ctx.ShouldBind(&password)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		err = c.userService.NewPassword(ctx, password.Password)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Response(ctx, http.StatusOK, nil)
// 	}
// }

// func (c *UserHandler) UpdateMedia() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		auxUrl := ""
// 		var user domain.User
// 		typeField := ctx.Param("type")
// 		formValue := ctx.Request.FormValue("data")

// 		err := json.Unmarshal([]byte(formValue), &user)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}

// 		file, handler, errImage := ctx.Request.FormFile("image")
// 		if errImage == nil && handler != nil && handler.Size != 0 {
// 			url, err := c.mediaService.UploadMedia(ctx, file, handler)
// 			if err != nil {
// 				web.Error(ctx, 400, err.Error())
// 				return
// 			}
// 			if typeField == "banner" {
// 				auxUrl = user.Banner
// 				user.Banner = strings.Replace(url, "upload", "upload/c_fill,g_auto,h_199,w_800", 1)
// 			}
// 			if typeField == "avatar" {
// 				auxUrl = user.Avatar
// 				user.Avatar = strings.Replace(url, "upload", "upload/c_fill,g_auto,h_800,w_800", 1)
// 			}
// 			// updates := map[string]string{typeField: url}
// 			// update(&user, updates)
// 		}

// 		err = c.userService.UpdateMedia(ctx, user)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Response(ctx, 200, user)
// 		if auxUrl != "" {
// 			c.mediaService.DeleteMedia(ctx, auxUrl)
// 		}
// 	}
// }

// func (c *UserHandler) DeleteMedia() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var user domain.User
// 		mediaType := ctx.Param("type")

// 		err := ctx.ShouldBind(&user)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 		}

// 		err = c.userService.DeleteMedia(ctx, user, mediaType)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Response(ctx, 200, user)
// 		if mediaType == "avatar" && user.Avatar != "" {
// 			c.mediaService.DeleteMedia(ctx, user.Avatar)
// 		}
// 		if mediaType == "banner" && user.Banner != "" {
// 			c.mediaService.DeleteMedia(ctx, user.Banner)
// 		}
// 	}
// }

// // func update(v interface{}, updates map[string]string) {
// // 	rv := reflect.ValueOf(v).Elem()
// // 	for key, val := range updates {
// // 		fv := rv.FieldByName(key)
// // 		fv.SetString(val)
// // 	}
// // }

// func NewHandlerUser(p user.Service, m media.Service) *UserHandler {
// 	return &UserHandler{
// 		userService:  p,
// 		mediaService: m,
// 	}
// }
